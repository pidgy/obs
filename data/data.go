package data

import (
	"syscall"

	"github.com/pkg/errors"

	"github.com/pidgy/obs/dll"
	"github.com/pidgy/obs/uptr"
)

type (
	// Type wraps obs_data_t.
	Type uintptr
)

const (
	// Null represents a Null obs_data_t.
	Null = Type(0)
)

// New wraps obs_data_create.
func New() (Type, error) {
	r, _, err := dll.OBS.NewProc("obs_data_create").Call()
	if err != syscall.Errno(0) {
		return Type(0), errors.Wrap(err, "obs_data_create")
	}
	return Type(r), nil
}

// IsNull returns true or false as to whether or not Type has been initialized.
func (t Type) IsNull() bool {
	return t == Null
}

// Release wraps void obs_data_release(obs_data_t *data).
func (t Type) Release() error {
	_, _, err := dll.OBS.NewProc("obs_data_release").Call()
	if err != syscall.Errno(0) {
		return errors.Wrap(err, "obs_data_release")
	}
	return nil
}

// SaveJSON wraps bool obs_data_save_json(obs_data_t *data, const char *file).
func (t Type) SaveJSON(file string) error {
	_, _, err := dll.OBS.NewProc("obs_data_save_json").Call(
		uintptr(t),
		uptr.FromString(file),
	)
	if err != syscall.Errno(0) {
		return errors.Wrap(err, "obs_data_save_json")
	}
	return nil
}

// SetString wraps void obs_data_set_string(obs_data_t *data, const char *name,const char *val).
func (t Type) SetString(name, val string) error {
	_, _, err := dll.OBS.NewProc("obs_data_set_string").Call(
		uintptr(t),
		uptr.FromString(name),
		uptr.FromString(val),
	)
	if err != syscall.Errno(0) {
		return errors.Wrap(err, "obs_data_set_string")
	}
	return nil
}
