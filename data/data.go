package data

import (
	"syscall"

	"github.com/pidgy/obs/lib"
	"github.com/pidgy/obs/uptr"
)

type (
	// Type wraps obs_data_t.
	Type uintptr
)

const (
	// NULL represents a NULL obs_data_t.
	NULL = Type(0)
)

// New wraps obs_data_create.
func New() (Type, error) {
	r, _, err := lib.OBS.NewProc("obs_data_create").Call()
	if err != syscall.Errno(0) {
		return Type(0), err
	}

	return Type(r), nil
}

// SetString wraps void obs_data_set_string(obs_data_t *data, const char *name,const char *val).
func (t Type) SetString(name, val string) error {
	_, _, err := lib.OBS.NewProc("obs_data_set_string").Call(
		uintptr(t),
		uptr.FromString(name),
		uptr.FromString(val),
	)
	if err != syscall.Errno(0) {
		return err
	}

	return nil
}
