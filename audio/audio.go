package audio

// #cgo windows CFLAGS: -I ../libobs/include -I libobs/include -I libobs
// #include "obs.h"
import "C"

import (
	"fmt"
	"syscall"
	"unsafe"

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

	Info struct {
		SamplesPerSecond
		SpeakerLayout
	}

	SpeakerLayout    uint32
	SamplesPerSecond uint32
)

const (
	SpeakerLayoutUnknown SpeakerLayout = iota
	SpeakerLayoutMono
	SpeakerLayoutStereo
	SpeakerLayout2p1
	SpeakerLayout4p0
	SpeakerLayout4p1
	SpeakerLayout5p1
	SpeakerLayout7p1
)

const (
	SamplesPerSecond44p1kHz SamplesPerSecond = 44100
)

// Get wraps audio_t *obs_get_audio(void).
func Get() (Type, error) {
	r, err := dll.OBSuintptr("obs_get_audio")
	return (Type)(unsafe.Pointer(&r)), err
}

// Reset wraps bool obs_reset_audio(const struct obs_audio_info *oai).
func (oai Info) Reset() (bool, error) {
	return dll.OBSbool("obs_reset_audio", uintptr(unsafe.Pointer(&oai)))
}

// Set wraps bool obs_set_audio_monitoring_device(const char *name, const char *id).
func (m *MonitorDevice) Set() (bool, error) {
	return dll.OBSbool("obs_set_audio_monitoring_device", uptr.FromString(m.Name), uptr.FromString(m.ID))
}

// String returns the string representation of a MonitorDevice.
func (m *MonitorDevice) String() string {
	return fmt.Sprintf("Audio: %s / %s", m.Name, m.ID)
}

// MonitoringDevice wraps void obs_get_audio_monitoring_device(const char **name, const char **id).
func MonitoringDevice() (*MonitorDevice, error) {
	m := &MonitorDevice{}
	return m, dll.OBS("obs_get_audio_monitoring_device", uptr.ReferenceFromString(m.Name), uptr.ReferenceFromString(m.ID))
}

// MonitoringAvailable wraps bool obs_audio_monitoring_available(void)
func MonitoringAvailable() (bool, error) {
	return dll.OBSbool("obs_audio_monitoring_available")
}

// EnumMonitoringDevices wraps void obs_enum_audio_monitoring_devices(obs_enum_audio_device_cb cb, void *data)
// and implements bool (*obs_enum_audio_device_cb)(void *data, const char *name, const char *id)
func EnumMonitoringDevices() ([]*MonitorDevice, error) {
	var a []*MonitorDevice
	callback := syscall.NewCallback(func(data, name, id uintptr) int {
		a = append(a, &MonitorDevice{Name: uptr.String(name), ID: uptr.String(id)})
		return 1
	})
	return a, dll.OBS("obs_enum_audio_monitoring_devices", callback)
}
