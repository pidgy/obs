package properties

import "github.com/pidgy/obs/dll"

type (
	// Type wraps obs_data_t.
	Type uintptr
)

const (
	// Null represents a nil obs_data_t.
	Null = Type(0)
)

// Destroy wraps void obs_properties_destroy(obs_properties_t *props).
func (t Type) Destroy() error {
	return dll.OBS("obs_properties_destroy", uintptr(t))
}

// IsNull returns true or false as to whether or not Type has been initialized.
func (t Type) IsNull() bool {
	return t == Null
}
