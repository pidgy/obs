package obs

import (
	"testing"

	"github.com/pidgy/obs/core"
	"github.com/pidgy/obs/data"
	"github.com/pidgy/obs/graphics"
	"github.com/pidgy/obs/locale"
	"github.com/pidgy/obs/profiler"
	"github.com/pidgy/obs/source"
)

func TestCreateSource(t *testing.T) {
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

	capture, err := source.Create("game_capture", "gameplay", video, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer capture.Release()

	src, err := source.NewInfo(
		"unitehud_capture_source",
		source.InfoTypeInput,
		source.InfoOutputAsync|source.InfoOutputCustomDraw,
	).Create(
		video,
		capture,
	)
	if err != nil {
		t.Fatal(err)
	}
	defer src.Free()
	defer src.Destroy()

	println("id:", src.ID())
	println("name:", src.Name())
	println("width:", src.Width())
	println("height:", src.Height())

	src.VideoRender(graphics.Effect(0))

	src.VideoTick(5)

	err = src.Register()
	if err != nil {
		t.Fatal(err)
	}

	ts, err := source.EnumTypes()
	if err != nil {
		t.Fatal(err)
	}

	println("sources", len(ts))

	for _, t := range ts {
		println("type", t)
	}
}
