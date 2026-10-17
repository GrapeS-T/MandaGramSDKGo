# Building and consuming the Go SDK

## Lifecycle restriction — read first

A Go `-buildmode=c-shared` DLL embeds the Go runtime. Unloading it with `FreeLibrary` is unsupported and unsafe, even if no plugin callback has run: loading can start runtime initialization and runtime-owned threads. The current MandaGram client opens and closes libraries during metadata discovery and unloads them on disable. Consequently this SDK's artifacts are **not compatible with that normal client lifecycle**. Static C metadata avoids entering Go to query metadata, but does not prevent runtime startup on library load and does not make unloading safe.

Use a fresh, disposable native harness process that loads the DLL once, keeps the DLL and API table alive, invokes metadata/init and callbacks, and exits without explicitly unloading. Do not use a harness with automatic library cleanup or reload this DLL in the same process. No client modifications, library-retention tricks, or unload suppression are included here. This limitation applies to your own Go plugins too. A DLL build and isolated callback verification are not proof that installation in the client is supported.

## Requirements and ABI

- Go 1.22 or later (the module's minimum); the Windows DLL was built and exercised with Go 1.27.1. Older toolchains have not been verified.
- Windows x64 only, using LLVM-MinGW 20261006 UCRT and its `x86_64-w64-mingw32-clang` compiler. Add the toolchain's `bin` directory to `PATH`. MSVC is not the cgo compiler used here.
- `CGO_ENABLED=1`, `GOOS=windows`, and `GOARCH=amd64`. Both Go source files and the example C translation unit have `windows && amd64 && cgo` build constraints; other configurations are excluded. The command below is a native Windows build, not a cross-compilation recipe.
- A host implementing the exact canonical header revision recorded in `../sdk-source.json`. There is no version/size field or automatic negotiation in `PluginApi`. Supplying a smaller or reordered table is unsafe.

`api.go` is the reusable `mandagram` package. It exposes one method for each of the 53 slots; `bridge.h` calls native function pointers because Go cannot call them directly. Its compile-time assertions require 64-bit pointers, a 53-pointer table, and the canonical 40-byte HookParam layout. The example is a separate `package main` under `examples/plain_text`.

## Build the actual plugin

Run from this repository root. No separate SDK installation or third-party Go modules are needed.

Windows PowerShell:

```powershell
$env:CGO_ENABLED = '1'
$env:GOOS = 'windows'
$env:GOARCH = 'amd64'
$env:CC = 'x86_64-w64-mingw32-clang'
New-Item -ItemType Directory -Force build | Out-Null
go build -buildmode=c-shared -o build/plain_text.dll ./examples/plain_text
```

This command produces the actual Windows x64 DLL and a generated cgo header beside it. The generated header describes internal Go bridge exports; hosts should use `include/mandagram_plugin.h`, not depend on those implementation exports. Building the SDK package alone does not produce a plugin. Runtime C library dependencies depend on the selected toolchain; inspect the resulting DLL before distributing it.

The required host entrypoints are defined in `plugin.c` with canonical C signatures:

```c
const PluginMetadata* mandagram_plugin_metadata(void);
bool mandagram_plugin_init(const PluginApi* api);
```

Metadata is an immutable C static struct referencing C string literals, valid while the library remains loaded. Init rejects a null table or missing message-send registration slot, then enters Go and uses `API.RegisterMessageSendHook`; registration status zero means success. The registered pointer is a C trampoline, not a Go function pointer. Other cgo-generated exports are implementation details.

## Use the SDK from another module

The module import is `github.com/GrapeS-T/MandaGramSDKGo`. Until a corresponding remote revision is published, use a local replacement rather than assuming `go get` can fetch this checkout. In your own module:

```powershell
go mod edit -require=github.com/GrapeS-T/MandaGramSDKGo@v0.0.0
go mod edit -replace=github.com/GrapeS-T/MandaGramSDKGo=C:/Projects/MandaGramSDKs/mandagram-sdk-go
```

Replace the local checkout path as needed. Import the package with the alias `mandagram`. Start your native plugin package from the three files in `examples/plain_text`, preserving their Windows amd64 cgo build constraints, change metadata and callbacks, and change `plugin.h`'s relative canonical-header include to match your layout (or configure a C include directory). Keep the SDK checkout and canonical header together. Build your `package main` as a Windows x64 DLL with `-buildmode=c-shared` and the environment above.

Use `mandagram.FromPointer` on the native API-table address passed to your init callback. The resulting API borrows that table and does not copy, validate, or own it. Each method requires that its slot exists and is non-null. Check the slots you need in your C init wrapper before calling into Go. Nullability, allowed calling threads, JNI references, and argument retention follow the host implementation, not just the header's types.

## Ownership and callbacks

All strings passed through SDK methods are pointers to NUL-terminated C memory, not Go strings. Output buffers are caller-owned; allocate C memory when storage must survive a call, and free it only after the host no longer needs it. `Object` wraps an opaque host-owned `jobject`; it is not a Go object, an owning JNI global reference, or a stable object identity. `HookParam` is a view over native memory; do not put Go pointers in it.

Native callback types represent C function addresses. Define correctly typed C trampolines that call `//export` Go functions; never cast a Go function value to a native callback address. Retained callbacks must outlive registration. Do not pass Go heap pointers to C for retention, and do not hide Go pointers in `uintptr` user data. Integer IDs or `runtime/cgo.Handle` values may represent Go state if you implement a lifetime that lasts through every callback; the current ABI provides no general unregistration mechanism. The plain-text example uses zero user data and no retained Go state.

The example's actual prefix decision runs in `goPlainBeforeSend`, not in the C trampoline. It compares at most seven bytes against `.plain `, stopping on the first mismatch (including an early NUL). A match returns the original C pointer advanced by seven bytes; otherwise it returns the original pointer. A null input returns null. Exact `.plain ` returns a pointer to its terminator. The host must copy the returned substring before releasing the original input. The callback does not mutate, copy, allocate, retain, or free text, and no Go pointer is returned or retained by C. It is stateless and safe for concurrent callbacks on valid independent host buffers. This describes the callback's code, not a guarantee that the Go runtime/cgo transition performs no internal allocation or runtime initialization.

## Verification boundary

The integration owner built the real Windows x64 DLL with Go 1.27.1 and LLVM-MinGW 20261006 UCRT (`x86_64-w64-mingw32-clang`) and loaded/called it from a native C++ harness. The checks passed for the exact required `mandagram_plugin_metadata` and `mandagram_plugin_init` export names, stable metadata, null API and missing registration-slot rejection, failed registration, 11 text-boundary/null-text cases, borrowed-pointer identity/offset, and a callback from a foreign native thread. The harness did not unload the DLL. These are isolated native DLL results, not a successful installation or lifecycle test in MandaGram.

Source review compared all 53 Go methods and C bridge wrappers with the canonical header: slot selection, argument order, signed/unsigned 32- and 64-bit integers, pointer-sized user data, C booleans, 32-bit floats, opaque objects, native callback addresses, and the 40-byte HookParam layout. No mapping defects were found. The native harness exercises the example's message-send registration/callback path, not the runtime semantics of the other API slots.

For subsequent changes, rebuild the Windows DLL and repeat the export, metadata, init, registration, boundary/pointer, and foreign-thread checks in a fresh native harness. Supply a full canonical table, keep API and text storage alive through their calls, and terminate the process without unloading. Neither these checks nor compilation establish normal client lifecycle compatibility.
