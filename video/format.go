package video

// Format wraps video_format.
type Format int32

const (
	FormatNone Format = iota
	/* planar 420 format */
	FormatI420 /* three-plane */
	FormatNV12 /* two-plane luma and packed chroma */

	/* packed 422 formats */
	FormatYVYU
	FormatYUY2 /* YUYV */
	FormatUYVY

	/* packed uncompressed formats */
	FormatRGBA
	FormatBGRA
	FormatBGRX
	FormatY800 /* grayscale */

	/* planar 4:4:4 */
	FormatI444

	/* more packed uncompressed formats */
	FormatBGR3

	/* planar 4:2:2 */
	FormatI422

	/* planar 4:2:0 with alpha */
	FormatI40A

	/* planar 4:2:2 with alpha */
	FormatI42A

	/* planar 4:4:4 with alpha */
	FormatYUVA

	/* packed 4:4:4 with alpha */
	FormatAYUV

	ColorspaceDefault Colorspace = iota
	Colorspace601
	Colorspace709
	ColorspaceSRGB

	RangeDefault Range = iota
	RangePartial
	RangeFull
)
