package source

import (
	"math"
	"syscall"
	"unsafe"

	"github.com/pkg/errors"

	"github.com/pidgy/obs/data"
	"github.com/pidgy/obs/dll"
	"github.com/pidgy/obs/frame"
	"github.com/pidgy/obs/uptr"
)

type (
	// Type wraps obs_source_t.
	Type uintptr
)

const (
	// Null represents a Null obs_source_t.
	Null = Type(0)
)

var (
	// MinVolumeDb is the minimum decible value for a source's volume.
	MinVolumeDb = float32(math.Inf(-1))
	// MaxVolumeDb is the maximum decible value for a source's volume.
	MaxVolumeDb = float32(0)
)

// Create wraps obs_source_t *obs_source_create(const char *id, const char *name,obs_data_t *settings,obs_data_t *hotkey_data).
func Create(id, name string, settings, hotkeys data.Type) (Type, error) {
	r, _, err := dll.OBS.NewProc("obs_source_create").Call(
		uptr.FromString(id),
		uptr.FromString(name),
		uintptr(settings),
		uintptr(hotkeys),
	)
	if err != syscall.Errno(0) {
		return Null, errors.Wrap(err, "obs_source_create")
	}

	return Type(r), nil
}

// Enum wraps void obs_enum_sources(bool (*enum_proc)(void*, obs_source_t*), void *param).
func Enum() ([]Type, error) {
	var t []Type

	// typedef bool (*obs_enum_audio_device_cb)(void *data, const char *name, const char *id).
	callback := syscall.NewCallback(
		func(void, source uintptr) int {
			t = append(t, Type(source))
			return 1
		},
	)
	_, _, err := dll.OBS.NewProc("obs_enum_sources").Call(
		callback,
	)
	if err != syscall.Errno(0) {
		return nil, errors.Wrap(err, "obs_enum_sources")
	}

	return t, nil
}

// EnumTypes wraps bool obs_enum_source_types(size_t idx, const char **id).
func EnumTypes() (ids []string, err error) {
	for idx := uintptr(0); idx < 1024; idx++ {
		id := uptr.NewBytePtr(4096)

		r, _, err := dll.OBS.NewProc("obs_enum_source_types").Call(
			idx,
			uintptr(unsafe.Pointer(&id)),
		)
		if err != syscall.Errno(0) {
			return nil, errors.Wrap(err, "obs_enum_source_types")
		}

		if !uptr.Bool(r) {
			break
		}
		ids = append(ids, uptr.BytePtrToString(id))
	}
	return ids, nil
}

// ID wraps const char *obs_source_get_id(const obs_source_t *source).
func (t Type) ID() (string, error) {
	r, _, err := dll.OBS.NewProc("obs_source_get_id").Call(
		uintptr(unsafe.Pointer(t)),
	)
	if err != syscall.Errno(0) {
		return "", errors.Wrap(err, "obs_source_get_id")
	}

	return uptr.String(r), nil
}

// IsNull returns true or false as to whether or not Type has been initialized.
func (t Type) IsNull() bool {
	return t == Null
}

// Muted wraps bool obs_source_muted(const obs_source_t *source).
func (t Type) Muted() (bool, error) {
	r, _, err := dll.OBS.NewProc("obs_source_muted").Call(
		uintptr(unsafe.Pointer(t)),
	)
	if err != syscall.Errno(0) {
		return false, errors.Wrap(err, "obs_source_muted")
	}

	return uptr.Bool(r), nil
}

// Name wraps const char *obs_source_get_name(const obs_source_t *source).
func (t Type) Name() (string, error) {
	r, _, err := dll.OBS.NewProc("obs_source_get_name").Call(
		uintptr(unsafe.Pointer(t)),
	)
	if err != syscall.Errno(0) {
		return "", errors.Wrap(err, "obs_source_get_name")
	}

	return uptr.String(r), nil
}

// Output wraps obs_source_t *obs_get_output_source(uint32_t channel).
func Output(channel uint32) (Type, error) {
	r, _, err := dll.OBS.NewProc("obs_get_output_source").Call(
		uintptr(channel),
	)
	if err != syscall.Errno(0) {
		return Null, errors.Wrap(err, "obs_get_output_source")
	}

	return Type(r), nil
}

// OutputAudio wraps void obs_source_output_audio(obs_source_t *source, const struct obs_source_audio *audio).
func (t Type) OutputAudio(a *frame.Audio) error {
	_, _, err := dll.OBS.NewProc("obs_source_output_audio").Call(
		uintptr(unsafe.Pointer(t)),
		uintptr(unsafe.Pointer(a)),
	)
	if err != syscall.Errno(0) {
		return errors.Wrap(err, "obs_source_output_audio")
	}

	return nil
}

// OutputVideo wraps void obs_source_output_video(obs_source_t *source, const struct obs_source_frame *frame).
func (t Type) OutputVideo(v *frame.Video) error {
	_, _, err := dll.OBS.NewProc("obs_source_output_video").Call(
		uintptr(unsafe.Pointer(t)),
		uintptr(unsafe.Pointer(v)),
	)
	if err != syscall.Errno(0) {
		return errors.Wrap(err, "obs_source_output_video")
	}

	return nil
}

// Release wraps void obs_source_release(obs_source_t *source).
func (t Type) Release() error {
	_, _, err := dll.OBS.NewProc("obs_source_release").Call(
		uintptr(unsafe.Pointer(t)),
	)
	if err != syscall.Errno(0) {
		return errors.Wrap(err, "obs_source_release")
	}
	return nil
}

// SetMuted wraps void obs_source_set_muted(obs_source_t *source, bool muted).
func (t Type) SetMuted(m bool) error {
	_, _, err := dll.OBS.NewProc("obs_source_set_muted").Call(
		uintptr(unsafe.Pointer(t)),
		uptr.FromBool(m),
	)
	if err != syscall.Errno(0) {
		return errors.Wrap(err, "obs_source_set_muted")
	}

	return nil
}

// SetName wraps void obs_source_set_name(obs_source_t *source, const char *name).
func (t Type) SetName(n string) error {
	_, _, err := dll.OBS.NewProc("obs_source_set_name").Call(
		uintptr(unsafe.Pointer(t)),
		uptr.FromString(n),
	)
	if err != syscall.Errno(0) {
		return errors.Wrap(err, "obs_source_set_name")
	}

	return nil
}

// SetVolume wraps void obs_source_set_volume(obs_source_t *source, float volume).
func (t Type) SetVolume(db float32) error {
	_, _, err := dll.OBS.NewProc("obs_source_set_volume").Call(
		uintptr(unsafe.Pointer(t)),
		uptr.FromFloat(db),
	)
	if err != syscall.Errno(0) {
		return errors.Wrap(err, "obs_source_set_volume")
	}

	return nil
}

// Scene wraps bool obs_source_is_scene(const obs_source_t *source).
func (t Type) Scene() (bool, error) {
	r, _, err := dll.OBS.NewProc("obs_source_is_scene").Call(
		uintptr(unsafe.Pointer(t)),
	)
	if err != syscall.Errno(0) {
		return false, errors.Wrap(err, "obs_source_is_scene")
	}

	return uptr.Bool(r), nil
}

// Volume wraps float obs_source_get_volume(const obs_source_t *source).
func (t Type) Volume() (float32, error) {
	_, r, err := dll.OBS.NewProc("obs_source_get_volume").Call(
		uintptr(unsafe.Pointer(t)),
	)
	if err != syscall.Errno(0) {
		return 0, errors.Wrap(err, "obs_source_get_volume")
	}

	return uptr.Float(r), nil
}
