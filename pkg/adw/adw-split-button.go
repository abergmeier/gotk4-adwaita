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

// SplitButtonClass: instance of this type is always passed by reference.
type SplitButtonClass struct {
	*splitButtonClass
}

// splitButtonClass is the struct that's finalized.
type splitButtonClass struct {
	native *C.AdwSplitButtonClass
}

func (s *SplitButtonClass) ParentClass() *gtk.WidgetClass {
	valptr := &s.native.parent_class
	var _v *gtk.WidgetClass // out
	_v = (*gtk.WidgetClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
