package video

// #cgo windows CFLAGS: -I ../libobs/include -I libobs/include -I libobs
//
// #include "info.h"
//
// static bool ctrue() { return 1; }
// static bool cfalse() { return 0; }
import "C"

import (
	"github.com/pidgy/obs/graphics"
)

type (
	// Info wraps obs_source_info... clearly.
	Info C.struct_obs_video_info

	InfoOptions struct {
		Adapter uint32

		Base,
		Output Resolution

		FPS

		Format
		ColorSpace
		Range

		graphics.Module
		graphics.Scale

		GPUConversion bool
	}

	FPS struct {
		n, d uint32
	}

	Resolution struct {
		w, h uint32
	}
)

var (
	FPS60 = FPS{60, 1}

	Resolution1920x1080 = Resolution{1920, 1080}
)

// NewInfo wraps the creation of obs_source_info.
func NewInfo(o InfoOptions) *Info {
	return &Info{
		graphics_module: C.CString(o.Module.String()),

		fps_num: C.uint32_t(o.FPS.n),
		fps_den: C.uint32_t(o.FPS.d),

		base_width:  C.uint32_t(o.Base.w),
		base_height: C.uint32_t(o.Base.h),

		output_width:  C.uint32_t(o.Output.w),
		output_height: C.uint32_t(o.Output.h),
		output_format: uint32(o.Format),

		adapter: C.uint32_t(o.Adapter),

		gpu_conversion: boolean(o.GPUConversion),

		colorspace: uint32(o.ColorSpace),
		_range:     uint32(o.Range),

		scale_type: uint32(o.Scale),
	}
}

func boolean(ok bool) C.bool {
	if ok {
		return C.ctrue()
	}
	return C.cfalse()
}
