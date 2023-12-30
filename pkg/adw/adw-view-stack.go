// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <adwaita.h>
import "C"

// ViewStackClass: instance of this type is always passed by reference.
type ViewStackClass struct {
	*viewStackClass
}

// viewStackClass is the struct that's finalized.
type viewStackClass struct {
	native *C.AdwViewStackClass
}

func (v *ViewStackClass) ParentClass() *gtk.WidgetClass {
	valptr := &v.native.parent_class
	var _v *gtk.WidgetClass // out
	_v = (*gtk.WidgetClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}

// ViewStackPageClass: instance of this type is always passed by reference.
type ViewStackPageClass struct {
	*viewStackPageClass
}

// viewStackPageClass is the struct that's finalized.
type viewStackPageClass struct {
	native *C.AdwViewStackPageClass
}
