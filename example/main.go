package main

import (
	"github.com/pidgy/obs/core"
	"github.com/pidgy/obs/data"
	"github.com/pidgy/obs/locale"
	"github.com/pidgy/obs/module"
	"github.com/pidgy/obs/module/dshow"
	"github.com/pidgy/obs/profiler"
	"github.com/pidgy/obs/source"
)

func main() {
	err := core.Startup(locale.EnUS, "", profiler.Null)
	if err != nil {
		panic(err)
	}
	defer core.Shutdown()

	ok, err := core.Initialized()
	if err != nil {
		panic(err)
	}
	if !ok {
		panic("core is not initialized")
	}

	mod, err := module.New("win-dshow")
	if err != nil {
		panic(err)
	}

	dsc, err := mod.Description()
	if err != nil {
		panic(err)
	}

	name, err := mod.Name()
	if err != nil {
		panic(err)
	}

	println("loaded module:", name, "-", dsc)

	d, err := data.New()
	if err != nil {
		panic(err)
	}
	defer d.Release()

	err = d.SetString(dshow.SettingVideoDeviceID, "AVerMedia HD Capture GC573 1")
	if err != nil {
		panic(err)
	}

	err = d.SetString(dshow.SettingResolution, "1920x1080")
	if err != nil {
		panic(err)
	}

	capture, err := source.Create("dshow_input", "AVerMedia HD Capture GC573 1", d, data.Null)
	if err != nil {
		panic(err)
	}
	defer capture.Release()

	ok, err = capture.Configurable()
	if err != nil {
		panic(err)
	}
	println("configurable:", ok)
}
