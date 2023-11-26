package kernel32

import (
	"syscall"
	"unsafe"

	"github.com/pkg/errors"
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
		return 0, errors.Wrap(err, "AddDllDirectory")
	}
	if r == 0 {
		return 0, windows.GetLastError()
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
		return errors.Wrap(err, "RemoveDllDirectory")
	}
	if r == 0 {
		return windows.GetLastError()
	}

	return nil
}

func (c Cookie) LoadLibrary(dll string) error {
	r, err := windows.LoadLibrary(dll)
	if err != nil {
		return errors.Wrap(err, "LoadLibrary")
	}
	if r == 0 {
		return windows.GetLastError()
	}

	return nil
}

func (c Cookie) LoadLibraryEx(dll string) error {
	r, err := windows.LoadLibraryEx(
		dll,
		windows.Handle(0),
		windows.LOAD_WITH_ALTERED_SEARCH_PATH,
	)
	if err != nil {
		return errors.Wrap(err, "LoadLibraryEx")
	}
	if r == 0 {
		return windows.GetLastError()
	}

	return nil
}
