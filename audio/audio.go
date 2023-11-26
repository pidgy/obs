package audio

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/pkg/errors"

	"github.com/pidgy/obs/dll"
	"github.com/pidgy/obs/uptr"
)

type (
	// Type wraps audio_t
	Type uintptr

	// MonitorDevice wraps parameters used in *_audio_monitoring_device* calls.
	MonitorDevice struct {
		Name,
		ID string
	}
)

// Get wraps audio_t *obs_get_audio(void).
func Get() (Type, error) {
	r, _, err := dll.OBS.NewProc("obs_get_audio").Call()
	if err != syscall.Errno(0) {
		return 0, errors.Wrap(err, "obs_get_audio")
	}
	return (Type)(unsafe.Pointer(&r)), nil
}

// Set wraps bool obs_set_audio_monitoring_device(const char *name, const char *id).
func (m *MonitorDevice) Set() (bool, error) {
	r, _, err := dll.OBS.NewProc("obs_set_audio_monitoring_device").Call(
		uptr.FromString(m.Name),
		uptr.FromString(m.ID),
	)
	if err != syscall.Errno(0) {
		return false, errors.Wrap(err, "obs_set_audio_monitoring_device")
	}

	return uptr.Bool(r), nil
}

// String returns the string representation of a MonitorDevice.
func (m *MonitorDevice) String() string {
	return fmt.Sprintf("Audio: %s / %s", m.Name, m.ID)
}

// MonitoringDevice wraps void obs_get_audio_monitoring_device(const char **name, const char **id).
func MonitoringDevice() (*MonitorDevice, error) {
	m := &MonitorDevice{}

	_, _, err := dll.OBS.NewProc("obs_get_audio_monitoring_device").Call(
		uptr.ReferenceFromString(m.Name),
		uptr.ReferenceFromString(m.ID),
	)
	if err != syscall.Errno(0) {
		return nil, errors.Wrap(err, "obs_get_audio_monitoring_device")
	}

	return m, nil
}

// MonitoringAvailable wraps bool obs_audio_monitoring_available(void)
func MonitoringAvailable() (bool, error) {
	r, _, err := dll.OBS.NewProc("obs_audio_monitoring_available").Call()
	if err != syscall.Errno(0) {
		return false, errors.Wrap(err, "obs_audio_monitoring_available")
	}

	return uptr.Bool(r), nil
}

// EnumMonitoringDevices wraps void obs_enum_audio_monitoring_devices(obs_enum_audio_device_cb cb, void *data)
func EnumMonitoringDevices() ([]*MonitorDevice, error) {
	var a []*MonitorDevice

	// bool
	// (*obs_enum_audio_device_cb)(void *data, const char *name, const char *id)
	callback := syscall.NewCallback(
		func(data, name, id uintptr) int {
			a = append(a, &MonitorDevice{
				Name: uptr.String(name),
				ID:   uptr.String(id),
			})
			return 1
		},
	)
	_, _, err := dll.OBS.NewProc("obs_enum_audio_monitoring_devices").Call(callback)
	if err != syscall.Errno(0) {
		return nil, errors.Wrap(err, "obs_enum_audio_monitoring_devices")
	}

	return a, nil
}
