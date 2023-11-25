package obs

import (
	"testing"

	"github.com/pidgy/obs/audio"
	"github.com/pidgy/obs/core"
	"github.com/pidgy/obs/data"
	"github.com/pidgy/obs/frame"
	"github.com/pidgy/obs/graphics"
	"github.com/pidgy/obs/locale"
	"github.com/pidgy/obs/profiler"
	"github.com/pidgy/obs/source"
	"github.com/pidgy/obs/video"
)

const (
	version = "27.5.32"
)

func TestStartupShowdown(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.NULL)
	if err != nil {
		t.Fatal(err)
	}

	err = core.Shutdown()
	if err != nil {
		t.Fatal(err)
	}
}

func TestLocale(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.NULL)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	l, err := core.Locale()
	if err != nil {
		t.Fatal(err)
	}

	if l != locale.EnUS {
		t.Fatalf("unexpected locale.EnUS: %s", l)
	}
}

func TestVersion(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.NULL)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	s, _, err := core.Version()
	if err != nil {
		t.Fatal(err)
	}

	if s != version {
		t.Fatalf("unexpected version: %s", s)
	}
}

func TestProfilerNameStore(t *testing.T) {
	n, err := profiler.New()
	if err != nil {
		t.Fatal(err)
	}

	err = core.Startup(locale.EnUS, "", n)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	name, err := n.Store("testing")
	if err != nil {
		t.Fatal(err)
	}
	if name != "testing" {
		t.Fatalf("unexpected name store: %s", name)
	}

	err = n.Close()
	if err != nil {
		t.Fatal(err)
	}

	n2, err := profiler.Get()
	if err != nil {
		t.Fatal(err)
	}
	if n2 == profiler.NULL {
		t.Fatalf("unexpected profiler name store")
	}
}

func TestGraphics(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.NULL)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	err = graphics.Enter()
	if err != nil {
		t.Fatal(err)
	}
	defer graphics.Leave()
}

func TestGetAudio(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.NULL)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	a, err := audio.Get()
	if err != nil {
		t.Fatal(err)
	}
	if a == 0 {
		t.Fatalf("unexpected nil audio_t")
	}
}

func TestVideo(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.NULL)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	v, err := video.Get()
	if err != nil {
		t.Fatal(err)
	}

	if v == nil {
		t.Fatalf("unexpected nil audio_t")
	}
}

func TestAudioMonitoringDeviceGetSet(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.NULL)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	m, err := audio.MonitoringDevice()
	if err != nil {
		t.Fatal(err)
	}

	if m.Name != "" || m.ID != "" {
		t.Fatalf("expected unset monitoring device")
	}

	d, err := audio.EnumMonitoringDevices()
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
	err := core.Startup(locale.EnUS, "", profiler.NULL)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	ok, err := audio.MonitoringAvailable()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("TestAudioMonitoringAvailable: %t", ok)
}

func TestEnumAudioMonitoringDevices(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.NULL)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	devices, err := audio.EnumMonitoringDevices()
	if err != nil {
		t.Fatal(err)
	}

	for _, device := range devices {
		t.Logf("%s", device)
	}
}

func TestAddRemoveRawVideoCallback(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.NULL)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	raw := video.NewRawCallback(
		func(d *video.Data) {
			t.Logf("raw video callback: %t", d == nil)
		},
	)

	err = raw.Add(&video.ScaleInfo{
		Format:     video.FormatRGBA,
		Width:      1920,
		Height:     1080,
		Colorspace: video.ColorspaceDefault,
		Range:      video.RangeDefault,
	})
	if err != nil {
		t.Fatal(err)
	}

	err = raw.Remove()
	if err != nil {
		t.Fatal(err)
	}
}

func TestResetVideo(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.NULL)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	i := &video.Type{
		FPSNumerator:   60,
		FPSDenominator: 1,
		BaseWidth:      1920,
		BaseHeight:     1080,
		OutputWidth:    1920,
		OutputHeight:   1080,

		Format: video.FormatRGBA,

		Adapter: 0,

		GPUConversion: false,

		Colorspace: video.ColorspaceDefault,
		Range:      video.RangeDefault,

		Scale: graphics.Bilinear,
	}

	err = video.Reset(i)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSourceCreateRelease(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.NULL)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	video, err := data.New()
	if err != nil {
		t.Fatal(err)
	}

	err = video.SetString("capture_mode", "window")
	if err != nil {
		t.Fatal(err)
	}

	err = video.SetString("window", "foo:bar:foobar.exe")
	if err != nil {
		t.Fatal(err)
	}

	// obs_source_create("game_capture", "gameplay", videoSourceSettings, IntPtr.Zero)
	s, err := source.New("game_capture", "gameplay", video, 0)
	if err != nil {
		t.Fatal(err)
	}

	err = s.Release()
	if err != nil {
		t.Fatal(err)
	}
}

func TestEnumSources(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.NULL)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	video, err := data.New()
	if err != nil {
		t.Fatal(err)
	}

	err = video.SetString("capture_mode", "window")
	if err != nil {
		t.Fatal(err)
	}

	err = video.SetString("window", "foo:bar:foobar.exe")
	if err != nil {
		t.Fatal(err)
	}

	s, err := source.New("game_capture", "gameplay", video, 0)
	if err != nil {
		t.Fatal(err)
	}

	ss, err := source.Enum()
	if err != nil {
		t.Fatal(err)
	}

	if len(ss) != 1 {
		t.Fatal(err)
	}

	err = s.Release()
	if err != nil {
		t.Fatal(err)
	}

	ss, err = source.Enum()
	if err != nil {
		t.Fatal(err)
	}

	if len(ss) != 0 {
		t.Fatal(err)
	}
}

