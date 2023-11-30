package call

import (
	"syscall"

	"github.com/pidgy/obs/dll"
	"github.com/pidgy/obs/uptr"
	"github.com/pkg/errors"
)

type (
	// Data wraps calldata_t.
	Data uintptr
)

const (
	// Null represents a nil calldata_t pointer.
	Null = Data(0)
)

// New wraps void calldata_init(calldata_t *data).
func New() (Data, error) {
	d := Null

	_, _, err := dll.OBS.NewProc("calldata_init").Call(
		uintptr(d),
	)
	if err != syscall.Errno(0) {
		return Null, errors.Wrap(err, "calldata_init")
	}

	return d, nil

}

// Bool wraps void calldata_set_bool(calldata_t *data, const char *name, bool val).
func (d Data) Bool(name string, val bool) error {
	_, _, err := dll.OBS.NewProc("calldata_set_bool").Call(
		uintptr(d),
		uptr.FromString(name),
		uptr.FromBool(val),
	)
	if err != syscall.Errno(0) {
		return errors.Wrap(err, "calldata_set_bool")
	}
	return nil
}

// Free wraps void calldata_free(calldata_t *data).
func (d Data) Free() error {
	if d.IsNull() {
		return nil
	}

	_, _, err := dll.OBS.NewProc("calldata_free").Call(
		uintptr(d),
	)
	if err != syscall.Errno(0) {
		return errors.Wrap(err, "calldata_free")
	}
	return nil
}

// IsNull returns true or false as to whether or not Data has been initialized.
func (d Data) IsNull() bool {
	return d == Null
}
