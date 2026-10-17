# MandaGram Go SDK

Reusable Go bindings for all 53 native MandaGram PluginApi slots, with a Go plain-text plugin example that builds into a Windows x64 DLL. Windows amd64 with cgo is the only supported target. This is a native plugin SDK, not a Telegram network API.

**Lifecycle limitation:** Go `c-shared` libraries cannot safely be unloaded. The current client's metadata-probe and disable/unload lifecycle is incompatible with this embedded runtime. Do not install this example in a normal client session; a successful isolated load/call is not end-to-end client compatibility.

See [building and consuming the SDK](docs/building.md) for requirements, commands, memory ownership, and safe verification boundaries. The [canonical header](include/mandagram_plugin.h) and [source manifest](sdk-source.json) define the pinned ABI.
