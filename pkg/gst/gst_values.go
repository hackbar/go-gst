package gst

import (
	"runtime"
	"unsafe"

	"github.com/go-gst/go-glib/pkg/gobject/v2"
)

// #cgo pkg-config: gstreamer-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gst/gst.h>
import "C"

var (
	TypeValueArray = gobject.Type(C.gst_value_array_get_type())
	TypeValueList  = gobject.Type(C.gst_value_list_get_type())
)

func init() {
	gobject.RegisterGValueMarshalers([]gobject.TypeMarshaler{
		{T: TypeValueArray, F: marshalValueArray},
		{T: TypeValueList, F: marshalValueList},
	})
}

// ValueArray is a Go representation of a GStreamer value array (GST_TYPE_ARRAY). Each element is
// recursively marshaled via GoValue().
type ValueArray []any

// ValueList is a Go representation of a GStreamer value list (GST_TYPE_LIST). Each element is
// recursively marshaled via GoValue().
type ValueList []any

var _ gobject.GoValueInitializer = ValueArray(nil)
var _ gobject.GoValueInitializer = ValueList(nil)

// GoValueType implements GoValueInitializer.
func (a ValueArray) GoValueType() gobject.Type { return TypeValueArray }

// SetGoValue implements GoValueInitializer.
func (a ValueArray) SetGoValue(v *gobject.Value) {
	for _, elem := range a {
		elemValue := gobject.NewValue(elem)
		C.gst_value_array_append_value(
			(*C.GValue)(gobject.UnsafeValueToGlibNone(v)),
			(*C.GValue)(gobject.UnsafeValueToGlibNone(elemValue)),
		)
		runtime.KeepAlive(elemValue)
	}
	runtime.KeepAlive(v)
}

// GoValueType implements GoValueInitializer.
func (l ValueList) GoValueType() gobject.Type { return TypeValueList }

// SetGoValue implements GoValueInitializer.
func (l ValueList) SetGoValue(v *gobject.Value) {
	for _, elem := range l {
		elemValue := gobject.NewValue(elem)
		C.gst_value_list_append_value(
			(*C.GValue)(gobject.UnsafeValueToGlibNone(v)),
			(*C.GValue)(gobject.UnsafeValueToGlibNone(elemValue)),
		)
		runtime.KeepAlive(elemValue)
	}
	runtime.KeepAlive(v)
}

func marshalValueArray(p unsafe.Pointer) (any, error) {
	gvalue := (*C.GValue)(p)
	size := int(C.gst_value_array_get_size(gvalue))
	result := make(ValueArray, size)
	for i := range size {
		elemValue := C.gst_value_array_get_value(gvalue, C.guint(i))
		v := gobject.ValueFromNative(unsafe.Pointer(elemValue))
		result[i] = v.GoValue()
	}
	return result, nil
}

func marshalValueList(p unsafe.Pointer) (any, error) {
	gvalue := (*C.GValue)(p)
	size := int(C.gst_value_list_get_size(gvalue))
	result := make(ValueList, size)
	for i := range size {
		elemValue := C.gst_value_list_get_value(gvalue, C.guint(i))
		v := gobject.ValueFromNative(unsafe.Pointer(elemValue))
		result[i] = v.GoValue()
	}
	return result, nil
}
