package core

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/pkg/errors"

	"github.com/pidgy/obs/dll"
	"github.com/pidgy/obs/locale"
	"github.com/pidgy/obs/profiler"
	"github.com/pidgy/obs/uptr"
)

// Shutdown wraps obs_shutdown.
func Shutdown() error {
	_, _, err := dll.OBS.NewProc("obs_shutdown").Call()
	if err != syscall.Errno(0) {
		return errors.Wrap(err, "obs_shutdown")
	}

	return dll.Cleanup()
}

// Startup wraps obs_startup, use profiler.None as NULL value.
func Startup(locale locale.Type, moduleConfigPath string, ns profiler.NameStore) error {
	file, _, err := dll.Core("obs.dll")
	if err != nil {
		return err
	}

	println("loaded obs.dll from", file)

	_, _, err = dll.OBS.NewProc("obs_startup").Call(
		uptr.FromString(locale.String()),
		uptr.FromString(moduleConfigPath),
		uintptr(unsafe.Pointer(ns)),
	)
	if err != syscall.Errno(0) {
		return errors.Wrap(err, "obs_startup")
	}

	return nil
}

// Locale wraps obs_get_locale.
func Locale() (locale.Type, error) {
	l, _, err := dll.OBS.NewProc("obs_get_locale").Call()
	if err != syscall.Errno(0) {
		return "", errors.Wrap(err, "obs_get_locale")
	}
	if l == 0 {
		return "", fmt.Errorf("unknown locale")
	}

	return locale.Type(uptr.String(l)), nil
}

// SetLocale wraps void obs_set_locale(const char *locale).
func SetLocale(l locale.Type) error {
	_, _, err := dll.OBS.NewProc("obs_set_locale").Call(
		uptr.FromString(l.String()),
	)
	if err != syscall.Errno(0) {
		return errors.Wrap(err, "obs_set_locale")
	}
	return nil
}

// Version wraps obs_get_version_string and obs_get_version.
func Version() (string, uint32, error) {
	v1, _, err := dll.OBS.NewProc("obs_get_version_string").Call()
	if err != syscall.Errno(0) {
		return "", 0, errors.Wrap(err, "obs_get_version_string")
	}

	v2, _, err := dll.OBS.NewProc("obs_get_version").Call()
	if err != syscall.Errno(0) {
		return "", 0, errors.Wrap(err, "obs_get_version")
	}

	return uptr.String(v1), uint32(v2), nil
}
