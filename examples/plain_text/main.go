//go:build windows && amd64 && cgo

package main

/*
#cgo CFLAGS: -std=c11
#include "plugin.h"
*/
import "C"

import (
	mandagram "github.com/GrapeS-T/MandaGramSDKGo"
	"unsafe"
)

//export goPlainInit
func goPlainInit(ptr unsafe.Pointer) C.int {
	api := mandagram.FromPointer(ptr)
	if api.RegisterMessageSendHook(mandagram.MessageBeforeSendFn(C.plain_before_send_address()), nil, 0) != 0 {
		return 0
	}
	return 1
}

//export goPlainBeforeSend
func goPlainBeforeSend(text *C.char) *C.char {
	if text == nil {
		return nil
	}
	const prefix = ".plain "
	for i := range len(prefix) {
		// Compare one byte at a time: a short NUL-terminated string must
		// never be read beyond its terminator. No string conversion or copy.
		if *(*byte)(unsafe.Add(unsafe.Pointer(text), i)) != prefix[i] {
			return text
		}
	}
	// The host copies this borrowed substring before releasing the input.
	return (*C.char)(unsafe.Add(unsafe.Pointer(text), len(prefix)))
}

func main() {}
