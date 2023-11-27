package obs

import (
	"testing"

	"github.com/pidgy/obs/core"
	"github.com/pidgy/obs/data"
	"github.com/pidgy/obs/locale"
	"github.com/pidgy/obs/module"
	"github.com/pidgy/obs/module/dshow"
	"github.com/pidgy/obs/profiler"
	"github.com/pidgy/obs/source"
)

func TestOBS(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.Null)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	ok, err := core.Initialized()
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatalf("core is not initialized")
	}

	mod, err := module.New("win-dshow")
	if err != nil {
		t.Fatal(err)
	}

	dsc, err := mod.Description()
	if err != nil {
		t.Fatal(err)
	}

	name, err := mod.Name()
	if err != nil {
		t.Fatal(err)
	}

	println("loaded module:", name, "-", dsc)

	d, err := data.New()
	if err != nil {
		t.Fatal(err)
	}
	defer d.Release()

	err = d.SetString(dshow.SettingVideoDeviceID, "AVerMedia HD Capture GC573 1")
	if err != nil {
		t.Fatal(err)
	}

	err = d.SetString(dshow.SettingResolution, "1920x1080")
	if err != nil {
		t.Fatal(err)
	}

	capture, err := source.Create("dshow_input", "AVerMedia HD Capture GC573 1", d, data.Null)
	if err != nil {
		t.Fatal(err)
	}
	defer capture.Release()

	ok, err = capture.Configurable()
	if err != nil {
		t.Fatal(err)
	}
	println("configurable:", ok)

	err = capture.VideoRender()
	if err != nil {
		t.Fatal(err)
	}

	// p, err := capture.Properties()
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// println(p.IsNull())
}
