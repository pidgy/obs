package output

import (
	"syscall"
	"unsafe"

	"github.com/pkg/errors"

	"github.com/pidgy/obs/dll"
	"github.com/pidgy/obs/uptr"
)

type (
	// Type wraps obs_output_info.
	Type uintptr
)

const (
	// Null represents a nil obs_output_info.
	Null = Type(0)
)

// EnumTypes wraps bool obs_enum_output_types(size_t idx, const char **id).
func EnumTypes() (ids []string, err error) {
	for idx := uintptr(0); idx < 1024; idx++ {
		id := uptr.NewBytePtr(4096)

		r, _, err := dll.OBS.NewProc("obs_enum_output_types").Call(
			idx,
			uintptr(unsafe.Pointer(&id)),
		)
		if err != syscall.Errno(0) {
			return nil, errors.Wrap(err, "obs_enum_output_types")
		}

		if !uptr.Bool(r) {
			break
		}
		ids = append(ids, uptr.BytePtrToString(id))
	}
	return ids, nil
}

// IsNull returns true or false as to whether or not Type has been initialized.
func (t Type) IsNull() bool {
	return t == Null
}
