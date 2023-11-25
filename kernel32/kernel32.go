package kernel32

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

// Cookie wraps the return value from AddDllDirectory.
type Cookie uintptr

var kernel32 = windows.NewLazySystemDLL("kernel32.dll")

// AddDllDirectory is borrowed from "golang.org/x/sys/windows".
func AddDllDirectory(path string) (Cookie, error) {
	add := kernel32.NewProc("AddDllDirectory")

	err := add.Find()
	if err != nil {
		return 0, err
	}

	p := windows.StringToUTF16Ptr(path)

	r, _, err := syscall.Syscall(add.Addr(), 1, uintptr(unsafe.Pointer(p)), 0, 0)
	if err != syscall.Errno(0) {
		return 0, err
	}
	if r == 0 {
		return 0, fmt.Errorf("failed to add dll directory: %s", path)
	}

	return Cookie(r), nil
}

// RemoveDllDirectory is borrowed from "golang.org/x/sys/windows".
func (c Cookie) RemoveDllDirectory() error {
	remove := kernel32.NewProc("RemoveDllDirectory")

	err := remove.Find()
	if err != nil {
		return err
	}

	r, _, err := syscall.Syscall(remove.Addr(), 1, uintptr(c), 0, 0)
	if err != syscall.Errno(0) {
		return err
	}
	if r == 0 {
		return fmt.Errorf("failed to remove dll directory: %d", c)
	}

	return nil
}
