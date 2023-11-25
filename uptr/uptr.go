package uptr

import (
	"math"
	"unsafe"

	"golang.org/x/sys/windows"
)

// Bool converts a uintptr to bool.
func Bool(u uintptr) bool {
	return bool(*(*bool)(unsafe.Pointer(&u)))
}

// BytePtrToString converts a *byte to string.
func BytePtrToString(b *byte) string {
	return windows.BytePtrToString((*byte)(unsafe.Pointer(b)))
}

// Float converts a uintptr to float32.
func Float(u uintptr) float32 {
	return math.Float32frombits(uint32(u))
}

// FromBool converts a bool to uintptr.
func FromBool(b bool) uintptr {
	if b {
		return uintptr(1)
	}
	return uintptr(0)
}

// FromFloat converts a float32 to uintptr.
func FromFloat(f float32) uintptr {
	return uintptr(math.Float32bits(f))
}

// FromString converts a string to uintptr.
func FromString(s string) uintptr {
	b, err := windows.BytePtrFromString(s)
	if err != nil {
		panic(err)
	}
	return uintptr(unsafe.Pointer(b))
}

// Int converts a uintptr to int.
func Int(u uintptr) int {
	return int(*(*int)(unsafe.Pointer(&u)))
}

// NewBytePtr creates amd converts a a *byte to bool.
func NewBytePtr() uintptr {
	b := make([]byte, 4096)
	return uintptr(unsafe.Pointer(&b[0]))
}

// FromString converts a string to uintptr.
func ReferenceFromString(s string) uintptr {
	b, err := windows.BytePtrFromString(s)
	if err != nil {
		panic(err)
	}
	return uintptr(unsafe.Pointer(&b))
}

// String converts a uintptr to string.
func String(u uintptr) string {
	return BytePtrToString((*byte)(unsafe.Pointer(u)))
}
