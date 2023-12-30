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

// WindowClass: instance of this type is always passed by reference.
type WindowClass struct {
	*windowClass
}

// windowClass is the struct that's finalized.
type windowClass struct {
	native *C.AdwWindowClass
}

func (w *WindowClass) ParentClass() *gtk.WindowClass {
	valptr := &w.native.parent_class
	var _v *gtk.WindowClass // out
	_v = (*gtk.WindowClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
