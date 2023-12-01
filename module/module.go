package module

// #cgo windows CFLAGS: -I ../libobs/include -I libobs/include -I libobs
//
// #include "obs-module.h"
import "C"

import (
	"fmt"
	"path/filepath"
	"syscall"
	"unsafe"

	"github.com/pkg/errors"

	"github.com/pidgy/obs/dll"
	"github.com/pidgy/obs/uptr"
)

// Module provides an interface to opaquely use obs-plugin modules across packages.
type (
	// FailureInfo wraps obs_module_failure_info.
	FailureInfo uintptr

	// Type wraps obs_module_t.
	Type uintptr

	Return int
)

var (
	ErrNotImplemented = fmt.Errorf("module does not implement this procedure")
)

const (
	Null = Type(0)

	Success             Return = -iota // - Successful
	Error                              // - A generic error occurred
	FileNotFound                       // - The module was not found
	MissingExports                     // - Required exports are missing
	IncompatibleVersion                // - Incompatible version
	HardcodedSkip                      // - Skipped by hardcoded rules (e.g. obsolete obs-browser macOS plugin)
)

// Current wraps obs_module_t *obs_current_module(void).
func Current() (Type, error) {
	r, err := dll.OBSuintptr("obs_current_module")
	return Type(r), err
}

// New wraps
// - int obs_open_module(obs_module_t **module, const char *path, const char *data_path).
// - bool obs_init_module(obs_module_t *module)
func New(name string) (Type, error) {
	m := Type(Null)

	file, dir, err := dll.Module(name)
	if err != nil {
		return Null, err
	}
	dir = filepath.Join(dir, "../", "../", "data", "obs-plugins", name)

	err = dll.OBS("obs_open_module", uintptr(unsafe.Pointer(&m)), uptr.FromString(file), uptr.FromString(dir))
	if err != nil {
		return Null, errors.Wrap(err, Return(err.(syscall.Errno)).String())
	}

	ok, err := dll.OBSbool("obs_init_module", uintptr(m))
	if err != nil {
		return Null, err
	}
	if !ok {
		return Null, errors.Wrap(errors.Errorf("module was not loaded successfully"), "obs_init_module")
	}

	return m, nil
}

// IsNull returns true or false as to whether or not Type has been initialized.
func (f FailureInfo) IsNull() bool {
	return f == FailureInfo(Null)
}

// LoadAll wraps void obs_load_all_modules(void).
func LoadAll() error {
	return dll.OBS("obs_load_all_modules")
}

// LoadAll2 wraps void obs_load_all_modules2(struct obs_module_failure_info *mfi).
func LoadAll2() (FailureInfo, error) {
	f := FailureInfo(Null)
	return f, dll.OBS("obs_load_all_modules2", uintptr(f))
}

// Log wraps obs_log_loaded_modules.
func Log() error {
	return dll.OBS("obs_log_loaded_modules")
}

// String returns the human-readable representation of the return value obs_open_module.
func (r Return) String() string {
	switch r {
	case Success:
		return "successful"
	case Error:
		return "generic error occurred"
	case FileNotFound:
		return "the module was not found"
	case MissingExports:
		return "required exports are missing"
	case IncompatibleVersion:
		return "incompatible version"
	case HardcodedSkip:
		return "skipped by hardcoded rules"
	default:
		return fmt.Sprintf("unknown return value: %d", r)
	}
}

// Description wraps const char *obs_get_module_description(obs_module_t *module).
func (t Type) Description() (string, error) {
	return dll.OBSstring("obs_get_module_description", uintptr(t))
}

// IsNull returns true or false as to whether or not Type has been initialized.
func (t Type) IsNull() bool {
	return t == Null
}

// Name wraps const char *obs_get_module_name(obs_module_t *module).
func (t Type) Name() (string, error) {
	return dll.OBSstring("obs_get_module_name", uintptr(t))
}
