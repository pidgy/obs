package source

/*
#cgo windows CFLAGS: -I ../libobs/include -I libobs/include -I libobs

#include "info.h"
*/
import "C"

import (
	"syscall"
	"unsafe"

	"github.com/pkg/errors"

	"github.com/pidgy/obs/data"
	"github.com/pidgy/obs/dll"
	"github.com/pidgy/obs/frame"
	"github.com/pidgy/obs/graphics"
)

// info_get_name implements const char *(*obs_source_info.info_get_name)(void *type_data).
//
//export info_get_name
func info_get_name(type_data *C.void) *C.cchar_t {
	return C.CString("UniteHUD Capture Source")
}

// info_create wraps void *(*obs_source_info.create)(obs_data_t *settings, obs_source_t *source).
//
//export info_create
func info_create(settings *C.obs_data_t, source *C.obs_source_t) *C.void {
	b := [0]byte{}
	return (*C.void)(&b)
}

// info_destroy wraps void (*obs_source_info.destroy)(void *data).
//
//export info_destroy
func info_destroy(data *C.void) {}

// info_video_render wraps void (*obs_source_info.video_render)(void *data, gs_effect_t *effect).
//
//export info_video_render
func info_video_render(data *C.void, effect *C.gs_effect_t) {}

// info_video_tick wraps void (*obs_source_info.video_tick)(void *data, float seconds).
//
//export info_video_tick
func info_video_tick(data *C.void, seconds C.float) {
	println("tick", int(seconds))
}

// info_get_width wraps uint32_t (*obs_source_info.get_width)(void *data).
//
//export info_get_width
func info_get_width(data *C.void) C.uint32_t {
	return 1920
}

// info_get_height wraps uint32_t (*obs_source_info.get_height)(void *data).
//
//export info_get_height
func info_get_height(data *C.void) C.uint32_t {
	return 1080
}

type (
	// InfoType wraps obs_source_info... clearly.
	Info C.struct_obs_source_info

	// InfoType wraps enum obs_source_type obs_source_info.type.
	InfoType int

	// InfoOutput wraps uint32_t obs_source_info.output_flags.
	InfoOutput int
)

const (
	InfoTypeInput           InfoType = C.OBS_SOURCE_TYPE_INPUT
	InfoTypeFilter          InfoType = C.OBS_SOURCE_TYPE_FILTER
	InfoTypeInputTransition InfoType = C.OBS_SOURCE_TYPE_TRANSITION

	InfoOutputVideo          InfoOutput = C.OBS_SOURCE_VIDEO
	InfoOutputAudio          InfoOutput = C.OBS_SOURCE_AUDIO
	InfoOutputAsync          InfoOutput = C.OBS_SOURCE_ASYNC
	InfoOutputAsyncVideo     InfoOutput = C.OBS_SOURCE_ASYNC_VIDEO
	InfoOutputCustomDraw     InfoOutput = C.OBS_SOURCE_CUSTOM_DRAW
	InfoOutputDoNotDuplicate InfoOutput = C.OBS_SOURCE_DO_NOT_DUPLICATE
)

// NewInfo returns a wrapped obs_source_info.
func NewInfo(id string, t InfoType, o InfoOutput) *Info {
	/*
		struct obs_source_info my_source {
			.id           = "my_source",
			.type         = OBS_SOURCE_TYPE_INPUT,
			.output_flags = OBS_SOURCE_VIDEO,
			.get_name     = my_source_name,
			.create       = my_source_create,
			.destroy      = my_source_destroy,
			.update       = my_source_update,
			.video_render = my_source_render,
			.get_width    = my_source_width,
			.get_height   = my_source_height
		};
	*/
	i := &Info{
		id:           C.CString(id),
		_type:        uint32(t),
		output_flags: C.uint(o),
		get_name:     (*[0]byte)(unsafe.Pointer(C.info_get_name)),
		create:       (*[0]byte)(C.info_create),
		destroy:      (*[0]byte)(C.info_destroy),
		video_render: (*[0]byte)(C.info_video_render),
		video_tick:   (*[0]byte)(C.info_video_tick),
		get_width:    (*[0]byte)(C.info_get_width),
		get_height:   (*[0]byte)(C.info_get_height),
	}

	return i
}

