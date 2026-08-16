package gstvideo

// #cgo pkg-config: gstreamer-video-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gst/video/video.h>
import "C"
import (
	"unsafe"

	"github.com/go-gst/go-gst/pkg/gst"
)

// VideoMeta wraps GstVideoMeta
//
// see also https://gstreamer.freedesktop.org/documentation/video/gstvideometa.html#GstVideoMeta
type VideoMeta struct {
	*videoMeta
}

// videoMeta is the internal struct holding the C pointer
type videoMeta struct {
	native *C.GstVideoMeta
}

// Instance returns the underlying C pointer.
func (v *VideoMeta) Instance() *C.GstVideoMeta {
	if v == nil {
		return nil
	}
	return v.native
}

// instance returns the underlying C pointer. This is used by the bindings internally.
func (v *VideoMeta) instance() *C.GstVideoMeta {
	if v == nil {
		return nil
	}
	return v.native
}

// Flags returns the video frame flags of the metadata.
func (v *VideoMeta) Flags() VideoFrameFlags {
	return VideoFrameFlags(v.native.flags)
}

// Format returns the video format of the metadata.
func (v *VideoMeta) Format() VideoFormat {
	return VideoFormat(v.native.format)
}

// ID returns the metadata ID.
func (v *VideoMeta) ID() int32 {
	return int32(v.native.id)
}

// Width returns the width of the video.
func (v *VideoMeta) Width() uint {
	return uint(v.native.width)
}

// Height returns the height of the video.
func (v *VideoMeta) Height() uint {
	return uint(v.native.height)
}

// NPlanes returns the number of planes.
func (v *VideoMeta) NPlanes() uint {
	return uint(v.native.n_planes)
}

// Offset returns the plane offsets.
func (v *VideoMeta) Offset() [4]uint {
	return [4]uint{
		uint(v.native.offset[0]),
		uint(v.native.offset[1]),
		uint(v.native.offset[2]),
		uint(v.native.offset[3]),
	}
}

// Stride returns the plane strides.
func (v *VideoMeta) Stride() [4]int32 {
	return [4]int32{
		int32(v.native.stride[0]),
		int32(v.native.stride[1]),
		int32(v.native.stride[2]),
		int32(v.native.stride[3]),
	}
}

// Alignment returns the alignment info of the metadata.
func (v *VideoMeta) Alignment() *VideoAlignment {
	return UnsafeVideoAlignmentFromGlibBorrow(unsafe.Pointer(&v.native.alignment))
}

// VideoMetaGetInfo wraps gst_video_meta_get_info
func VideoMetaGetInfo() *gst.MetaInfo {
	cret := C.gst_video_meta_get_info()
	if cret == nil {
		return nil
	}
	return gst.UnsafeMetaInfoFromGlibNone(unsafe.Pointer(cret))
}

// UnsafeVideoMetaFromGlibBorrow is used to convert raw C.GstVideoMeta pointers to go. This is used by the bindings internally.
func UnsafeVideoMetaFromGlibBorrow(p unsafe.Pointer) *VideoMeta {
	if p == nil {
		return nil
	}
	return &VideoMeta{&videoMeta{(*C.GstVideoMeta)(p)}}
}

// UnsafeVideoMetaFromGlibNone is used to convert raw C.GstVideoMeta pointers to go without transferring ownership.
// Since VideoMeta is owned by the parent GstBuffer, it is safely borrowed without attaching a standalone finalizer.
func UnsafeVideoMetaFromGlibNone(p unsafe.Pointer) *VideoMeta {
	return UnsafeVideoMetaFromGlibBorrow(p)
}

// UnsafeVideoMetaFromGlibFull is used to convert raw C.GstVideoMeta pointers to go while taking ownership.
func UnsafeVideoMetaFromGlibFull(p unsafe.Pointer) *VideoMeta {
	return UnsafeVideoMetaFromGlibBorrow(p)
}

// UnsafeVideoMetaFree is a no-op because VideoMeta lifecycle is tied to the parent GstBuffer.
func UnsafeVideoMetaFree(v *VideoMeta) {
}

// UnsafeVideoMetaToGlibNone returns the underlying C pointer. This is used by the bindings internally.
func UnsafeVideoMetaToGlibNone(v *VideoMeta) unsafe.Pointer {
	if v == nil {
		return nil
	}
	return unsafe.Pointer(v.native)
}

// UnsafeVideoMetaToGlibFull returns the underlying C pointer and gives up ownership.
func UnsafeVideoMetaToGlibFull(v *VideoMeta) unsafe.Pointer {
	if v == nil {
		return nil
	}
	_p := unsafe.Pointer(v.native)
	v.native = nil
	return _p
}
