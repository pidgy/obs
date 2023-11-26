package uptr

import (
	"fmt"
	"math"
	"unsafe"

	"golang.org/x/sys/windows"
)

const Null = uintptr(0)

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

// IsNull returns true or false as to whether or not Type has been initialized.
func IsNull(u uintptr) bool {
	return u == Null
}

// NewBytePtr creates amd converts a a *byte to bool.
func NewBytePtr(size int) *byte {
	b := make([]byte, size)
	return &b[0]
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

// Version converts an unsigned 32-bit integer into a version string.
// This implementation is incorrect... yet to figure it out!
// 27.5.32
// 11011000000100000000000000100
// 30.0.0
// 11110000000000000000000000000
func Version(u uint32) string {
	major := (u & 0b11111111000000000000000000000000) >> 24
	minor := (u & 0b000000000000000001111111100000000) << 5
	patch := (u & 0b00000000000000000000000011111111) << 3
	return fmt.Sprintf("%d.%d.%d", major, minor, patch)
}
