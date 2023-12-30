// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <adwaita.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeCarouselIndicatorLines = coreglib.Type(C.adw_carousel_indicator_lines_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeCarouselIndicatorLines, F: marshalCarouselIndicatorLines},
	})
}

// CarouselIndicatorLinesOverrides contains methods that are overridable.
type CarouselIndicatorLinesOverrides struct {
}

func defaultCarouselIndicatorLinesOverrides(v *CarouselIndicatorLines) CarouselIndicatorLinesOverrides {
	return CarouselIndicatorLinesOverrides{}
}

// CarouselIndicatorLines lines indicator for carousel.
//
// <picture> <source srcset="carousel-indicator-dots-lines.png"
// media="(prefers-color-scheme: dark)"> <img src="carousel-indicator-lines.png"
// alt="carousel-indicator-lines"> </picture>
//
// The AdwCarouselIndicatorLines widget shows a set of lines for each page of
// a given carousel. The carousel's active page is shown as another line that
// moves between them to match the carousel's position.
//
// See also carouselindicatordots.
//
// # CSS nodes
//
// AdwCarouselIndicatorLines has a single CSS node with name
// carouselindicatorlines.
type CarouselIndicatorLines struct {
	_ [0]func() // equal guard
	gtk.Widget

	*coreglib.Object
	gtk.Orientable
}

var (
	_ gtk.Widgetter     = (*CarouselIndicatorLines)(nil)
	_ coreglib.Objector = (*CarouselIndicatorLines)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*CarouselIndicatorLines, *CarouselIndicatorLinesClass, CarouselIndicatorLinesOverrides](
		GTypeCarouselIndicatorLines,
		initCarouselIndicatorLinesClass,
		wrapCarouselIndicatorLines,
		defaultCarouselIndicatorLinesOverrides,
	)
}

func initCarouselIndicatorLinesClass(gclass unsafe.Pointer, overrides CarouselIndicatorLinesOverrides, classInitFunc func(*CarouselIndicatorLinesClass)) {
	if classInitFunc != nil {
		class := (*CarouselIndicatorLinesClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapCarouselIndicatorLines(obj *coreglib.Object) *CarouselIndicatorLines {
	return &CarouselIndicatorLines{
		Widget: gtk.Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
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
		Object: obj,
		Orientable: gtk.Orientable{
			Object: obj,
		},
	}
}

func marshalCarouselIndicatorLines(p uintptr) (interface{}, error) {
	return wrapCarouselIndicatorLines(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewCarouselIndicatorLines creates a new AdwCarouselIndicatorLines.
//
// The function returns the following values:
//
//   - carouselIndicatorLines: newly created AdwCarouselIndicatorLines.
//
func NewCarouselIndicatorLines() *CarouselIndicatorLines {
	var _cret *C.GtkWidget // in

	_cret = C.adw_carousel_indicator_lines_new()

	var _carouselIndicatorLines *CarouselIndicatorLines // out

	_carouselIndicatorLines = wrapCarouselIndicatorLines(coreglib.Take(unsafe.Pointer(_cret)))

	return _carouselIndicatorLines
}

// Carousel gets the displayed carousel.
//
// The function returns the following values:
//
//   - carousel (optional): displayed carousel.
//
func (self *CarouselIndicatorLines) Carousel() *Carousel {
	var _arg0 *C.AdwCarouselIndicatorLines // out
	var _cret *C.AdwCarousel               // in

	_arg0 = (*C.AdwCarouselIndicatorLines)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_carousel_indicator_lines_get_carousel(_arg0)
	runtime.KeepAlive(self)

	var _carousel *Carousel // out

	if _cret != nil {
		_carousel = wrapCarousel(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _carousel
}

// SetCarousel sets the displayed carousel.
//
// The function takes the following parameters:
//
//   - carousel (optional): carousel.
//
func (self *CarouselIndicatorLines) SetCarousel(carousel *Carousel) {
	var _arg0 *C.AdwCarouselIndicatorLines // out
	var _arg1 *C.AdwCarousel               // out

	_arg0 = (*C.AdwCarouselIndicatorLines)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if carousel != nil {
		_arg1 = (*C.AdwCarousel)(unsafe.Pointer(coreglib.InternObject(carousel).Native()))
	}

	C.adw_carousel_indicator_lines_set_carousel(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(carousel)
}
