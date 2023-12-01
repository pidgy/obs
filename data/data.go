package data

import (
	"encoding/json"
	"fmt"

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
	r, err := dll.OBSuintptr("obs_data_create")
	return Type(r), err
}

// IsNull returns true or false as to whether or not Type has been initialized.
func (t Type) IsNull() bool {
	return t == Null
}

// String wraps const char *obs_data_get_json(obs_data_t *data).
func (t Type) String() string {
	s, _ := dll.OBSstring("obs_data_get_json", uintptr(t))
	return s
}

// Pretty wraps const char *obs_data_get_json(obs_data_t *data).
func (t Type) Pretty() string {
	s, err := dll.OBSstring("obs_data_release", uintptr(t))
	if err != nil {
		return fmt.Sprintf(`{"error": "%v"}`, err)
	}

	pretty := map[string]interface{}{}

	err = json.Unmarshal(json.RawMessage(s), &pretty)
	if err != nil {
		return fmt.Sprintf(`{"error": "%v"}`, err)
	}

	raw, err := json.MarshalIndent(pretty, "    ", "")
	if err != nil {
		return fmt.Sprintf(`{"error": "%v"}`, err)
	}

	return string(raw)
}

// Release wraps void obs_data_release(obs_data_t *data).
func (t Type) Release() error {
	return dll.OBS("obs_data_release")
}

// SaveJSON wraps bool obs_data_save_json(obs_data_t *data, const char *file).
func (t Type) SaveJSON(file string) error {
	return dll.OBS("obs_data_save_json", uintptr(t), uptr.FromString(file))
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
	return dll.OBS("obs_data_set_bool", uintptr(t), uptr.FromString(name), uptr.FromBool(val))
}

// SetBools wraps void obs_data_set_bool(obs_data_t *data, const char *name, bool val).
func (t Type) SetBools(m map[string]bool) error {
	for k, v := range m {
		err := t.SetBool(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

// SetString wraps void void obs_data_set_int(obs_data_t *data, const char *name, long long val).
func (t Type) SetInt(name string, val int) error {
	return dll.OBS("obs_data_set_int", uintptr(t), uptr.FromString(name), uintptr(val))
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
	return dll.OBS("obs_data_set_string", uintptr(t), uptr.FromString(name), uptr.FromString(val))
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
