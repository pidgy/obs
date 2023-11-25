### obs 

###### `obs` is a [libobs](https://docs.obsproject.com/) wrapper written in Go. 

##### Install
```
go get github.com/pidgy/obs
```

##### Install
This library is developed against libobs-windows64-release-27.5.32. Prebuilt DLL files can be downloaded [here](https://obsstudios3.streamlabs.com/libobs-windows64-release-27.5.32.7z).

##### Testing
`obs_test.go` tests all of the functionality implemented in this library.

##### Startup / Shutdown
```
err := core.Startup(locale.EnUS, "", profiler.NULL)
if err != nil {
    t.Fatal(err)
}
defer core.Shutdown()
```

##### Audio Monitoring Devices
```
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
```

##### Video Capture Source
```
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
```