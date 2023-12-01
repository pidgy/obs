// Package dll provides access to libobs .dll files.
//
// By default, dll will try and add multiple directories starting from "$PWD\libobs\bin\64bit" to the .dll search path.
package dll

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/pkg/errors"

	"github.com/pidgy/obs/kernel32"
	"github.com/pidgy/obs/uptr"
)

// obs errorss obs.dll access in a single package.
var (
	obs = syscall.NewLazyDLL(`obs.dll`)
	// frontend = syscall.NewLazyDLL(`obs-frontend-api.dll`)
)

const (
	pwd = "./libobs"
	sys = ""
	// sys = "C:/Program Files/obs-studio/"
)

var (
	cores = []string{
		pwd + "/bin/64bit/",
		"../" + pwd + "/bin/64bit/",
		"../../" + pwd + "/bin/64bit/",
	}

	modules = []string{
		pwd + "/obs-plugins/64bit/",
		"../" + pwd + "/obs-plugins/64bit/",
		"../../" + pwd + "/obs-plugins/64bit/",
	}

	cookies []kernel32.Cookie
)

// Cleanup removes dll directory cookies and returns accumulated errors.
func Cleanup() {
	for _, c := range cookies {
		_ = c.RemoveDllDirectory()
	}
}

// Core loads 64bit bin .dlls from "${bin}\libobs\bin...}".
func Core(dll string) (file, dir string, err error) {
	if !strings.HasSuffix(dll, ".dll") {
		dll = fmt.Sprintf("%s.dll", dll)
	}

	dir = filepath.Join(sys, "/bin/64bit")
	file, err = load(dir, dll)
	if err == nil {
		return
	}

	abs, err := os.Getwd()
	if err != nil {
		return "", "", err
	}
	for _, c := range cores {
		dir = filepath.Join(abs, c)
		file, err = load(dir, dll)
		if err == nil {
			return
		}
	}

	return "", "", fmt.Errorf("failed to load %s", dll)
}

// File errorss a syscall.LazyDLL for use in other packages.
func File(name string) *syscall.LazyDLL {
	return syscall.NewLazyDLL(name)
}

// Module loads 64bit obs-plugin .dlls from "${plugins}\libobs\${obs-plugins...}".
func Module(dll string) (file, dir string, err error) {
	dir = filepath.Join(sys, "obs-plugins/64bit")

	file, err = load(dir, dll)
	if err == nil {
		return
	}

	abs, err := os.Getwd()
	if err != nil {
		return "", "", err
	}

	for _, m := range modules {
		dir = filepath.Join(abs, m)

		file, err = load(dir, dll)
		if err == nil {
			return
		}
	}

	return "", "", fmt.Errorf("failed to load %s", dll)
}

// OBS is a convenience function for executing procedure calls to obs.dll.
func OBS(name string, args ...uintptr) error {
	_, err := OBSuintptr(name, args...)
	return err
}

// OBSbool is a convenience function for executing procedure calls to obs.dll.
func OBSbool(name string, args ...uintptr) (bool, error) {
	r, err := OBSuintptr(name, args...)
	return uptr.Bool(r), err
}

// OBSfloat32 is a convenience function for executing procedure calls to obs.dll.
func OBSfloat32(name string, args ...uintptr) (float32, error) {
	r, err := OBSuintptr2(name, args...)
	return uptr.Float(r), err
}

// OBSint32 is a convenience function for executing procedure calls to obs.dll.
func OBSint32(name string, args ...uintptr) (int32, error) {
	r, err := OBSuintptr(name, args...)
	return int32(r), err
}

// OBSstring is a convenience function for executing procedure calls to obs.dll.
func OBSstring(name string, args ...uintptr) (string, error) {
	r, err := OBSuintptr(name, args...)
	return uptr.String(r), err
}

// OBSuint32 is a convenience function for executing procedure calls to obs.dll.
func OBSuint32(name string, args ...uintptr) (uint32, error) {
	r, err := OBSuintptr(name, args...)
	return uint32(r), err
}

// OBSuintptr is a convenience function for executing procedure calls to obs.dll.
func OBSuintptr(name string, args ...uintptr) (uintptr, error) {
	r, _, err := obs.NewProc(name).Call(args...)
	if err != syscall.Errno(0) {
		return 0, errors.Wrap(err, name)
	}
	return r, nil
}

// OBSuintptr2 is a convenience function for executing procedure calls to obs.dll.
func OBSuintptr2(name string, args ...uintptr) (uintptr, error) {
	_, r, err := obs.NewProc(name).Call(args...)
	if err != syscall.Errno(0) {
		return 0, errors.Wrap(err, name)
	}
	return r, nil
}

// load will add dir to the windows search path and return the absolute path of dll.
func load(dir string, dll string) (string, error) {
	cookie, err := kernel32.AddDllDirectory(dir)
	if err != nil {
		return "", errors.Wrap(err, dir)
	}
	cookies = append(cookies, cookie)

	file := filepath.Join(dir, dll)

	return file, errors.Wrap(cookie.LoadLibraryEx(file), file)
}
