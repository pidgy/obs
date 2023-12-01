package graphics

import (
	"github.com/pidgy/obs/dll"
)

type (
	// Scale wraps obs_scale_type.
	Scale uint32

	// Effect wraps gs_effect_t.
	Effect uintptr

	// Module wraps graphics_module argument for video_info_t.
	Module string
)

const (
	ScaleDisable Scale = iota
	ScalePoint
	ScaleBicubic
	ScaleBilinear
	ScaleLanczos
	ScaleArea

	ModuleOpenGL Module = "libobs-opengl"
	ModuleD3D11  Module = "libobs-d3d11"
)

func (m Module) String() string {
	switch m {
	case ModuleOpenGL:
		return "libobs-opengl"
	case ModuleD3D11:
		return "libobs-d3d11"
	default:
		return "unknown-module"
	}
}

// Enter wraps obs_enter_graphics.
func Enter() error {
	return dll.OBS("obs_enter_graphics")
}

// Leave wraps obs_leave_graphics.
func Leave() error {
	return dll.OBS("obs_leave_graphics")
}

// Uint32 returns Scale as a uint32 type.
func (s Scale) Uint32() uint32 {
	return uint32(s)
}
