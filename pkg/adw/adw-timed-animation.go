// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #cgo pkg-config: libadwaita-1
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <adwaita.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.adw_timed_animation_get_type()), F: marshalTimedAnimationer},
	})
}

// TimedAnimation: time-based animation.
//
// AdwTimedAnimation implements a simple animation interpolating the given value
// from timedanimation:value-from to timedanimation:value-to over
// timedanimation:duration milliseconds using the curve described by
// timedanimation:easing.
//
// If timedanimation:reverse is set to TRUE, AdwTimedAnimation will instead
// animate from timedanimation:value-to to timedanimation:value-from, and the
// easing curve will be inverted.
//
// The animation can repeat a certain amount of times, or endlessly, depending
// on the timedanimation:repeat-count value. If timedanimation:alternate is set
// to TRUE, it will also change the direction every other iteration.
type TimedAnimation struct {
	Animation
}

var (
	_ Animationer = (*TimedAnimation)(nil)
)

func wrapTimedAnimation(obj *externglib.Object) *TimedAnimation {
	return &TimedAnimation{
		Animation: Animation{
			Object: obj,
		},
	}
}

func marshalTimedAnimationer(p uintptr) (interface{}, error) {
	return wrapTimedAnimation(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewTimedAnimation creates a new AdwTimedAnimation on widget to animate target
// from from to to.
//
// The function takes the following parameters:
//
//    - widget to create animation on.
//    - from: value to animate from.
//    - to: value to animate to.
//    - duration for the animation.
//    - target value to animate.
//
func NewTimedAnimation(widget gtk.Widgetter, from, to float64, duration uint, target AnimationTargetter) *TimedAnimation {
	var _arg1 *C.GtkWidget          // out
	var _arg2 C.double              // out
	var _arg3 C.double              // out
	var _arg4 C.guint               // out
	var _arg5 *C.AdwAnimationTarget // out
	var _cret *C.AdwAnimation       // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = C.double(from)
	_arg3 = C.double(to)
	_arg4 = C.guint(duration)
	_arg5 = (*C.AdwAnimationTarget)(unsafe.Pointer(target.Native()))
	C.g_object_ref(C.gpointer(target.Native()))

	_cret = C.adw_timed_animation_new(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(from)
	runtime.KeepAlive(to)
	runtime.KeepAlive(duration)
	runtime.KeepAlive(target)

	var _timedAnimation *TimedAnimation // out

	_timedAnimation = wrapTimedAnimation(externglib.Take(unsafe.Pointer(_cret)))

	return _timedAnimation
}

// Alternate gets whether self changes direction on every iteration.
func (self *TimedAnimation) Alternate() bool {
	var _arg0 *C.AdwTimedAnimation // out
	var _cret C.gboolean           // in

	_arg0 = (*C.AdwTimedAnimation)(unsafe.Pointer(self.Native()))

	_cret = C.adw_timed_animation_get_alternate(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Duration gets the duration of self in milliseconds.
func (self *TimedAnimation) Duration() uint {
	var _arg0 *C.AdwTimedAnimation // out
	var _cret C.guint              // in

	_arg0 = (*C.AdwTimedAnimation)(unsafe.Pointer(self.Native()))

	_cret = C.adw_timed_animation_get_duration(_arg0)
	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Easing gets the easing function self uses.
func (self *TimedAnimation) Easing() Easing {
	var _arg0 *C.AdwTimedAnimation // out
	var _cret C.AdwEasing          // in

	_arg0 = (*C.AdwTimedAnimation)(unsafe.Pointer(self.Native()))

	_cret = C.adw_timed_animation_get_easing(_arg0)
	runtime.KeepAlive(self)

	var _easing Easing // out

	_easing = Easing(_cret)

	return _easing
}

// RepeatCount gets the number of times self will play.
func (self *TimedAnimation) RepeatCount() uint {
	var _arg0 *C.AdwTimedAnimation // out
	var _cret C.guint              // in

	_arg0 = (*C.AdwTimedAnimation)(unsafe.Pointer(self.Native()))

	_cret = C.adw_timed_animation_get_repeat_count(_arg0)
	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Reverse gets whether self plays backwards.
func (self *TimedAnimation) Reverse() bool {
	var _arg0 *C.AdwTimedAnimation // out
	var _cret C.gboolean           // in

	_arg0 = (*C.AdwTimedAnimation)(unsafe.Pointer(self.Native()))

	_cret = C.adw_timed_animation_get_reverse(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ValueFrom gets the value self will animate from.
func (self *TimedAnimation) ValueFrom() float64 {
	var _arg0 *C.AdwTimedAnimation // out
	var _cret C.double             // in

	_arg0 = (*C.AdwTimedAnimation)(unsafe.Pointer(self.Native()))

	_cret = C.adw_timed_animation_get_value_from(_arg0)
	runtime.KeepAlive(self)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// ValueTo gets the value self will animate to.
func (self *TimedAnimation) ValueTo() float64 {
	var _arg0 *C.AdwTimedAnimation // out
	var _cret C.double             // in

	_arg0 = (*C.AdwTimedAnimation)(unsafe.Pointer(self.Native()))

	_cret = C.adw_timed_animation_get_value_to(_arg0)
	runtime.KeepAlive(self)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// SetAlternate sets whether self changes direction on every iteration.
//
// The function takes the following parameters:
//
//    - alternate: whether self alternates.
//
func (self *TimedAnimation) SetAlternate(alternate bool) {
	var _arg0 *C.AdwTimedAnimation // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.AdwTimedAnimation)(unsafe.Pointer(self.Native()))
	if alternate {
		_arg1 = C.TRUE
	}

	C.adw_timed_animation_set_alternate(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(alternate)
}

// SetDuration sets the duration of self in milliseconds.
//
// If the animation repeats more than once, sets the duration of one iteration.
//
// The function takes the following parameters:
//
//    - duration to use.
//
func (self *TimedAnimation) SetDuration(duration uint) {
	var _arg0 *C.AdwTimedAnimation // out
	var _arg1 C.guint              // out

	_arg0 = (*C.AdwTimedAnimation)(unsafe.Pointer(self.Native()))
	_arg1 = C.guint(duration)

	C.adw_timed_animation_set_duration(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(duration)
}

// SetEasing sets the easing function self will use.
//
// See easing for the description of specific easing functions.
//
// The function takes the following parameters:
//
//    - easing function to use.
//
func (self *TimedAnimation) SetEasing(easing Easing) {
	var _arg0 *C.AdwTimedAnimation // out
	var _arg1 C.AdwEasing          // out

	_arg0 = (*C.AdwTimedAnimation)(unsafe.Pointer(self.Native()))
	_arg1 = C.AdwEasing(easing)

	C.adw_timed_animation_set_easing(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(easing)
}

// SetRepeatCount sets the number of times self will play.
//
// If set to 0, self will repeat endlessly.
//
// The function takes the following parameters:
//
//    - repeatCount: number of times self will play.
//
func (self *TimedAnimation) SetRepeatCount(repeatCount uint) {
	var _arg0 *C.AdwTimedAnimation // out
	var _arg1 C.guint              // out

	_arg0 = (*C.AdwTimedAnimation)(unsafe.Pointer(self.Native()))
	_arg1 = C.guint(repeatCount)

	C.adw_timed_animation_set_repeat_count(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(repeatCount)
}

// SetReverse sets whether self plays backwards.
//
// The function takes the following parameters:
//
//    - reverse: whether self plays backwards.
//
func (self *TimedAnimation) SetReverse(reverse bool) {
	var _arg0 *C.AdwTimedAnimation // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.AdwTimedAnimation)(unsafe.Pointer(self.Native()))
	if reverse {
		_arg1 = C.TRUE
	}

	C.adw_timed_animation_set_reverse(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(reverse)
}

// SetValueFrom sets the value self will animate from.
//
// The function takes the following parameters:
//
//    - value to animate from.
//
func (self *TimedAnimation) SetValueFrom(value float64) {
	var _arg0 *C.AdwTimedAnimation // out
	var _arg1 C.double             // out

	_arg0 = (*C.AdwTimedAnimation)(unsafe.Pointer(self.Native()))
	_arg1 = C.double(value)

	C.adw_timed_animation_set_value_from(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetValueTo sets the value self will animate to.
//
// The function takes the following parameters:
//
//    - value to animate to.
//
func (self *TimedAnimation) SetValueTo(value float64) {
	var _arg0 *C.AdwTimedAnimation // out
	var _arg1 C.double             // out

	_arg0 = (*C.AdwTimedAnimation)(unsafe.Pointer(self.Native()))
	_arg1 = C.double(value)

	C.adw_timed_animation_set_value_to(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}
