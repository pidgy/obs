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

	"github.com/pidgy/obs/dll"
	"github.com/pidgy/obs/frame"
)

type (
	// Info wraps obs_source_info... clearly.
	Info C.struct_obs_source_info

	// InfoType wraps enum obs_source_type obs_source_info.type.
	InfoType int

	// InfoOutputFlag wraps uint32_t obs_source_info.output_flags.
	InfoOutputFlag int

	InfoOptions struct {
		ID string

		Type  InfoType
		Flags InfoOutputFlag
	}
)

const (
	InfoTypeInput           InfoType = C.OBS_SOURCE_TYPE_INPUT
	InfoTypeFilter          InfoType = C.OBS_SOURCE_TYPE_FILTER
	InfoTypeInputTransition InfoType = C.OBS_SOURCE_TYPE_TRANSITION

	InfoOutputFlagVideo          InfoOutputFlag = C.OBS_SOURCE_VIDEO
	InfoOutputFlagAudio          InfoOutputFlag = C.OBS_SOURCE_AUDIO
	InfoOutputFlagAsync          InfoOutputFlag = C.OBS_SOURCE_ASYNC
	InfoOutputFlagAsyncVideo     InfoOutputFlag = C.OBS_SOURCE_ASYNC_VIDEO
	InfoOutputFlagCustomDraw     InfoOutputFlag = C.OBS_SOURCE_CUSTOM_DRAW
	InfoOutputFlagDoNotDuplicate InfoOutputFlag = C.OBS_SOURCE_DO_NOT_DUPLICATE
)

// Register wraps void obs_register_source(struct obs_source_info *info).
func Register(o *InfoOptions) error {
	i := &Info{
		id:           C.CString(o.ID),
		_type:        uint32(o.Type),
		output_flags: C.uint(o.Flags),

		get_name:     (*[0]byte)(unsafe.Pointer(C.info_get_name)),
		create:       (*[0]byte)(C.info_create),
		destroy:      (*[0]byte)(C.info_destroy),
		video_render: (*[0]byte)(C.info_video_render),
		video_tick:   (*[0]byte)(C.info_video_tick),
		get_width:    (*[0]byte)(C.info_get_width),
		get_height:   (*[0]byte)(C.info_get_height),
		update:       (*[0]byte)(C.info_update),
	}

	_, _, err := dll.OBS.NewProc("obs_register_source_s").Call(
		uintptr(unsafe.Pointer(i)),
		uintptr((C.size_t)(unsafe.Sizeof(C.struct_obs_source_info{}))),
	)
	if err != syscall.Errno(0) {
		return errors.Wrap(err, "obs_register_source_s")
	}

	return nil
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

// update implements void (*obs_source_info.update)(void *data, obs_data_t *settings).
//
//export info_update
func info_update(data *C.void, settings *C.obs_data_t) {
	println("-> info_update")
}

// info_get_name implements const char *(*obs_source_info.info_get_name)(void *type_data).
//
//export info_get_name
func info_get_name(type_data *C.void) *C.const_char_t {
	i := (*Info)(unsafe.Pointer(type_data))

	println("->", C.GoString(i.id), "info_get_name")

	return i.id
}

// info_create wraps void *(*obs_source_info.create)(obs_data_t *settings, obs_source_t *source).
//
//export info_create
func info_create(settings *C.obs_data_t, source *C.obs_source_t) *C.void {
	println("-> info_create")

	b := [0]byte{}
	return (*C.void)(&b)
}

// info_destroy wraps void (*obs_source_info.destroy)(void *data).
//
//export info_destroy
func info_destroy(data *C.void) {
	println("-> info_destroy")
}

// info_video_render wraps void (*obs_source_info.video_render)(void *data, gs_effect_t *effect).
//
//export info_video_render
func info_video_render(data *C.void, effect *C.gs_effect_t) {
	println("-> info_video_render")
}

// info_video_tick wraps void (*obs_source_info.video_tick)(void *data, float seconds).
//
//export info_video_tick
func info_video_tick(data *C.void, seconds C.float) {
	println("-> info_video_tick", int(seconds))
}

// info_get_width wraps uint32_t (*obs_source_info.get_width)(void *data).
//
//export info_get_width
func info_get_width(data *C.void) C.uint32_t {
	println("-> info_get_width")
	return 0

	// i := (*Info)(unsafe.Pointer(data))
	// return C.uint32_t(i.opts.Width)
}

// info_get_height wraps uint32_t (*obs_source_info.get_height)(void *data).
//
//export info_get_height
func info_get_height(data *C.void) C.uint32_t {
	i := (*Info)(unsafe.Pointer(data))
	println("->", C.GoString(i.id), "info_get_height")
	return C.uint32_t(0)
}
