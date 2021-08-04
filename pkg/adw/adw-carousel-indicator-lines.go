// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #cgo pkg-config: libadwaita-1
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <adwaita.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.adw_carousel_indicator_lines_get_type()), F: marshalCarouselIndicatorLinesser},
	})
}

// CarouselIndicatorLines lines indicator for adw.Carousel.
//
// The AdwCarouselIndicatorLines widget shows a set of lines for each page of a
// given adw.Carousel. The carousel's active page is shown as another line that
// moves between them to match the carousel's position.
//
// See also adw.CarouselIndicatorDots.
//
//
// CSS nodes
//
// AdwCarouselIndicatorLines has a single CSS node with name
// carouselindicatorlines.
type CarouselIndicatorLines struct {
	gtk.Widget

	gtk.Orientable
	*externglib.Object
}

func wrapCarouselIndicatorLines(obj *externglib.Object) *CarouselIndicatorLines {
	return &CarouselIndicatorLines{
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
			Object: obj,
		},
		Orientable: gtk.Orientable{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalCarouselIndicatorLinesser(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCarouselIndicatorLines(obj), nil
}

// NewCarouselIndicatorLines creates a new AdwCarouselIndicatorLines.
func NewCarouselIndicatorLines() *CarouselIndicatorLines {
	var _cret *C.GtkWidget // in

	_cret = C.adw_carousel_indicator_lines_new()

	var _carouselIndicatorLines *CarouselIndicatorLines // out

	_carouselIndicatorLines = wrapCarouselIndicatorLines(externglib.Take(unsafe.Pointer(_cret)))

	return _carouselIndicatorLines
}

// Carousel gets the displayed carousel.
func (self *CarouselIndicatorLines) Carousel() *Carousel {
	var _arg0 *C.AdwCarouselIndicatorLines // out
	var _cret *C.AdwCarousel               // in

	_arg0 = (*C.AdwCarouselIndicatorLines)(unsafe.Pointer(self.Native()))

	_cret = C.adw_carousel_indicator_lines_get_carousel(_arg0)

	var _carousel *Carousel // out

	if _cret != nil {
		_carousel = wrapCarousel(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _carousel
}

// SetCarousel sets the displayed carousel.
func (self *CarouselIndicatorLines) SetCarousel(carousel *Carousel) {
	var _arg0 *C.AdwCarouselIndicatorLines // out
	var _arg1 *C.AdwCarousel               // out

	_arg0 = (*C.AdwCarouselIndicatorLines)(unsafe.Pointer(self.Native()))
	if carousel != nil {
		_arg1 = (*C.AdwCarousel)(unsafe.Pointer(carousel.Native()))
	}

	C.adw_carousel_indicator_lines_set_carousel(_arg0, _arg1)
}
