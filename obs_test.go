package obs

import (
	"testing"
	"time"

	"github.com/pidgy/obs/core"
	"github.com/pidgy/obs/data"
	"github.com/pidgy/obs/locale"
	"github.com/pidgy/obs/module"
	"github.com/pidgy/obs/module/dshow"
	"github.com/pidgy/obs/profiler"
	"github.com/pidgy/obs/source"
	"github.com/pidgy/obs/util/block"
	"github.com/pidgy/obs/util/can"
	"github.com/pidgy/obs/video"
)

func TestOBSCaptureSource(t *testing.T) {
	device := 0

	v, err := dshow.NewDevice(device)
	if err != nil {
		t.Fatal(err)
	}

	err = core.Startup(locale.EnUS, "", profiler.Null)
	if err != nil {
		t.Fatal(err)
	}
	defer can.Panic(core.Shutdown)

	mod, err := module.New("win-dshow")
	if err != nil {
		t.Fatal(err)
	}

	settings, err := data.New()
	if err != nil {
		t.Fatal(err)
	}
	defer can.Release(settings)

	s := &dshow.Settings{
		Active:          true,
		Name:            v.Name,
		VideoDeviceID:   v.ID,
		AudioOutputMode: dshow.DirectSound,

		Resolution: "1920x1080",
		Res:        dshow.Preferred,

		DeactivateWhenNotShowing: false,

		Buffering:  dshow.BufferingOn,
		Format:     video.FormatNone,
		Range:      video.RangeDefault,
		ColorSpace: video.ColorspaceDefault,

		HWDecode:             true,
		UseCustomAudioDevice: false,
		AutoRotation:         false,
		FlipImage:            false,
	}

	err = settings.Set(s)
	if err != nil {
		t.Fatal(err)
	}

	src, err := source.New("dshow_input", s.VideoDeviceID, settings)
	if err != nil {
		t.Fatal(err)
	}
	defer can.Release(src)

	println("module:", can.S(mod.Name))
	println("source:", can.S(src.Name))

	block.For(time.Second * 10)
}
