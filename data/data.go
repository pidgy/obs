package data

import (
	"encoding/json"
	"syscall"

	"github.com/pkg/errors"

	"github.com/pidgy/obs/array"
	"github.com/pidgy/obs/dll"
	"github.com/pidgy/obs/uptr"
)

type (
	// Type wraps obs_data_t.
	Type uintptr

	Setter interface {
		Arrays() map[string]array.Type
		Bools() map[string]bool
		Doubles() map[string]float64
		Ints() map[string]int
		Objects() map[string]Type
		Strings() map[string]string
	}
)

const (
	// Null represents a nil obs_data_t.
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

// String wraps const char *obs_data_get_json(obs_data_t *data).
func (t Type) String() string {
	r, _, err := dll.OBS.NewProc("obs_data_get_json").Call(
		uintptr(t),
	)
	if err != syscall.Errno(0) {
		return errors.Wrap(err, "obs_data_get_json").Error()
	}
	return uptr.String(r)
}

// MustRelease wraps void obs_data_release(obs_data_t *data).
func (t Type) MustRelease() {
	_, _, err := dll.OBS.NewProc("obs_data_release").Call()
	if err != syscall.Errno(0) {
		panic(errors.Wrap(err, "obs_data_release"))
	}
}

// Pretty wraps const char *obs_data_get_json(obs_data_t *data).
func (t Type) Pretty() string {
	r, _, err := dll.OBS.NewProc("obs_data_get_json").Call(
		uintptr(t),
	)
	if err != syscall.Errno(0) {
		return errors.Wrap(err, "obs_data_get_json").Error()
	}

	s := json.RawMessage(uptr.String(r))
	pretty := map[string]interface{}{}
	err = json.Unmarshal(s, &pretty)
	if err != nil {
		return errors.Wrap(err, "obs_data_get_json: unmarshal").Error()
	}

	raw, err := json.MarshalIndent(pretty, "    ", "")
	if err != nil {
		return errors.Wrap(err, "obs_data_get_json: marshal").Error()
	}

	return string(raw)
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

// Set wraps obs_data_set_string/obs_data_set_int/obs_data_set_double/obs_data_set_bool/obs_data_set_obj/obs_data_set_array.
func (t Type) Set(s Setter) error {
	err := t.SetStrings(s.Strings())
	if err != nil {
		return err
	}

	err = t.SetBools(s.Bools())
	if err != nil {
		return err
	}

	err = t.SetInts(s.Ints())
	if err != nil {
		return err
	}

	return nil
}

// SetBool wraps void obs_data_set_bool(obs_data_t *data, const char *name, bool val).
func (t Type) SetBool(name string, val bool) error {
	_, _, err := dll.OBS.NewProc("obs_data_set_bool").Call(
		uintptr(t),
		uptr.FromString(name),
		uptr.FromBool(val),
	)
	if err != syscall.Errno(0) {
		return errors.Wrap(err, "obs_data_set_bool")
	}
	return nil
}

// SetBools wraps void obs_data_set_bool(obs_data_t *data, const char *name, bool val).
func (t Type) SetBools(m map[string]bool) error {
	for k, v := range m {
		_, _, err := dll.OBS.NewProc("obs_data_set_bool").Call(
			uintptr(t),
			uptr.FromString(k),
			uptr.FromBool(v),
		)
		if err != syscall.Errno(0) {
			return errors.Wrap(err, "obs_data_set_bool")
		}
	}
	return nil
}

// SetString wraps void void obs_data_set_int(obs_data_t *data, const char *name, long long val).
func (t Type) SetInt(name string, val int) error {
	_, _, err := dll.OBS.NewProc("obs_data_set_int").Call(
		uintptr(t),
		uptr.FromString(name),
		uintptr(val),
	)
	if err != syscall.Errno(0) {
		return errors.Wrap(err, "obs_data_set_int")
	}
	return nil
}

// SetStrings wraps void obs_data_set_int(obs_data_t *data, const char *name, long long val).
func (t Type) SetInts(m map[string]int) error {
	for k, v := range m {
		err := t.SetInt(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

// SetString wraps void obs_data_set_string(obs_data_t *data, const char *name, const char *val).
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

// SetStrings wraps void obs_data_set_string(obs_data_t *data, const char *name, const char *val).
func (t Type) SetStrings(m map[string]string) error {
	for k, v := range m {
		err := t.SetString(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
