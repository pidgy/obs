package source

import (
	"testing"

	"github.com/pidgy/obs/core"
	"github.com/pidgy/obs/locale"
	"github.com/pidgy/obs/profiler"
)

func TestRegisterSourceInfo(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.Null)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	err = Register(
		&InfoOptions{
			ID:    "unitehud_capture_source",
			Flags: InfoOutputFlagAsyncVideo,
			Type:  InfoTypeInput,
		},
	)
	if err != nil {
		t.Fatal(err)
	}
}
