// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: libadwaita-1
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <adwaita.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.adw_bin_get_type()), F: marshalBinner},
	})
}

type Bin struct {
	gtk.Widget
}

var _ gextras.Nativer = (*Bin)(nil)

func wrapBin(obj *externglib.Object) *Bin {
	return &Bin{
		Widget: gtk.Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: gtk.Accessible{
				Object: obj,
			},
			Buildable: gtk.Buildable{
				Object: obj,
			},
			ConstraintTarget: gtk.ConstraintTarget{
				Object: obj,
			},
		},
	}
}

func marshalBinner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapBin(obj), nil
}

// NewBin creates a new Bin.
func NewBin() *Bin {
	var _cret *C.GtkWidget // in

	_cret = C.adw_bin_new()

	var _bin *Bin // out

	_bin = wrapBin(externglib.Take(unsafe.Pointer(_cret)))

	return _bin
}

// Child gets the child widget of self.
func (self *Bin) Child() gtk.Widgetter {
	var _arg0 *C.AdwBin    // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.AdwBin)(unsafe.Pointer(self.Native()))

	_cret = C.adw_bin_get_child(_arg0)

	var _widget gtk.Widgetter // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(gtk.Widgetter)

	return _widget
}

// SetChild sets the child widget of self.
func (self *Bin) SetChild(child gtk.Widgetter) {
	var _arg0 *C.AdwBin    // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.AdwBin)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))

	C.adw_bin_set_child(_arg0, _arg1)
}