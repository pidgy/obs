package main

import (
	"time"

	"github.com/pidgy/obs/audio"
	"github.com/pidgy/obs/core"
	"github.com/pidgy/obs/data"
	"github.com/pidgy/obs/graphics"
	"github.com/pidgy/obs/locale"
	"github.com/pidgy/obs/module"
	"github.com/pidgy/obs/module/dshow"
	"github.com/pidgy/obs/profiler"
	"github.com/pidgy/obs/source"
	"github.com/pidgy/obs/util/block"
	"github.com/pidgy/obs/util/can"
	"github.com/pidgy/obs/video"
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

	println("--------------------------------")
	println("device_name:", v.Name)
	println("device_id:", v.Path)
	println("--------------------------------")

	err = core.Startup(locale.EnUS, "", profiler.Null)
	if err != nil {
		panic(err)
	}
	defer can.Panic(core.Shutdown)

	result, err := video.Info{
		Module:        graphics.ModuleD3D11,
		FPS:           video.FPS60,
		Base:          video.Resolution1920x1080,
		Output:        video.Resolution1920x1080,
		Format:        video.FormatYUY2,
		Adapter:       0,
		GPUConversion: true,
		ColorSpace:    video.ColorspaceDefault,
		Range:         video.RangeDefault,
		Scale:         graphics.ScaleBicubic,
	}.Reset()
	if err != nil {
		panic(err)
	}
	if result != video.ResetResultSuccess {
		panic(result.String())
	}

	ok, err := audio.Info{
		SamplesPerSecond: audio.SamplesPerSecond44p1kHz,
		SpeakerLayout:    audio.SpeakerLayoutStereo,
	}.Reset()
	if err != nil {
		panic(err)
	}
	if !ok {
		println("failed to reset audio")
	}

	version, _, err := core.Version()
	if err != nil {
		panic(err)
	}

	println("--------------------------------")
	println("obs_version:", version)
	println("--------------------------------")

	err = graphics.Enter()
	if err != nil {
		panic(err)
	}
	defer can.Panic(graphics.Leave)

	println("--------------------------------")
	println("graphics_enter")
	println("--------------------------------")

	_, err = module.New("win-dshow")
	if err != nil {
		panic(err)
	}

	settings, err := data.New()
	if err != nil {
		panic(err)
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
		panic(err)
	}

	src, err := source.New("dshow_input", v.ID, settings)
	if err != nil {
		panic(err)
	}
	defer can.Release(src)

	block.For(time.Second * 5)

	err = src.VideoRender()
	if err != nil {
		panic(err)
	}

	// println(p.IsNull())

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
