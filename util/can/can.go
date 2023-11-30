// Package must provides convenience functions for callers that can safely assume no errors will be encountered.
package can

// Releaser represents a contract for types that have a Release method.
type Releaser interface {
	Release() error
}

func B(fn func() (bool, error)) bool {
	b, err := fn()
	panics(err)
	return b
}

// I executes fn and panics if an error is encountered.
func I(fn func() (int, error)) int {
	i, err := fn()
	panics(err)
	return i
}

// Release calls a Releasers Release method and panics if an error is encountered.
func Release(r Releaser) {
	panics(r.Release())
}

// S executes fn and panics if an error is encountered.
func S(fn func() (string, error)) string {
	s, err := fn()
	panics(err)
	return s
}

// Panic executes fn and panics if an error is encountered.
func Panic(fn func() error) {
	panics(fn())
}

// panics if err is a valid type.
func panics(err error) {
	if err != nil {
		panic(err)
	}
}
