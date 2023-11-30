package frame

// #cgo windows CFLAGS: -I ../libobs/include -I libobs/include -I libobs
// #include "obs.h"
import "C"

// MaxAVPlanes wraps MAX_AV_PLANES.
const MaxAVPlanes = 8

type (
	/*
		Audio wraps struct obs_source_audio.

		* Source audio output structure.  Used with obs_source_output_audio to output
		* source audio.  Audio is automatically resampled and remixed as necessary.

		struct obs_source_audio {
		const uint8_t *data[MAX_AV_PLANES];
		uint32_t frames;

		enum speaker_layout speakers;
		enum audio_format format;
		uint32_t samples_per_sec;

		uint64_t            timestamp;
		int64_t             dec_frame_pts;
		};
	*/
	Audio C.struct_obs_source_audio

	/*
		Video wraps obs_source_frame.

		* Source asynchronous video output structure.  Used with
		* obs_source_output_video to output asynchronous video.  Video is buffered as
		* necessary to play according to timestamps.  When used with audio output,
		* audio is synced to video as it is played.
		*
		* If a YUV format is specified, it will be automatically upsampled and
		* converted to RGB via shader on the graphics processor.
		*
		* NOTE: Non-YUV formats will always be treated as full range with this
		* structure!  Use obs_source_frame2 along with obs_source_output_video2
		* instead if partial range support is desired for non-YUV video formats.

		struct obs_source_frame {
		uint8_t *data[MAX_AV_PLANES];
		uint32_t linesize[MAX_AV_PLANES];
		uint32_t width;
		uint32_t height;
		uint64_t timestamp;
		uint64_t duration;

		enum video_format format;
		float color_matrix[16];
		bool full_range;
		float color_range_min[3];
		float color_range_max[3];
		bool flip;
		uint8_t flags;

		// Used internally by libobs.
		volatile long refs;
		bool prev_frame;
		};
	*/
	Video C.struct_obs_source_frame
)

// Format wraps access to obs_source_audio's format field.
func (a Audio) Format() uint32 {
	return a.format
}

// Format wraps access to obs_source_frame's format field.
func (v Video) Format() uint32 {
	return v.format
}

// Height wraps access to obs_source_frame's height field.
func (v Video) Height() uint32 {
	return uint32(v.height)
}

// Width wraps access to obs_source_frame's width field.
func (v Video) Width() uint32 {
	return uint32(v.width)
}
