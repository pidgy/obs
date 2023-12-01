package profiler

import (
	"unsafe"

	"github.com/pidgy/obs/dll"
	"github.com/pidgy/obs/uptr"
)

// NameStore wraps profiler_name_store_t.
type NameStore uintptr

const (
	// Null wraps the Null value for *profiler_name_store_t.
	Null = NameStore(0)
)

// New wraps profiler_name_store_create.
func New() (NameStore, error) {
	r, err := dll.OBSuintptr("profiler_name_store_create")
	return NameStore(r), err
}

// Get wraps obs_get_profiler_name_store.
func Get() (NameStore, error) {
	r, err := dll.OBSuintptr("obs_get_profiler_name_store")
	return NameStore(r), err
}

// Close wraps profiler_name_store_free.
func (n NameStore) Close() error {
	return dll.OBS("profiler_name_store_free", uintptr(unsafe.Pointer(n)))
}

// IsNull returns true or false as to whether or not NameStore has been initialized.
func (n NameStore) IsNull() bool {
	return n == Null
}

// Store wraps const char *profile_store_name(profiler_name_store_t *store, const char *format, ...).
func (n NameStore) Store(name string) (string, error) {
	return dll.OBSstring("profile_store_name", uintptr(unsafe.Pointer(n)), uptr.FromString(name))
}
