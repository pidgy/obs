package video

import (
	"testing"

	"github.com/pidgy/obs/core"
	"github.com/pidgy/obs/graphics"
	"github.com/pidgy/obs/locale"
	"github.com/pidgy/obs/profiler"
)

func TestVideo(t *testing.T) {
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

	v, err := Get()
	if err != nil {
		t.Fatal(err)
	}
	if v.IsNull() {
		t.Fatalf("unexpected nil video_t")
	}
}

func TestResetVideo(t *testing.T) {
	// err := core.Startup(locale.EnUS, "", profiler.NULL)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// defer core.Shutdown()

	// i := &Type{
	// 	FPSNumerator:   60,
	// 	FPSDenominator: 1,
	// 	BaseWidth:      1920,
	// 	BaseHeight:     1080,
	// 	OutputWidth:    1920,
	// 	OutputHeight:   1080,

	// 	Format: FormatRGBA,

	// 	Adapter: 0,

	// 	GPUConversion: false,

	// 	Colorspace: ColorspaceDefault,
	// 	Range:      RangeDefault,

	// 	Scale: graphics.Bilinear,
	// }

	// err = Reset(i)
	// if err != nil {
	// 	t.Fatal(err)
	// }
}

func TestAddRemoveRawVideoCallback(t *testing.T) {
	err := core.Startup(locale.EnUS, "", profiler.Null)
	if err != nil {
		t.Fatal(err)
	}
	defer core.Shutdown()

	raw := NewRawCallback(
		func(d *Data) {
			t.Logf("raw video callback: %t", d == nil)
		},
	)

	err = raw.Add(&ScaleInfo{
		Format:     FormatRGBA,
		Width:      1920,
		Height:     1080,
		Colorspace: ColorspaceDefault,
		Range:      RangeDefault,
	})
	if err != nil {
		t.Fatal(err)
	}

	err = raw.Remove()
	if err != nil {
		t.Fatal(err)
	}
}
