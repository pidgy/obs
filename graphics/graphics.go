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
)

const (
	Disable Scale = iota
	Point
	Bicubic
	Bilinear
	Lanczos
	Area
)

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
