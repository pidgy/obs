package video

// #cgo windows CFLAGS: -I ../libobs/include -I libobs/include -I libobs
// #include "obs.h"
import "C"

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"unsafe"

	"github.com/pidgy/obs/core"
	"github.com/pidgy/obs/dll"
	"github.com/pidgy/obs/graphics"
	"github.com/pkg/errors"
)

const Null = Type(0)

type (
	// Info wraps struct obs_video_info.
	Info struct {
		graphics.Module
		FPS
		Base,
		Output Resolution
		Format
		Adapter       uint32
		GPUConversion bool
		ColorSpace
		Range
		graphics.Scale
	}

	infoC C.struct_obs_video_info

	// Data wraps video_data.
	Data uintptr

	// RawCallback wraps the callback parameter for obs_add/remove_raw_video_callback.
	RawCallback struct {
		handle uintptr
	}

	// ScaleInfo wraps struct video_scale_info.
	ScaleInfo struct {
		Format Format
		Width  uint32
		Height uint32

		Range      Range
		Colorspace ColorSpace
	}

	// Type wraps obs_video_info.
	Type uintptr

	// ResetResult wraps the response of obs_reset_video.
	ResetResult int

	FPS struct {
		N, D uint
	}

	Resolution struct {
		W, H uint
	}
)

const (
	ResetResultSuccess ResetResult = -iota
	ResetResultFail
	ResetResultNotSupported
	ResetResultInvalidParam
	ResetResultCurrentlyActive
	ResetResultModuleNotFound
)

var (
	FPS60 = FPS{60, 1}

	Resolution1920x1080 = Resolution{1920, 1080}
)

// NewRawCallback returns a RawCallback used for obs_add/remove_raw_video_callback.
func NewRawCallback(callback func(*Data)) *RawCallback {
	return &RawCallback{
		handle: syscall.NewCallback(
			func(_, frame uintptr) uintptr {
				callback((*Data)(unsafe.Pointer(&frame)))
				return 0
			},
		),
	}
}

// Get wraps obs_get_video.
func Get() (Type, error) {
	r, err := dll.OBSuintptr("obs_get_video")
	return Type(r), err
}

// c returns the cgo converstion of an Info.
func (i Info) c() *infoC {
	return &infoC{
		graphics_module: C.CString(i.Module.String()),
		fps_num:         C.uint(i.FPS.N),
		fps_den:         C.uint(i.FPS.D),
		base_width:      C.uint(i.Base.W),
		base_height:     C.uint(i.Base.H),
		output_width:    C.uint(i.Output.W),
		output_height:   C.uint(i.Output.H),
		output_format:   i.Format.Uint32(),
		gpu_conversion:  C.bool(true),
		colorspace:      i.ColorSpace.Uint32(),
		_range:          i.Range.Uint32(),
		scale_type:      i.Scale.Uint32(),
	}
}

// Add wraps void obs_add_raw_video_callback(
//
//	const struct video_scale_info *conversion,
//	void (*callback)(void *param, struct video_data *frame),
//	void *param
//
// ).
func (r *RawCallback) Add(s *ScaleInfo) error {
	return dll.OBS("obs_add_raw_video_callback", uintptr(unsafe.Pointer(s)), r.handle, uintptr(0))
}

// Remove wraps void obs_remove_raw_video_callback(
//
//	void (*callback)(void *param, struct video_data *frame),
//	void *param
//
// ).
func (r *RawCallback) Remove() error {
	return dll.OBS("obs_remove_raw_video_callback", r.handle)
}

// Reset wraps int obs_reset_video(struct obs_video_info *ovi).
// libobs example: https://gist.github.com/fzwoch/9e925aab37238006efb1e001241509a8
func (i Info) Reset() (ResetResult, error) {
	d, err := os.Getwd()
	if err != nil {
		return ResetResultFail, errors.Wrap(err, i.Module.String())
	}

	path := filepath.Join(d, "libobs", "data", "libobs")
	println("--------------------------")
	println(path)
	println("--------------------------")
	err = core.AddDataPath(path)
	if err != nil {
		return ResetResultFail, errors.Wrap(err, i.Module.String())
	}

	_, _, err = dll.Core(i.Module.String())
	if err != nil {
		return ResetResultFail, errors.Wrap(err, i.Module.String())
	}

	c := i.c()
	defer c.free()

	r, err := dll.OBSint32("obs_reset_video", uintptr(unsafe.Pointer(c)))
	return ResetResult(r), err
}

// String returns a string representation of a ResetResult.
func (r ResetResult) String() string {
	switch r {
	case ResetResultSuccess:
		return "obs_video_reset: success"
	case ResetResultNotSupported:
		return "obs_video_reset: adapter lacks capabilities"
	case ResetResultInvalidParam:
		return "obs_video_reset: parameter is invalid"
	case ResetResultCurrentlyActive:
		return "obs_video_reset: video is currently active"
	case ResetResultFail:
		return "obs_video_reset: graphics module is not found"
	default:
		return fmt.Sprintf("obs_video_reset: unknown result %d", r)
	}
}

// IsNull returns true or false as to whether or not Type has been initialized.
func (t Type) IsNull() bool {
	return t == Null
}

// free releases allocated cgo memory from Info conversions.
func (c *infoC) free() {
	C.free(unsafe.Pointer(c.graphics_module))
	C.free(unsafe.Pointer(c))
}
