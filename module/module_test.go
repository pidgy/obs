package module

import (
	"testing"

	"github.com/pidgy/obs/core"
	"github.com/pidgy/obs/locale"
	"github.com/pidgy/obs/profiler"
)

func TestModuleNew(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.Null)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	m, err := New("win-dshow.dll")
	if err != nil {
		t.Fatal(err)
	}

	if m.IsNull() {
		t.Fatalf("module is null")
	}

	err = Log()
	if err != nil {
		t.Fatal(err)
	}
}

func TestModuleLoadAll(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.Null)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	err = LoadAll()
	if err != nil {
		t.Fatal(err)
	}

	err = Log()
	if err != nil {
		t.Fatal(err)
	}
}

func TestModuleLog(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.Null)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	err = Log()
	if err != nil {
		t.Fatal(err)
	}
}
