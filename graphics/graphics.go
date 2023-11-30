package graphics

import (
	"syscall"

	"github.com/pkg/errors"

	"github.com/pidgy/obs/dll"
)

type (
	// Scale wraps obs_scale_type.
	Scale int

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
	_, _, err := dll.OBS.NewProc("obs_enter_graphics").Call()
	if err != syscall.Errno(0) {
		return errors.Wrap(err, "obs_enter_graphics")
	}
	return nil
}

// Leave wraps obs_leave_graphics.
func Leave() error {
	_, _, err := dll.OBS.NewProc("obs_leave_graphics").Call()
	if err != syscall.Errno(0) {
		return errors.Wrap(err, "obs_leave_graphics")
	}
	return nil
}

// MustLeave wraps obs_leave_graphics.
func MustLeave() {
	_, _, err := dll.OBS.NewProc("obs_leave_graphics").Call()
	if err != syscall.Errno(0) {
		panic(errors.Wrap(err, "obs_leave_graphics"))
	}
}
