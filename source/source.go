package source

import (
	"math"
	"syscall"
	"unsafe"

	"github.com/pidgy/obs/data"
	"github.com/pidgy/obs/frame"
	"github.com/pidgy/obs/lib"
	"github.com/pidgy/obs/uptr"
)

type (
	// Type wraps obs_source_t.
	Type uintptr
)

const (
	// NULL represents a NULL obs_source_t.
	NULL = Type(0)
)

var (
	// MinVolumeDb is the minimum decible value for OBS an source's volume.
	MinVolumeDb = float32(math.Inf(-1))
	// MaxVolumeDb is the maximum decible value for OBS an source's volume.
	MaxVolumeDb = float32(0)
)

// New wraps obs_source_t *obs_source_create(const char *id, const char *name,obs_data_t *settings,obs_data_t *hotkey_data).
func New(id, name string, settings, hotkeys data.Type) (Type, error) {
	r, _, err := lib.OBS.NewProc("obs_source_create").Call(
		uptr.FromString(id),
		uptr.FromString(name),
		uintptr(settings),
		uintptr(hotkeys),
	)
	if err != syscall.Errno(0) {
		return 0, err
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
	_, _, err := lib.OBS.NewProc("obs_enum_sources").Call(
		callback,
	)
	if err != syscall.Errno(0) {
		return nil, err
	}

	return t, nil
}

// ID wraps const char *obs_source_get_id(const obs_source_t *source).
func (t Type) ID() (string, error) {
	r, _, err := lib.OBS.NewProc("obs_source_get_id").Call(
		uintptr(unsafe.Pointer(t)),
	)
	if err != syscall.Errno(0) {
		return "", err
	}

	return uptr.String(r), nil
}

// Muted wraps bool obs_source_muted(const obs_source_t *source).
func (t Type) Muted() (bool, error) {
	r, _, err := lib.OBS.NewProc("obs_source_muted").Call(
		uintptr(unsafe.Pointer(t)),
	)
	if err != syscall.Errno(0) {
		return false, err
	}

	return uptr.Bool(r), nil
}

// Name wraps const char *obs_source_get_name(const obs_source_t *source).
func (t Type) Name() (string, error) {
	r, _, err := lib.OBS.NewProc("obs_source_get_name").Call(
		uintptr(unsafe.Pointer(t)),
	)
	if err != syscall.Errno(0) {
		return "", err
	}

	return uptr.String(r), nil
}

// OutputAudio wraps void obs_source_output_audio(obs_source_t *source, const struct obs_source_audio *audio).
func (t Type) OutputAudio(a *frame.Audio) error {
	_, _, err := lib.OBS.NewProc("obs_source_output_audio").Call(
		uintptr(unsafe.Pointer(t)),
		uintptr(unsafe.Pointer(a)),
	)
	if err != syscall.Errno(0) {
		return err
	}

	return nil
}

// OutputVideo wraps void obs_source_output_video(obs_source_t *source, const struct obs_source_frame *frame).
func (t Type) OutputVideo(v *frame.Video) error {
	_, _, err := lib.OBS.NewProc("obs_source_output_video").Call(
		uintptr(unsafe.Pointer(t)),
		uintptr(unsafe.Pointer(v)),
	)
	if err != syscall.Errno(0) {
		return err
	}

	return nil
}

// Release wraps void obs_source_release(obs_source_t *source).
func (t Type) Release() error {
	_, _, err := lib.OBS.NewProc("obs_source_release").Call(
		uintptr(unsafe.Pointer(t)),
	)
	if err != syscall.Errno(0) {
		return err
	}
	return nil
}

// SetMuted wraps void obs_source_set_muted(obs_source_t *source, bool muted).
func (t Type) SetMuted(m bool) error {
	_, _, err := lib.OBS.NewProc("obs_source_set_muted").Call(
		uintptr(unsafe.Pointer(t)),
		uptr.FromBool(m),
	)
	if err != syscall.Errno(0) {
		return err
	}

	return nil
}

// SetName wraps void obs_source_set_name(obs_source_t *source, const char *name).
func (t Type) SetName(n string) error {
	_, _, err := lib.OBS.NewProc("obs_source_set_name").Call(
		uintptr(unsafe.Pointer(t)),
		uptr.FromString(n),
	)
	if err != syscall.Errno(0) {
		return err
	}

	return nil
}

// SetVolume wraps void obs_source_set_volume(obs_source_t *source, float volume).
func (t Type) SetVolume(db float32) error {
	_, _, err := lib.OBS.NewProc("obs_source_set_volume").Call(
		uintptr(unsafe.Pointer(t)),
		uptr.FromFloat(db),
	)
	if err != syscall.Errno(0) {
		return err
	}

	return nil
}

// Scene wraps bool obs_source_is_scene(const obs_source_t *source).
func (t Type) Scene() (bool, error) {
	r, _, err := lib.OBS.NewProc("obs_source_is_scene").Call(
		uintptr(unsafe.Pointer(t)),
	)
	if err != syscall.Errno(0) {
		return false, err
	}

	return uptr.Bool(r), nil
}

// Volume wraps float obs_source_get_volume(const obs_source_t *source).
func (t Type) Volume() (float32, error) {
	_, r, err := lib.OBS.NewProc("obs_source_get_volume").Call(
		uintptr(unsafe.Pointer(t)),
	)
	if err != syscall.Errno(0) {
		return 0, err
	}

	return uptr.Float(r), nil
}
