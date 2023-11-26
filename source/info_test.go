package source

import (
	"testing"

	"github.com/pidgy/obs/core"
	"github.com/pidgy/obs/data"
	"github.com/pidgy/obs/locale"
	"github.com/pidgy/obs/profiler"
)

func TestSourceInfo(t *testing.T) {
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
	src, err := Create("game_capture", "gameplay", video, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer src.Release()

	s, err := NewInfo(
		"unitehud_capture_source",
		InfoTypeInput,
		InfoOutputAsync,
	).Create(
		data.Null,
		src,
	)
	if err != nil {
		t.Fatal(err)
	}

	println("id:", s.ID())
	println("name:", s.Name())
}
