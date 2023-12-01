package call

// #cgo windows CFLAGS: -I ../libobs/include -I libobs/include -I libobs
// #include "obs.h"
import "C"
import (
	"unsafe"

	"github.com/pidgy/obs/dll"
	"github.com/pidgy/obs/uptr"
)

type (
	// Data wraps calldata_t.
	Data C.calldata_t
)

// New wraps void calldata_init(calldata_t *data).
func New() (*Data, error) {
	d := &Data{}
	return d, dll.OBS("calldata_init", uintptr(unsafe.Pointer(d)))

}

// Bool wraps void calldata_set_bool(calldata_t *data, const char *name, bool val).
func (d *Data) Bool(name string, val bool) error {
	return dll.OBS("calldata_set_bool", uintptr(unsafe.Pointer(d)), uptr.FromString(name), uptr.FromBool(val))
}

// Free wraps void calldata_free(calldata_t *data).
func (d *Data) Free() error { return dll.OBS("calldata_free", uintptr(unsafe.Pointer(d))) }

// IsNull returns true or false as to whether or not Data has been initialized.
func (d *Data) IsNull() bool {
	return d == nil
}
