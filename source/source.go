package source

// #cgo windows CFLAGS: -I ../libobs/include -I libobs/include -I libobs
// #include "obs.h"
import "C"

import (
	"math"
	"syscall"
	"unsafe"

	"github.com/pkg/errors"

	"github.com/pidgy/obs/data"
	"github.com/pidgy/obs/dll"
	"github.com/pidgy/obs/frame"
	"github.com/pidgy/obs/proc"
	"github.com/pidgy/obs/properties"
	"github.com/pidgy/obs/uptr"
)

type (
	// Type wraps obs_source_t.
	Type uintptr
)

const (
	// Null represents a nil obs_source_t.
	Null = Type(0)
)

var (
	// MinVolumeDb is the minimum decible value for a source's volume.
	MinVolumeDb = float32(math.Inf(-1))
	// MaxVolumeDb is the maximum decible value for a source's volume.
	MaxVolumeDb = float32(0)
)

// New wraps obs_source_t *obs_source_create(const char *id, const char *name, obs_data_t *settings, obs_data_t *hotkey_data).
func New(id, name string, settings data.Type) (Type, error) {
	r, _, err := dll.OBS.NewProc("obs_source_create").Call(
		uptr.FromString(id),
		uptr.FromString(name),
		uintptr(settings),
		uintptr(0), // Hotkeys are not supported.
	)
	if err != syscall.Errno(0) {
		return Null, errors.Wrap(err, "obs_source_create")
	}
	return Type(r), nil
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

// Sources wraps void obs_enum_sources(bool (*enum_proc)(void*, obs_source_t*), void *param).
func Sources() ([]Type, error) {
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

// Configurable wraps bool obs_source_configurable(const obs_source_t *source).
func (t Type) Configurable() (bool, error) {
	r, _, err := dll.OBS.NewProc("obs_source_configurable").Call(
		uintptr(t),
	)
	if err != syscall.Errno(0) {
		return false, errors.Wrap(err, "obs_source_configurable")
	}
	return uptr.Bool(r), nil
}

// DisplayName wraps const char *obs_source_get_display_name(const char *id).
func (t Type) DisplayName() (string, error) {
	id, err := t.ID()
	if err != nil {
		return "", err
	}
	r, _, err := dll.OBS.NewProc("obs_source_get_display_name").Call(
		uptr.FromString(id),
	)
	if err != syscall.Errno(0) {
		return "", errors.Wrap(err, "obs_source_get_display_name")
	}
	return uptr.String(r), nil
}

// OutputVideo wraps void obs_source_output_video(obs_source_t *source, const struct obs_source_frame *frame).
func (t Type) OutputVideo(v *frame.Video) error {
	_, _, err := dll.OBS.NewProc("obs_source_output_video").Call(
		uintptr(t),
		uintptr(unsafe.Pointer(v)),
	)
	if err != syscall.Errno(0) {
		return errors.Wrap(err, "obs_source_output_video")
	}
	return nil
}

// Height wraps uint32_t obs_source_get_height(obs_source_t *source).
func (t Type) Height() (int32, error) {
	r, _, err := dll.OBS.NewProc("obs_source_get_height").Call(
		uintptr(t),
	)
	if err != syscall.Errno(0) {
		return 0, errors.Wrap(err, "obs_source_get_height")
	}
	return int32(r), nil
}

// Hidden wraps bool obs_source_is_hidden(obs_source_t *source).
func (t Type) Hidden() (bool, error) {
	r, _, err := dll.OBS.NewProc("obs_source_is_hidden").Call(
		uintptr(t),
	)
	if err != syscall.Errno(0) {
		return false, errors.Wrap(err, "obs_source_is_hidden")
	}
	return uptr.Bool(r), nil
}

// ID wraps const char *obs_source_get_id(const obs_source_t *source).
func (t Type) ID() (string, error) {
	r, _, err := dll.OBS.NewProc("obs_source_get_id").Call(
		uintptr(t),
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
		uintptr(t),
	)
	if err != syscall.Errno(0) {
		return false, errors.Wrap(err, "obs_source_muted")
	}

	return uptr.Bool(r), nil
}

// Name wraps const char *obs_source_get_name(const obs_source_t *source).
func (t Type) Name() (string, error) {
	r, _, err := dll.OBS.NewProc("obs_source_get_name").Call(
		uintptr(t),
	)
	if err != syscall.Errno(0) {
		return "", errors.Wrap(err, "obs_source_get_name")
	}

	return uptr.String(r), nil
}

// OutputAudio wraps void obs_source_output_audio(obs_source_t *source, const struct obs_source_audio *audio).
func (t Type) OutputAudio(a *frame.Audio) error {
	_, _, err := dll.OBS.NewProc("obs_source_output_audio").Call(
		uintptr(t),
		uintptr(unsafe.Pointer(a)),
	)
	if err != syscall.Errno(0) {
		return errors.Wrap(err, "obs_source_output_audio")
	}

	return nil
}

// ProcHandler wraps proc_handler_t *obs_source_get_proc_handler(const obs_source_t *source).
func (t Type) ProcHandler() (proc.Handler, error) {
	r, _, err := dll.OBS.NewProc("obs_source_get_proc_handler").Call(
		uintptr(t),
	)
	if err != syscall.Errno(0) {
		return proc.Null, errors.Wrap(err, "obs_source_get_proc_handler")
	}
	return proc.Handler(r), nil
}

// Properties wraps obs_properties_t *obs_source_properties(const obs_source_t *source).
func (t Type) Properties() (properties.Type, error) {
	r, _, err := dll.OBS.NewProc("obs_source_properties").Call(
		uintptr(t),
	)
	if err != syscall.Errno(0) {
		return properties.Null, errors.Wrap(err, "obs_source_properties")
	}
	return properties.Type(r), nil
}

// PropertiesByID wraps obs_properties_t *obs_get_source_properties(const char *id).
func (t Type) PropertiesByID(id string) (properties.Type, error) {
	r, _, err := dll.OBS.NewProc("obs_get_source_properties").Call(
		uptr.FromString(id),
	)
	if err != syscall.Errno(0) {
		return properties.Null, errors.Wrap(err, "obs_get_source_properties")
	}
	return properties.Type(r), nil
}

// MustRelease wraps void obs_source_release(obs_source_t *source).
func (t Type) MustRelease() {
	_, _, err := dll.OBS.NewProc("obs_source_release").Call()
	if err != syscall.Errno(0) {
		panic(errors.Wrap(err, "obs_source_release"))
	}
}

// Release wraps void obs_source_release(obs_source_t *source).
func (t Type) Release() error {
	_, _, err := dll.OBS.NewProc("obs_source_release").Call(
		uintptr(t),
	)
	if err != syscall.Errno(0) {
		return errors.Wrap(err, "obs_source_release")
	}
	return nil
}

// Resolution wraps uint32_t obs_source_get_width/height(obs_source_t *source).
func (t Type) Resolution() (w, h uint32, err error) {
	w, err = t.Width()
	if err != nil {
		return 0, 0, err
	}
	h, err = t.Width()
	if err != nil {
		return 0, 0, err
	}
	return
}

// SetHidden wraps void obs_source_set_hidden(obs_source_t *source, bool hidden).
func (t Type) SetHidden(b bool) error {
	_, _, err := dll.OBS.NewProc("obs_source_set_hidden").Call(
		uintptr(t),
		uptr.FromBool(b),
	)
	if err != syscall.Errno(0) {
		return errors.Wrap(err, "obs_source_set_hidden")
	}
	return nil
}

// SetMuted wraps void obs_source_set_muted(obs_source_t *source, bool muted).
func (t Type) SetMuted(m bool) error {
	_, _, err := dll.OBS.NewProc("obs_source_set_muted").Call(
		uintptr(t),
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
		uintptr(t),
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
		uintptr(t),
		uptr.FromFloat(db),
	)
	if err != syscall.Errno(0) {
		return errors.Wrap(err, "obs_source_set_volume")
	}
	return nil
}

// Settings wraps obs_data_t *obs_source_get_settings(const obs_source_t *source).
func (t Type) Settings() (data.Type, error) {
	r, _, err := dll.OBS.NewProc("obs_source_get_settings").Call(
		uintptr(t),
	)
	if err != syscall.Errno(0) {
		return data.Null, errors.Wrap(err, "obs_source_get_settings")
	}
	return data.Type(r), nil
}

// Scene wraps bool obs_source_is_scene(const obs_source_t *source).
func (t Type) Scene() (bool, error) {
	r, _, err := dll.OBS.NewProc("obs_source_is_scene").Call(
		uintptr(t),
	)
	if err != syscall.Errno(0) {
		return false, errors.Wrap(err, "obs_source_is_scene")
	}

	return uptr.Bool(r), nil
}

// Type wraps enum obs_source_type obs_source_get_type(const obs_source_t *source).
func (t Type) Type() (InfoType, error) {
	r, _, err := dll.OBS.NewProc("obs_source_get_type").Call(
		uintptr(t),
	)
	if err != syscall.Errno(0) {
		return -1, errors.Wrap(err, "obs_source_get_type")
	}

	return InfoType(r), nil
}

// Update wraps void obs_source_update(obs_source_t *source, obs_data_t *settings).
func (t Type) Update(settings data.Type) error {
	_, _, err := dll.OBS.NewProc("obs_source_update").Call(
		uintptr(t),
		uintptr(settings),
	)
	if err != syscall.Errno(0) {
		return errors.Wrap(err, "obs_source_update")
	}
	return nil
}

// VideoRender wraps void obs_source_video_render(obs_source_t *source).
func (t Type) VideoRender() error {
	_, _, err := dll.OBS.NewProc("obs_source_video_render").Call(
		uintptr(t),
	)
	if err != syscall.Errno(0) {
		return errors.Wrap(err, "obs_source_video_render")
	}
	return nil
}

// Volume wraps float obs_source_get_volume(const obs_source_t *source).
func (t Type) Volume() (float32, error) {
	_, r, err := dll.OBS.NewProc("obs_source_get_volume").Call(
		uintptr(t),
	)
	if err != syscall.Errno(0) {
		return 0, errors.Wrap(err, "obs_source_get_volume")
	}

	return uptr.Float(r), nil
}

// Width wraps uint32_t obs_source_get_width(obs_source_t *source).
func (t Type) Width() (uint32, error) {
	r, _, err := dll.OBS.NewProc("obs_source_get_width").Call(
		uintptr(t),
	)
	if err != syscall.Errno(0) {
		return 0, errors.Wrap(err, "obs_source_get_width")
	}
	return uint32(r), nil
}

// Types wraps bool obs_enum_source_types(size_t idx, const char **id).
func Types() (ids []string, err error) {
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
