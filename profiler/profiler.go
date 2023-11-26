package profiler

import (
	"syscall"
	"unsafe"

	"github.com/pkg/errors"

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
	r, _, err := dll.OBS.NewProc("profiler_name_store_create").Call()
	if err != syscall.Errno(0) {
		return 0, errors.Wrap(err, "profiler_name_store_create")
	}

	return NameStore(r), nil
}

// Get wraps obs_get_profiler_name_store.
func Get() (NameStore, error) {
	r, _, err := dll.OBS.NewProc("obs_get_profiler_name_store").Call()
	if err != syscall.Errno(0) {
		return 0, errors.Wrap(err, "obs_get_profiler_name_store")
	}

	return NameStore(r), nil
}

// Close wraps profiler_name_store_free.
func (n NameStore) Close() error {
	_, _, err := dll.OBS.NewProc("profiler_name_store_free").Call(
		uintptr(unsafe.Pointer(n)),
	)
	if err != syscall.Errno(0) {
		return errors.Wrap(err, "profiler_name_store_free")
	}

	return nil
}

// IsNull returns true or false as to whether or not NameStore has been initialized.
func (n NameStore) IsNull() bool {
	return n == Null
}

// Store wraps const char *profile_store_name(profiler_name_store_t *store, const char *format, ...).
func (n NameStore) Store(name string) (string, error) {
	r, _, err := dll.OBS.NewProc("profile_store_name").Call(
		uintptr(unsafe.Pointer(n)),
		uptr.FromString(name),
	)
	if err != syscall.Errno(0) {
		return "", errors.Wrap(err, "profile_store_name")
	}

	return uptr.String(r), nil
}