// Create wraps void *(*obs_source_info.create)(obs_data_t *settings, obs_source_t *source).
func (i *Info) Create(d data.Type, t Type) (*Info, error) {
	err := C.invoke_info_create(
		C.closure_info_create(i.create),
		(*C.obs_data_t)(unsafe.Pointer(d)),
		(*C.obs_source_t)(unsafe.Pointer(t)),
	)
	if err == nil {
		return i, errors.Wrap(errors.Errorf("failed to create source"), "invoke_info_create")
	}

	return i, nil
}

// Destroy wraps void (*obs_source_info.destroy)(void *data).
func (i *Info) Destroy() {
	b := [0]byte{}

	C.invoke_info_destroy(
		C.closure_info_destroy(
			i.destroy,
		),
		unsafe.Pointer(&b),
	)
}

// Free releases any memory allocated by Info.
func (i *Info) Free() {
	C.free(unsafe.Pointer(i.id))
}

// Height wraps uint32_t (*obs_source_info.get_height)(void *data).
func (i *Info) Height() uint32 {
	b := [0]byte{}

	return uint32(C.invoke_info_get_height(
		C.closure_info_get_height(
			i.get_height,
		),
		unsafe.Pointer(&b),
	))
}

// ID returns the unexported ID associated with a source.
func (i *Info) ID() string {
	return C.GoString(i.id)
}

// Name wraps const char *(*obs_source_info.info_get_name)(void *type_data).
func (i *Info) Name() string {
	b := [0]byte{}

	return C.GoString(
		C.invoke_info_get_name(
			C.closure_info_get_name(
				i.get_name,
			),
			unsafe.Pointer(&b),
		),
	)
}

// OutputVideo wraps void obs_source_output_video(obs_source_t *source, const struct obs_source_frame *frame).
func (i *Info) OutputVideo(v *frame.Video) error {
	_, _, err := dll.OBS.NewProc("obs_source_output_video").Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(v)),
	)
	if err != syscall.Errno(0) {
		return errors.Wrap(err, "obs_source_output_video")
	}

	return nil
}

// Register wraps void obs_register_source(struct obs_source_info *info).
func (i *Info) Register() error {
	_, _, err := dll.OBS.NewProc("obs_register_source_s").Call(
		uintptr(unsafe.Pointer(i)),
		uintptr((C.size_t)(unsafe.Sizeof(Info{}))),
	)
	if err != syscall.Errno(0) {
		return errors.Wrap(err, "obs_register_source_s")
	}

	return nil
}

// VideoRender wraps void (*obs_source_info.video_render)(void *data, gs_effect_t *effect).
func (i *Info) VideoRender(fx graphics.Effect) {
	b := [0]byte{}

	C.invoke_info_video_render(
		C.closure_info_video_render(
			i.video_render,
		),
		unsafe.Pointer(&b),
		(*C.gs_effect_t)(unsafe.Pointer(fx)),
	)
}

// VideoTick wraps void (*obs_source_info.video_tick)(void *data, float seconds).
func (i *Info) VideoTick(seconds float32) {
	b := [0]byte{}

	C.invoke_info_video_tick(
		C.closure_info_video_tick(
			i.video_tick,
		),
		unsafe.Pointer(&b),
		(C.float)(seconds),
	)
}

// Width wraps uint32_t (*obs_source_info.get_width)(void *data).
func (i *Info) Width() uint32 {
	b := [0]byte{}

	return uint32(C.invoke_info_get_width(
		C.closure_info_get_width(
			i.get_width,
		),
		unsafe.Pointer(&b),
	))
}
