package profiler

import (
	"syscall"
	"unsafe"

	"github.com/pidgy/obs/lib"
	"github.com/pidgy/obs/uptr"
)

// NameStore wraps profiler_name_store_t.
type NameStore uintptr

const (
	// NULL wraps the NULL value for *profiler_name_store_t.
	NULL = NameStore(0)
)

// New wraps profiler_name_store_create.
func New() (NameStore, error) {
	r, _, err := lib.OBS.NewProc("profiler_name_store_create").Call()
	if err != syscall.Errno(0) {
		return 0, err
	}

	return NameStore(r), nil
}

// Get wraps obs_get_profiler_name_store.
func Get() (NameStore, error) {
	r, _, err := lib.OBS.NewProc("obs_get_profiler_name_store").Call()
	if err != syscall.Errno(0) {
		return 0, err
	}

	return NameStore(r), nil
}

// Close wraps profiler_name_store_free.
func (n NameStore) Close() error {
	_, _, err := lib.OBS.NewProc("profiler_name_store_free").Call(
		uintptr(unsafe.Pointer(n)),
	)
	if err != syscall.Errno(0) {
		return err
	}

	return nil
}

// Store wraps const char *profile_store_name(profiler_name_store_t *store, const char *format, ...).
func (n NameStore) Store(name string) (string, error) {
	r, _, err := lib.OBS.NewProc("profile_store_name").Call(
		uintptr(unsafe.Pointer(n)),
		uptr.FromString(name),
	)
	if err != syscall.Errno(0) {
		return "", err
	}

	return uptr.String(r), nil
}
