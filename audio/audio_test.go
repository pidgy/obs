package audio

import (
	"testing"

	"github.com/pidgy/obs/core"
	"github.com/pidgy/obs/locale"
	"github.com/pidgy/obs/profiler"
)

func TestGetAudio(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.Null)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	a, err := Get()
	if err != nil {
		t.Fatal(err)
	}
	if a == 0 {
		t.Fatalf("unexpected nil audio_t")
	}
}

func TestAudioMonitoringDeviceGetSet(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.Null)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	m, err := MonitoringDevice()
	if err != nil {
		t.Fatal(err)
	}

	if m.Name != "" || m.ID != "" {
		t.Fatalf("expected unset monitoring device")
	}

	d, err := EnumMonitoringDevices()
	if err != nil {
		t.Fatal(err)
	}

	for _, m := range d {
		t.Logf("setting %s", m)

		ok, err := m.Set()
		if err != nil {
			t.Fatal(err)
		}
		if !ok {
			t.Fatalf("failed to set %s", m)
		}
	}
}

func TestAudioMonitoringAvailable(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.Null)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	ok, err := MonitoringAvailable()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("monitoring available: %t", ok)
}

func TestEnumAudioMonitoringDevices(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.Null)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	devices, err := EnumMonitoringDevices()
	if err != nil {
		t.Fatal(err)
	}

	for _, device := range devices {
		t.Logf("device: %s", device)
	}
}
