package source

// #cgo windows CFLAGS: -I ../libobs/include -I libobs/include -I libobs
// #include "obs.h"
import "C"

import (
	"math"
	"syscall"
	"unsafe"

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
// Hotkeys are not supported.
func New(id, name string, settings data.Type) (Type, error) {
	r, err := dll.OBSuintptr("obs_source_create", uptr.FromString(id), uptr.FromString(name), uintptr(settings), uintptr(0))
	return Type(r), err
}

// Output wraps obs_source_t *obs_get_output_source(uint32_t channel).
func Output(channel uint32) (Type, error) {
	r, err := dll.OBSuintptr("obs_get_output_source", uintptr(channel))
	return Type(r), err
}

// Properties wraps obs_properties_t *obs_get_source_properties(const char *id).
func Properties(id string) (properties.Type, error) {
	r, err := dll.OBSuintptr("obs_get_source_properties", uptr.FromString(id))
	return properties.Type(r), err
}

// Sources wraps void obs_enum_sources(bool (*enum_proc)(void*, obs_source_t*), void *param).
// as well as typedef bool (*obs_enum_audio_device_cb)(void *data, const char *name, const char *id).
func Sources() (types []Type, err error) {
	return types, dll.OBS("obs_enum_sources", syscall.NewCallback(
		func(void, source uintptr) int {
			types = append(types, Type(source))
			return 1
		},
	))
}

// Configurable wraps bool obs_source_configurable(const obs_source_t *source).
func (t Type) Configurable() (bool, error) {
	return dll.OBSbool("obs_source_configurable", uintptr(t))
}

// DisplayName wraps const char *obs_source_get_display_name(const char *id).
func (t Type) DisplayName() (string, error) {
	id, err := t.ID()
	if err != nil {
		return "", err
	}
	return dll.OBSstring("obs_source_get_display_name", uptr.FromString(id))
}

// OutputVideo wraps void obs_source_output_video(obs_source_t *source, const struct obs_source_frame *frame).
func (t Type) OutputVideo(v *frame.Video) error {
	return dll.OBS("obs_source_output_video", uintptr(t), uintptr(unsafe.Pointer(v)))
}

// Height wraps uint32_t obs_source_get_height(obs_source_t *source).
func (t Type) Height() (uint32, error) {
	return dll.OBSuint32("obs_source_get_height", uintptr(t))
}

// Hidden wraps bool obs_source_is_hidden(obs_source_t *source).
func (t Type) Hidden() (bool, error) {
	return dll.OBSbool("obs_source_is_hidden", uintptr(t))
}

// ID wraps const char *obs_source_get_id(const obs_source_t *source).
func (t Type) ID() (string, error) {
	return dll.OBSstring("obs_source_get_display_name", uintptr(t))
}

// IsNull returns true or false as to whether or not Type has been initialized.
func (t Type) IsNull() bool {
	return t == Null
}

// Muted wraps bool obs_source_muted(const obs_source_t *source).
func (t Type) Muted() (bool, error) {
	return dll.OBSbool("obs_source_muted", uintptr(t))
}

// Name wraps const char *obs_source_get_name(const obs_source_t *source).
func (t Type) Name() (string, error) {
	return dll.OBSstring("obs_source_get_name", uintptr(t))
}

// OutputAudio wraps void obs_source_output_audio(obs_source_t *source, const struct obs_source_audio *audio).
func (t Type) OutputAudio(a *frame.Audio) error {
	return dll.OBS("obs_source_output_audio", uintptr(t), uintptr(unsafe.Pointer(a)))
}

// ProcHandler wraps proc_handler_t *obs_source_get_proc_handler(const obs_source_t *source).
func (t Type) ProcHandler() (proc.Handler, error) {
	r, err := dll.OBSuintptr("obs_source_get_proc_handler", uintptr(t))
	return proc.Handler(r), err
}

// Properties wraps obs_properties_t *obs_source_properties(const obs_source_t *source).
func (t Type) Properties() (properties.Type, error) {
	r, err := dll.OBSuintptr("obs_source_properties", uintptr(t))
	return properties.Type(r), err
}

// Release wraps void obs_source_release(obs_source_t *source).
func (t Type) Release() error {
	return dll.OBS("obs_source_release")
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
	return dll.OBS("obs_source_set_hidden", uintptr(t), uptr.FromBool(b))
}

// SetMuted wraps void obs_source_set_muted(obs_source_t *source, bool muted).
func (t Type) SetMuted(b bool) error {
	return dll.OBS("obs_source_set_muted", uintptr(t), uptr.FromBool(b))
}

// SetName wraps void obs_source_set_name(obs_source_t *source, const char *name).
func (t Type) SetName(n string) error {
	return dll.OBS("obs_source_set_name", uintptr(t), uptr.FromString(n))
}

// SetVolume wraps void obs_source_set_volume(obs_source_t *source, float volume).
func (t Type) SetVolume(db float32) error {
	return dll.OBS("obs_source_set_volume", uintptr(t), uptr.FromFloat(db))
}

// Settings wraps obs_data_t *obs_source_get_settings(const obs_source_t *source).
func (t Type) Settings() (data.Type, error) {
	r, err := dll.OBSuintptr("obs_source_get_settings", uintptr(t))
	return data.Type(r), err
}

// Scene wraps bool obs_source_is_scene(const obs_source_t *source).
func (t Type) Scene() (bool, error) {
	return dll.OBSbool("obs_source_is_scene", uintptr(t))
}

// Type wraps enum obs_source_type obs_source_get_type(const obs_source_t *source).
func (t Type) Type() (InfoType, error) {
	r, err := dll.OBSuintptr("obs_source_get_type", uintptr(t))
	return InfoType(r), err
}

// Update wraps void obs_source_update(obs_source_t *source, obs_data_t *settings).
func (t Type) Update(settings data.Type) error {
	return dll.OBS("obs_source_update", uintptr(t), uintptr(settings))
}

// VideoRender wraps void obs_source_video_render(obs_source_t *source).
func (t Type) VideoRender() error {
	return dll.OBS("obs_source_video_render", uintptr(t))
}

// Volume wraps float obs_source_get_volume(const obs_source_t *source).
func (t Type) Volume() (float32, error) {
	return dll.OBSfloat32("obs_source_get_volume", uintptr(t))
}

// Width wraps uint32_t obs_source_get_width(obs_source_t *source).
func (t Type) Width() (uint32, error) {
	return dll.OBSuint32("obs_source_get_width", uintptr(t))
}

// Types wraps bool obs_enum_source_types(size_t idx, const char **id).
func Types() (ids []string, err error) {
	for idx := uintptr(0); idx < 1024; idx++ {
		id := uptr.NewBytePtr(4096)

		ok, err := dll.OBSbool("obs_enum_source_types", idx, uintptr(unsafe.Pointer(&id)))
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
