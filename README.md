### obs 

###### [libobs](https://docs.obsproject.com/) wrapper written in Go. 

##### Install
```
go get github.com/pidgy/obs
```

##### Dependencies
This library is developed against [libobs-windows64-release-27.5.32](https://obsstudios3.streamlabs.com/libobs-windows64-release-27.5.32.7z).

##### Testing
[obs_test.go](https://github.com/pidgy/obs/blob/main/obs_test.go) has examples of the functionality implemented in this library.

##### Startup / Shutdown
```go
err := core.Startup(locale.EnUS, "", profiler.NULL)
if err != nil {
    panic(err)
}
defer core.Shutdown()
```

##### Audio Monitoring Devices
```go
m, err := audio.MonitoringDevice()
if err != nil {
    panic(err)
}

if m.Name != "" || m.ID != "" {
    panic("expected unset monitoring device")
}

d, err := audio.EnumMonitoringDevices()
if err != nil {
    panic(err)
}

for _, m := range d {
    ok, err := m.Set()
    if err != nil {
        panic(err)
    }
    if !ok {
        panic("failed to set audio monitoring device")
    }
}
```

##### Video Capture Source
```go
video, err := data.New()
if err != nil {
    panic(err)
}

err = video.SetString("capture_mode", "window")
if err != nil {
    panic(err)
}

err = video.SetString("window", "foo:bar:foobar.exe")
if err != nil {
    panic(err)
}

s, err := source.New("game_capture", "gameplay", video, 0)
if err != nil {
    panic(err)
}

err = s.Release()
if err != nil {
    panic(err)
}
```

##### Creating a DirectShow source
```go
v, err := dshow.NewDevice(deviceIndex)
if err != nil {
	panic(err)
}

err = core.Startup(locale.EnUS, "", profiler.Null)
if err != nil {
	panic(err)
}
defer can.Panic(core.Shutdown)

mod, err := module.New("win-dshow")
if err != nil {
	panic(err)
}

dsc, err := mod.Description()
if err != nil {
	panic(err)
}
println("description:", dsc) // Prints "description: Windows DirectShow source/encoder".

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

src, err := source.New("dshow_input", "UniteHUD Capture", settings)
if err != nil {
	panic(err)
}
defer can.Release(src)

println("blocking... (Ctrl+C to exit)")
block.For(time.Minute)
```