func TestSourceSetGetName(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.NULL)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	video, err := data.New()
	if err != nil {
		t.Fatal(err)
	}

	err = video.SetString("capture_mode", "window")
	if err != nil {
		t.Fatal(err)
	}

	err = video.SetString("window", "foo:bar:foobar.exe")
	if err != nil {
		t.Fatal(err)
	}

	s, err := source.New("game_capture", "gameplay", video, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer s.Release()

	n, err := s.Name()
	if err != nil {
		t.Fatal(err)
	}
	if n != "gameplay" {
		t.Fatalf("unexpected source name: \"%s\"", n)
	}

	err = s.SetName("your mom")
	if err != nil {
		t.Fatal(err)
	}

	n, err = s.Name()
	if err != nil {
		t.Fatal(err)
	}

	if n != "your mom" {
		t.Fatalf("unexpected source name: \"%s\"", n)
	}
}

func TestSourceIsScene(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.NULL)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	video, err := data.New()
	if err != nil {
		t.Fatal(err)
	}

	err = video.SetString("capture_mode", "window")
	if err != nil {
		t.Fatal(err)
	}

	err = video.SetString("window", "foo:bar:foobar.exe")
	if err != nil {
		t.Fatal(err)
	}

	s, err := source.New("game_capture", "gameplay", video, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer s.Release()

	ok, err := s.Scene()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("source scene? %t", ok)
}

func TestSourceID(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.NULL)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	video, err := data.New()
	if err != nil {
		t.Fatal(err)
	}

	err = video.SetString("capture_mode", "window")
	if err != nil {
		t.Fatal(err)
	}

	err = video.SetString("window", "foo:bar:foobar.exe")
	if err != nil {
		t.Fatal(err)
	}

	s, err := source.New("game_capture", "gameplay", video, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer s.Release()

	id, err := s.ID()
	if err != nil {
		t.Fatal(err)
	}

	if id != "game_capture" {
		t.Fatalf("unexpected source id: \"%s\"", id)
	}
}

func TestSourceVolume(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.NULL)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	video, err := data.New()
	if err != nil {
		t.Fatal(err)
	}

	err = video.SetString("capture_mode", "window")
	if err != nil {
		t.Fatal(err)
	}

	err = video.SetString("window", "foo:bar:foobar.exe")
	if err != nil {
		t.Fatal(err)
	}

	s, err := source.New("video", "wtf", video, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer s.Release()

	ss, err := source.Enum()
	if err != nil {
		t.Fatal(err)
	}

	for _, s := range ss {
		n, err := s.Name()
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("source: %s", n)
	}

	_, err = s.Volume()
	if err != nil {
		t.Fatal(err)
	}

	err = s.SetVolume(source.MinVolumeDb)
	if err != nil {
		t.Fatal(err)
	}

	v, err := s.Volume()
	if err != nil {
		t.Fatal(err)
	}

	if v != source.MinVolumeDb {
		t.Fatalf("unexpected volume: %.2f", v)
	}

	err = s.SetVolume(source.MaxVolumeDb)
	if err != nil {
		t.Fatal(err)
	}

	v, err = s.Volume()
	if err != nil {
		t.Fatal(err)
	}

	if v != source.MaxVolumeDb {
		t.Fatalf("unexpected volume: %.2f", v)
	}

	err = s.SetVolume(-12)
	if err != nil {
		t.Fatal(err)
	}

	v, err = s.Volume()
	if err != nil {
		t.Fatal(err)
	}

	if v != -12 {
		t.Fatalf("unexpected volume: %.2f", v)
	}
}

func TestSourceMuted(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.NULL)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	video, err := data.New()
	if err != nil {
		t.Fatal(err)
	}

	err = video.SetString("capture_mode", "window")
	if err != nil {
		t.Fatal(err)
	}

	err = video.SetString("window", "foo:bar:foobar.exe")
	if err != nil {
		t.Fatal(err)
	}

	s, err := source.New("game_capture", "gameplay", video, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer s.Release()

	err = s.SetMuted(true)
	if err != nil {
		t.Fatal(err)
	}

	m, err := s.Muted()
	if err != nil {
		t.Fatal(err)
	}

	if !m {
		t.Fatal("unexpected muted value")
	}

	err = s.SetMuted(false)
	if err != nil {
		t.Fatal(err)
	}

	m, err = s.Muted()
	if err != nil {
		t.Fatal(err)
	}

	if m {
		t.Fatal("unexpected muted value")
	}
}

func TestSourceOutputVideo(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.NULL)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	video, err := data.New()
	if err != nil {
		t.Fatal(err)
	}

	err = video.SetString("capture_mode", "window")
	if err != nil {
		t.Fatal(err)
	}

	err = video.SetString("window", "foo:bar:foobar.exe")
	if err != nil {
		t.Fatal(err)
	}

	s, err := source.New("1", "gameplay", 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer s.Release()

	v := frame.Video{}

	err = s.OutputVideo(&v)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("video format: %d", v.Format())
}

func TestSourceOutputAudio(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.NULL)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	video, err := data.New()
	if err != nil {
		t.Fatal(err)
	}

	err = video.SetString("capture_mode", "window")
	if err != nil {
		t.Fatal(err)
	}

	err = video.SetString("window", "foo:bar:foobar.exe")
	if err != nil {
		t.Fatal(err)
	}

	s, err := source.New("1", "gameplay", 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer s.Release()

	a := frame.Audio{}

	err = s.OutputAudio(&a)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("audio format: %d", a.Format())
}
