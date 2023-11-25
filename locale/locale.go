package locale

// Type wraps locale parameter values.
type Type string

const (
	// EnUS wraps the "en-US" locale passed to obs.Startup.
	EnUS Type = "en-US"
)

func (l Type) String() string {
	return string(l)
}
