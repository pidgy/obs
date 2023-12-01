package input

import (
	"unsafe"

	"github.com/pkg/errors"

	"github.com/pidgy/obs/dll"
	"github.com/pidgy/obs/uptr"
)

// EnumTypes wraps bool obs_enum_input_types(size_t idx, const char **id).
func EnumTypes() (ids []string, err error) {
	for idx := uintptr(0); idx < 1024; idx++ {
		id := uptr.NewBytePtr(4096)

		ok, err := dll.OBSbool("obs_enum_input_types", idx, uintptr(unsafe.Pointer(&id)))
		if err != nil {
			return nil, errors.Wrap(err, "obs_enum_input_types")
		}
		if !ok {
			break
		}

		ids = append(ids, uptr.BytePtrToString(id))
	}
	return ids, nil
}
