package filter

import (
	"unsafe"

	"github.com/pidgy/obs/dll"
	"github.com/pidgy/obs/uptr"
)

// EnumTypes wraps bool obs_enum_filter_types(size_t idx, const char **id).
func EnumTypes() (ids []string, err error) {
	for idx := uintptr(0); idx < 1024; idx++ {
		id := uptr.NewBytePtr(4096)

		ok, err := dll.OBSbool("obs_enum_filter_types", idx, uintptr(unsafe.Pointer(&id)))
		if err != nil {
			return nil, err
		}
		if !ok {
			break
		}

		ids = append(ids, uptr.BytePtrToString(id))
	}
	return ids, nil
}
