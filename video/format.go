package video

type (
	// ColorSpace wraps video_colorspace.
	ColorSpace int

	// Format wraps video_format.
	Format int

	// Range wraps video_range_type.
	Range int
)

const (
	FormatNone Format = iota
	FormatI420
	FormatNV12
	FormatYVYU
	FormatYUY2
	FormatUYVY
	FormatRGBA
	FormatBGRA
	FormatBGRX
	FormatY800
	FormatI444
	FormatBGR3
	FormatI422
	FormatI40A
	FormatI42A
	FormatYUVA
	FormatAYUV
)
const (
	ColorspaceDefault ColorSpace = iota
	Colorspace601
	Colorspace709
	ColorspaceSRGB
)
const (
	RangeDefault Range = iota
	RangePartial
	RangeFull
)

// Int returns the integer representation of a Colorspace.
func (c ColorSpace) Int() int {
	return int(c)
}

// String returns the string representation of a Colorspace.
func (c ColorSpace) String() string {
	switch c {
	case Colorspace601:
		return "601"
	case Colorspace709:
		return "709"
	case ColorspaceSRGB:
		return "SRGB"
	default:
		return ""
	}
}

// Int returns the integer representation of a Format.
func (f Format) Int() int {
	return int(f)
}

// Int returns the integer representation of a Range.
func (r Range) Int() int {
	return int(r)
}

// String returns the string representation of a Range.
func (r Range) String() string {
	switch r {
	case RangeFull:
		return "full"
	case RangePartial:
		return "partial"
	default:
		return ""
	}
}
