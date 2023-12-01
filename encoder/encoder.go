package encoder

import (
	"syscall"
	"unsafe"

	"github.com/pidgy/obs/dll"
	"github.com/pidgy/obs/uptr"
)

// Type wraps *obs_encoder_t.
type Type uintptr

const (
	Null = Type(0)
)

// Enum wraps void void obs_enum_encoders(bool (*enum_proc)(void*, obs_encoder_t*), void *param).
func Enum() ([]Type, error) {
	var t []Type
	callback := syscall.NewCallback(
		func(void, e uintptr) int {
			t = append(t, Type(e))
			return 1
		},
	)
	return t, dll.OBS("obs_enum_encoders", callback)
}

// EnumTypes wraps bool obs_enum_encoder_types(size_t idx, const char **id).
func EnumTypes() (ids []string, err error) {
	for idx := uintptr(0); idx < 1024; idx++ {
		id := uptr.NewBytePtr(4096)

		ok, err := dll.OBSbool("obs_enum_encoder_types", idx, uintptr(unsafe.Pointer(&id)))
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

// IsNull returns true or false as to whether or not Type has been initialized.
func (t Type) IsNull() bool {
	return t == Null
}

// Ref wraps obs_encoder_t *obs_encoder_get_ref(obs_encoder_t *encoder).
func (t Type) Ref() (Type, error) {
	r, err := dll.OBSuintptr("obs_encoder_get_ref", uintptr(t))
	if err != nil {
		return Null, err
	}
	return Type(r), nil
}

// Release wraps void obs_encoder_release(obs_encoder_t *encoder).
func (t Type) Release() error {
	return dll.OBS("obs_encoder_release", uintptr(t))
}
