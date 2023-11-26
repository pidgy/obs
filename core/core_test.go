package core

import (
	"testing"

	"github.com/pidgy/obs/locale"
	"github.com/pidgy/obs/profiler"
	"github.com/pidgy/obs/uptr"
)

func TestCoreStartupShowdown(t *testing.T) {
	err := Startup(locale.EnUS, "", profiler.Null)
	if err != nil {
		t.Fatal(err)
	}

	err = Shutdown()
	if err != nil {
		t.Fatal(err)
	}
}

func TestCoreLocale(t *testing.T) {
	err := Startup(locale.EnUS, "", profiler.Null)
	if err != nil {
		t.Fatal(err)
	}
	defer Shutdown()

	l, err := Locale()
	if err != nil {
		t.Fatal(err)
	}

	if l != locale.EnUS {
		t.Fatalf("unexpected locale: %s", l)
	}
}

func TestCoreVersion(t *testing.T) {
	err := Startup(locale.EnUS, "", profiler.Null)
	if err != nil {
		t.Fatal(err)
	}
	defer Shutdown()

	s, u, err := Version()
	if err != nil {
		t.Fatal(err)
	}

	if s == "" {
		t.Fatalf("unexpected version: %s", s)
	}

	t.Logf("obs version: %s (%d -> %s)", s, u, uptr.Version(u))
}
