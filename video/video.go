package video

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/pkg/errors"

	"github.com/pidgy/obs/dll"
)

const Null = Type(0)

type (
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
	r, _, err := dll.OBS.NewProc("obs_get_video").Call()
	if err != syscall.Errno(0) {
		return Null, errors.Wrap(err, "obs_get_video")
	}

	return Type(r), nil
}

// Add wraps void obs_add_raw_video_callback(
//
//	const struct video_scale_info *conversion,
//	void (*callback)(void *param, struct video_data *frame),
//	void *param
//
// ).
func (r *RawCallback) Add(s *ScaleInfo) error {
	_, _, err := dll.OBS.NewProc("obs_add_raw_video_callback").Call(
		uintptr(unsafe.Pointer(s)),
		r.handle,
		uintptr(0),
	)
	if err != syscall.Errno(0) {
		return errors.Wrap(err, "obs_add_raw_video_callback")
	}
	return nil
}

// Remove wraps void obs_remove_raw_video_callback(
//
//	void (*callback)(void *param, struct video_data *frame),
//	void *param
//
// ).
func (r *RawCallback) Remove() error {
	_, _, err := dll.OBS.NewProc("obs_remove_raw_video_callback").Call(
		r.handle,
	)
	if err != syscall.Errno(0) {
		return errors.Wrap(err, "obs_remove_raw_video_callback")
	}
	return nil
}

// Reset wraps int obs_reset_video(struct obs_video_info *ovi).
// libobs example: https://gist.github.com/fzwoch/9e925aab37238006efb1e001241509a8
func Reset(i *Info) error {
	r, _, err := dll.OBS.NewProc("obs_reset_video").Call(
		uintptr(unsafe.Pointer(i)),
	)
	if err != syscall.Errno(0) {
		return errors.Wrap(err, "obs_reset_video")
	}

	result := Result(r)
	if result != Success {
		return fmt.Errorf("%s", result)
	}

	return nil
}

// IsNull returns true or false as to whether or not Type has been initialized.
func (t Type) IsNull() bool {
	return t == Null
}
