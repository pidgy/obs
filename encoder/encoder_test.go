package encoder

import (
	"testing"

	"github.com/pidgy/obs/core"
	"github.com/pidgy/obs/locale"
	"github.com/pidgy/obs/profiler"
)

func TestEncoderEnum(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.Null)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	e, err := Enum()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("encoders: %d", len(e))
}

func TestEncoderEnumTypes(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.Null)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	ids, err := EnumTypes()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("encoders: %d", len(ids))

	for _, id := range ids {
		t.Logf("encoder: %s", id)
	}
}
