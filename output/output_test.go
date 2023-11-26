package output

import (
	"testing"

	"github.com/pidgy/obs/core"
	"github.com/pidgy/obs/locale"
	"github.com/pidgy/obs/profiler"
)

func TestEnumOutputTypes(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.Null)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	ids, err := EnumTypes()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("outputs: %d", len(ids))

	for _, id := range ids {
		t.Logf("output: %s", id)
	}
}
