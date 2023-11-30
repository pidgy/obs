package properties

type (
	// Type wraps obs_data_t.
	Type uintptr
)

const (
	// Null represents a nil obs_data_t.
	Null = Type(0)
)

// IsNull returns true or false as to whether or not Type has been initialized.
func (t Type) IsNull() bool {
	return t == Null
}
