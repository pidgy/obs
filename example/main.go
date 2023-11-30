package main

import (
	"time"

	"github.com/pidgy/obs/call"
	"github.com/pidgy/obs/core"
	"github.com/pidgy/obs/data"
	"github.com/pidgy/obs/locale"
	"github.com/pidgy/obs/module"
	"github.com/pidgy/obs/module/dshow"
	"github.com/pidgy/obs/profiler"
	"github.com/pidgy/obs/source"
	"github.com/pidgy/obs/util/block"
	"github.com/pidgy/obs/util/can"
)

/*
info: Activate device 'AVerMedia HD Capture GC573 1:\\?\pci#ven_1461&dev_0054&subsys_57301461&rev_00#4&798c2c0&0&00dc#{65e8773d-8f56-11d0-a3b9-00a0c9223196}\{adef4cb5-1401-4177-84ee-fe8b26c13a5b}
info: Activate device 'AVerMedia HD Capture GC573 1:\\?\pci#ven_1461&dev_0054&subsys_57301461&rev_00#4&798c2c0&0&00dc#{65e8773d-8f56-11d0-a3b9-00a0c9223196}\{adef4cb5-1401-4177-84ee-fe8b26c13a5b}'
*/

const (
	deviceIndex = 0
)

func main() {
	v, err := dshow.NewDevice(deviceIndex)
	if err != nil {
		println(err.Error())
		return
	}

	err = core.Startup(locale.EnUS, "", profiler.Null)
	if err != nil {
		println(err.Error())
		return
	}
	defer can.Panic(core.Shutdown)

	mod, err := module.New("win-dshow")
	if err != nil {
		println(err.Error())
		return
	}

	dsc, err := mod.Description()
	if err != nil {
		println(err.Error())
		return
	}

	println("description:", dsc)

	settings, err := data.New()
	if err != nil {
		println(err.Error())
		return
	}
	defer can.Release(settings)

	config := &dshow.Settings{
		Active:          true,
		Name:            v.Name,
		VideoDeviceID:   v.ID,
		AudioOutputMode: dshow.DirectSound,

		HWDecode:  true,
		Buffering: dshow.BufferingOn,

		DeactivateWhenNotShowing: false,
	}
	err = settings.Set(config)
	if err != nil {
		println(err.Error())
		return
	}

	src, err := source.New("dshow_input", "UniteHUD Capture", settings)
	if err != nil {
		println(err.Error())
		return
	}
	defer can.Release(src)

	block.For(time.Millisecond * 100)

	h, err := src.ProcHandler()
	if err != nil {
		println(err.Error())
		return
	}
	println("proc:", h.IsNull())

	c, err := call.New()
	if err != nil {
		println(err.Error())
		return
	}

	err = c.Bool("activate", true)
	if err != nil {
		println(err.Error())
		return
	}
	defer can.Panic(c.Free)

	ok, err := h.Call("void activate(bool active)", c)
	if err != nil {
		println(err.Error())
		return
	}

	println("callback:", ok)

	println("blocking... (Ctrl+C to exit)")
	block.For(time.Minute)
}

/*
	v, err := dshow.NewDevice(0)
	if err != nil {
		println(err.Error()); return
	}

	// Instantiate the libobs core.
	err = core.Startup(locale.EnUS, "", profiler.Null)
	if err != nil {
		println(err.Error()); return
	}
	defer can.Panic(core.Shutdown)

	// Load the win-dshow module.
	mod, err := module.Load("win-dshow")
	if err != nil {
		println(err.Error()); return
	}
	if mod.IsNull() {
		panic("win-dshow failed to load")
	}

	// Define settings for a new Video Capture Device source.
	settings, err := data.New()
	if err != nil {
		println(err.Error()); return
	}
	if settings.IsNull() {
		panic("settings are null")
	}
	defer can.Release(settings)

	config := &dshow.Settings{
		Name:            v.Name,
		VideoDeviceID:   v.ID,
		AudioOutputMode: dshow.DirectSound,

		Resolution: "1920x1080",
		Res:        dshow.Preferred,
	}

	err = settings.Set(config)
	if err != nil {
		println(err.Error()); return
	}

	src, err := source.New("dshow_input", config.VideoDeviceID, data.Null, data.Null)
	if err != nil {
		println(err.Error()); return
	}
	defer can.Release(src)

	println("source:", can.S(src.Name))

	block.For(time.Second * 2)
}
*/
