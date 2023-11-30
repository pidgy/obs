package source

import (
	"testing"

	"github.com/pidgy/obs/core"
	"github.com/pidgy/obs/data"
	"github.com/pidgy/obs/frame"
	"github.com/pidgy/obs/locale"
	"github.com/pidgy/obs/profiler"
)

func TestSourceCreateRelease(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.Null)
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
	s, err := New("game_capture", "gameplay", video)
	if err != nil {
		t.Fatal(err)
	}

	err = s.Release()
	if err != nil {
		t.Fatal(err)
	}
}

func TestOutputSource(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.Null)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	s, err := Output(0)
	if err != nil {
		t.Fatal(err)
	}

	if s.IsNull() {
		t.Fatalf("source: null")
	}

	n, err := s.Name()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("source: %s", n)

	err = s.Release()
	if err != nil {
		t.Fatal(err)
	}
}

func TestEnumSources(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.Null)
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

	s, err := New("game_capture", "gameplay", video)
	if err != nil {
		t.Fatal(err)
	}

	ss, err := Sources()
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

	ss, err = Sources()
	if err != nil {
		t.Fatal(err)
	}

	if len(ss) != 0 {
		t.Fatal(err)
	}
}

func TestSourceSetGetName(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.Null)
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

	s, err := New("game_capture", "gameplay", video)
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
	err := core.Startup(locale.EnUS, "", profiler.Null)
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

	s, err := New("game_capture", "gameplay", video)
	if err != nil {
		t.Fatal(err)
	}
	defer s.Release()

	ok, err := s.Scene()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("scene: %t", ok)
}

func TestSourceID(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.Null)
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

	s, err := New("game_capture", "gameplay", video)
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
	err := core.Startup(locale.EnUS, "", profiler.Null)
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

	s, err := New("video", "wtf", video)
	if err != nil {
		t.Fatal(err)
	}
	defer s.Release()

	ss, err := Sources()
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

	err = s.SetVolume(MinVolumeDb)
	if err != nil {
		t.Fatal(err)
	}

	v, err := s.Volume()
	if err != nil {
		t.Fatal(err)
	}

	if v != MinVolumeDb {
		t.Fatalf("unexpected volume: %.2f", v)
	}

	err = s.SetVolume(MaxVolumeDb)
	if err != nil {
		t.Fatal(err)
	}

	v, err = s.Volume()
	if err != nil {
		t.Fatal(err)
	}

	if v != MaxVolumeDb {
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
	err := core.Startup(locale.EnUS, "", profiler.Null)
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

	s, err := New("game_capture", "gameplay", video)
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
	err := core.Startup(locale.EnUS, "", profiler.Null)
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

	s, err := New("1", "gameplay", data.Null)
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
	err := core.Startup(locale.EnUS, "", profiler.Null)
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

	s, err := New("1", "gameplay", data.Null)
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

func TestEnumSourceTypes(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.Null)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	ids, err := Types()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("sources: %d", len(ids))

	for _, id := range ids {
		t.Logf("source: %s", id)
	}
}
