package core

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/pidgy/obs/lib"
	"github.com/pidgy/obs/locale"
	"github.com/pidgy/obs/profiler"
	"github.com/pidgy/obs/uptr"
)

// Shutdown wraps obs_shutdown.
func Shutdown() error {
	_, _, err := lib.OBS.NewProc("obs_shutdown").Call()
	if err != syscall.Errno(0) {
		return err
	}
	return nil
}

// Startup wraps obs_startup, use profiler.None as NULL value.
func Startup(locale locale.Type, moduleConfigPath string, n profiler.NameStore) error {
	_, _, err := lib.OBS.NewProc("obs_startup").Call(
		uptr.FromString(locale.String()),
		uptr.FromString(moduleConfigPath),
		uintptr(unsafe.Pointer(n)),
	)
	if err != syscall.Errno(0) {
		return err
	}

	return nil
}

// Locale wraps obs_get_locale.
func Locale() (locale.Type, error) {
	l, _, err := lib.OBS.NewProc("obs_get_locale").Call()
	if err != syscall.Errno(0) {
		return "", err
	}
	if l == 0 {
		return "", fmt.Errorf("unknown locale")
	}

	return locale.Type(uptr.String(l)), nil
}

// Version wraps obs_get_version_string and obs_get_version.
func Version() (string, uint32, error) {
	v1, _, err := lib.OBS.NewProc("obs_get_version_string").Call()
	if err != syscall.Errno(0) {
		return "", 0, err
	}

	v2, _, err := lib.OBS.NewProc("obs_get_version").Call()
	if err != syscall.Errno(0) {
		return "", 0, err
	}

	return uptr.String(v1), uint32(v2), nil
}
