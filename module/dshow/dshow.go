package dshow

import (
	"fmt"

	"github.com/pidgy/obs/array"
	"github.com/pidgy/obs/data"
	"github.com/pidgy/obs/module/dshow/win32"
	"github.com/pidgy/obs/video"
	"github.com/pkg/errors"
)

type (
	Action int

	// Res wraps enum ResType.
	Res int

	// Buffering wraps enum class BufferingType : int64_t.
	Buffering int

	// AudioOutputMode wraps enum class AudioMode.
	AudioOutputMode int

	// enum class BufferingType : int64_t {
	// 	Auto,
	// 	On,
	// 	Off,
	// };

	Settings struct {
		Active bool
		Name   string

		AudioDeviceID string
		video.ColorSpace
		video.Format
		LastVideoDeviceID string
		LastResolution    string
		video.Range
		Resolution    string
		VideoDeviceID string

		Buffering
		FrameInterval int
		Res
		AudioOutputMode

		AutoRotation             bool
		DeactivateWhenNotShowing bool
		FlipImage                bool
		HWDecode                 bool
		UseCustomAudioDevice     bool
	}
)

const (
	None Action = iota
	Activate
	ActivateBlock
	Deactivate
	Shutdown
	ConfigVideo
	ConfigAudio
	ConfigCrossbar1
	ConfigCrossbar2
)

const (
	Capture AudioOutputMode = iota
	DirectSound
	WaveOut
)

const (
	BufferingAuto Buffering = iota
	BufferingOn
	BufferingOff
)

const (
	Preferred Res = iota
	Custom
)

var (
	NewDevice = win32.NewVideoCaptureDevice
)

func (a AudioOutputMode) Int() int {
	return int(a)
}

// Int returns the integer representation of a Buffering.
func (b Buffering) Int() int {
	return int(b)
}

// Int returns the integer representation of a Res.
func (r Res) Int() int {
	return int(r)
}

func (s *Settings) Arrays() map[string]array.Type {
	return map[string]array.Type{}
}

func (s *Settings) Bools() map[string]bool {
	return map[string]bool{
		"active":                      s.Active,
		"autorotation":                s.AutoRotation,
		"deactivate_when_not_showing": s.DeactivateWhenNotShowing,
		"flip_vertically":             s.FlipImage,
		"hw_decode":                   s.HWDecode,
		"use_custom_audio_device":     s.UseCustomAudioDevice,
	}
}

func (s *Settings) Doubles() map[string]float64 {
	return map[string]float64{}
}

func (s *Settings) Ints() map[string]int {
	return map[string]int{
		"audio_output_mode": s.AudioOutputMode.Int(),
		"buffering":         s.Buffering.Int(),
		"frame_interval":    s.FrameInterval,
		"res_type":          s.Res.Int(),
		"video_format":      int(s.Format.Uint32()),
	}
}

func (s *Settings) Objects() map[string]data.Type {
	return map[string]data.Type{}
}

func (s *Settings) Strings() map[string]string {
	return map[string]string{
		"name":                 s.Name,
		"audio_device_id":      s.AudioDeviceID,
		"color_space":          s.ColorSpace.String(),
		"color_range":          s.Range.String(),
		"last_resolution":      s.LastResolution,
		"last_video_device_id": s.LastVideoDeviceID,
		"resolution":           s.Resolution,
		"video_device_id":      s.VideoDeviceID,
	}
}

// VideoDeviceID returns the concatenation of L"DeviceName" and L"DevicePath".
func VideoDevice(index int) (name, path, id string, err error) {
	name, err = win32.VideoCaptureDeviceName(index)
	if err != nil {
		return
	}

	path, err = win32.VideoCaptureDevicePath(index)
	if err != nil {
		return
	}

	id = fmt.Sprintf("%s:%s", name, path)

	if name == "" || path == "" {
		err = errors.Errorf("device %d has invalid video device id: %s", index, id)
	}

	return
}

// void obs_data_set_string(obs_data_t *data, const char *name, const char *val)
// void obs_data_set_int(obs_data_t *data, const char *name, long long val)
// void obs_data_set_double(obs_data_t *data, const char *name, double val)
// void obs_data_set_bool(obs_data_t *data, const char *name, bool val)
// void obs_data_set_obj(obs_data_t *data, const char *name, obs_data_t *obj)
// void obs_data_set_array(obs_data_t *data, const char *name, obs_data_array_t *array)
