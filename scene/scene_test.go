package scene

import (
	"testing"

	"github.com/pidgy/obs/core"
	"github.com/pidgy/obs/locale"
	"github.com/pidgy/obs/profiler"
)

func TestSceneNew(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.Null)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	scene, err := New("test-scene")
	if err != nil {
		t.Fatal(err)
	}

	if scene.IsNull() {
		t.Fatalf("scene is null")
	}

	items, err := scene.Items()
	if err != nil {
		t.Fatal(err)
	}
	if len(items) != 0 {
		t.Fatalf("unexpected # items: %d", len(items))
	}
}
