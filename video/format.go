package video

type (
	// ColorSpace wraps video_colorspace.
	ColorSpace uint32

	// Format wraps video_format.
	Format uint32

	// Range wraps video_range_type.
	Range uint32
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
func (c ColorSpace) Uint32() uint32 {
	return uint32(c)
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
func (f Format) Uint32() uint32 {
	return uint32(f)
}

// Int returns the integer representation of a Range.
func (r Range) Uint32() uint32 {
	return uint32(r)
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
