package dshow

import (
	"testing"

	"github.com/pkg/errors"

	"github.com/pidgy/obs/core"
	"github.com/pidgy/obs/encoder"
	"github.com/pidgy/obs/graphics"
	"github.com/pidgy/obs/locale"
	"github.com/pidgy/obs/module"
	"github.com/pidgy/obs/profiler"
)

func TestModuleDShowClose(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.Null)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	p, err := New()
	if err != nil {
		t.Fatal(err)
	}

	err = p.Close()
	if err != nil && !errors.Is(err, module.ErrNotImplemented) {
		t.Fatal(err)
	}
}

func TestModuleDShow(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.Null)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	err = graphics.Enter()
	if err != nil {
		t.Fatal(err)
	}
	defer graphics.Leave()

	p, err := New()
	if err != nil {
		t.Fatal(err)
	}
	defer p.Close()

	d, err := p.Description()
	if err != nil {
		t.Fatal(err)
	}

	if d != "Windows DirectShow source/encoder" {
		t.Fatalf("unexpected description: %s", d)
	}

	v, err := p.Version()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(v)

	ids, err := encoder.EnumTypes()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("encoders: %d", len(ids))

	for _, id := range ids {
		t.Logf("encoder: %s", id)
	}

	e, err := encoder.Enum()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("encoders: %d", len(e))

}
