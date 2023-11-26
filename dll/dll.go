// Package dll provides access to libobs .dll files.
//
// By default, dll will try and add multiple directories starting from "$PWD\libobs\bin\64bit" to the .dll search path.
package dll

import (
	join "errors"
	"fmt"
	"os"
	"path/filepath"
	"syscall"

	"github.com/pkg/errors"

	"github.com/pidgy/obs/kernel32"
)

// OBS errorss obs.dll access in a single package.
var (
	OBS = syscall.NewLazyDLL(`obs.dll`)
)

const (
	pwd = "./libobs" // "./libobs-30.0.0/"
	sys = ""         // Don't use "C:/Program Files/obs-studio/"
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
func Cleanup() error {
	var errs []error

	for _, c := range cookies {
		err := c.RemoveDllDirectory()
		if err != nil {
			errs = append(errs, err)
		}
	}

	return join.Join(errs...)
}

// Core loads 64bit bin .dlls from "${bin}\libobs\bin...}".
func Core(dll string) (file, dir string, err error) {
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
