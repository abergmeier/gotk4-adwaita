// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"fmt"
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <adwaita.h>
// #include <glib-object.h>
// extern void _gotk4_adw1_Animation_ConnectDone(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeAnimationState = coreglib.Type(C.adw_animation_state_get_type())
	GTypeAnimation      = coreglib.Type(C.adw_animation_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeAnimationState, F: marshalAnimationState},
		coreglib.TypeMarshaler{T: GTypeAnimation, F: marshalAnimation},
	})
}

// AnimationState describes the possible states of an animation.
//
// The state can be controlled with animation.Play, animation.Pause,
// animation.Resume, animation.Reset and animation.Skip.
type AnimationState C.gint

const (
	// AnimationIdle: animation hasn't started yet.
	AnimationIdle AnimationState = iota
	// AnimationPaused: animation has been paused.
	AnimationPaused
	// AnimationPlaying: animation is currently playing.
	AnimationPlaying
	// AnimationFinished: animation has finished.
	AnimationFinished
)

func marshalAnimationState(p uintptr) (interface{}, error) {
	return AnimationState(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for AnimationState.
func (a AnimationState) String() string {
	switch a {
	case AnimationIdle:
		return "Idle"
	case AnimationPaused:
		return "Paused"
	case AnimationPlaying:
		return "Playing"
	case AnimationFinished:
		return "Finished"
	default:
		return fmt.Sprintf("AnimationState(%d)", a)
	}
}

// Animation: base class for animations.
//
// AdwAnimation represents an animation on a widget. It has a target that
// provides a value to animate, and a state indicating whether the animation
// hasn't been started yet, is playing, paused or finished.
//
// Currently there are two concrete animation types: timedanimation and
// springanimation.
//
// AdwAnimation will automatically skip the animation if animation:widget is
// unmapped, or if gtk.Settings:gtk-enable-animations is FALSE.
//
// The animation::done signal can be used to perform an action after
// the animation ends, for example hiding a widget after animating its
// gtk.Widget:opacity to 0.
//
// AdwAnimation will be kept alive while the animation is playing. As such,
// it's safe to create an animation, start it and immediately unref it:
// A fire-and-forget animation:
//
//    static void
//    animation_cb (double    value,
//                  MyObject *self)
//    {
//      // Do something with value
//    }
//
//    static void
//    my_object_animate (MyObject *self)
//    {
//      AdwAnimationTarget *target =
//        adw_callback_animation_target_new ((AdwAnimationTargetFunc) animation_cb,
//                                           self, NULL);
//      g_autoptr (AdwAnimation) animation =
//        adw_timed_animation_new (widget, 0, 1, 250, target);
//
//      adw_animation_play (animation);
//    }
//
// If there's a chance the previous animation for the same target hasn't yet
// finished, the previous animation should be stopped first, or the existing
// AdwAnimation object can be reused.
type Animation struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Animation)(nil)
)

// Animationer describes types inherited from class Animation.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type Animationer interface {
	coreglib.Objector
	baseAnimation() *Animation
}

var _ Animationer = (*Animation)(nil)

func wrapAnimation(obj *coreglib.Object) *Animation {
	return &Animation{
		Object: obj,
	}
}

func marshalAnimation(p uintptr) (interface{}, error) {
	return wrapAnimation(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (self *Animation) baseAnimation() *Animation {
	return self
}

// BaseAnimation returns the underlying base object.
func BaseAnimation(obj Animationer) *Animation {
	return obj.baseAnimation()
}

// ConnectDone: this signal is emitted when the animation has been completed,
// either on its own or via calling animation.Skip.
func (self *Animation) ConnectDone(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "done", false, unsafe.Pointer(C._gotk4_adw1_Animation_ConnectDone), f)
}

// State gets the current value of self.
//
// The state indicates whether self is currently playing, paused, finished or
// hasn't been started yet.
//
// The function returns the following values:
//
//   - animationState: animation value.
//
func (self *Animation) State() AnimationState {
	var _arg0 *C.AdwAnimation     // out
	var _cret C.AdwAnimationState // in

	_arg0 = (*C.AdwAnimation)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_animation_get_state(_arg0)
	runtime.KeepAlive(self)

	var _animationState AnimationState // out

	_animationState = AnimationState(_cret)

	return _animationState
}

// Target gets the target self animates.
//
// The function returns the following values:
//
//   - animationTarget: animation target.
//
func (self *Animation) Target() AnimationTargetter {
	var _arg0 *C.AdwAnimation       // out
	var _cret *C.AdwAnimationTarget // in

	_arg0 = (*C.AdwAnimation)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_animation_get_target(_arg0)
	runtime.KeepAlive(self)

	var _animationTarget AnimationTargetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type adw.AnimationTargetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(AnimationTargetter)
			return ok
		})
		rv, ok := casted.(AnimationTargetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching adw.AnimationTargetter")
		}
		_animationTarget = rv
	}

	return _animationTarget
}

// Value gets the current value of self.
//
// The function returns the following values:
//
//   - gdouble: current value.
//
func (self *Animation) Value() float64 {
	var _arg0 *C.AdwAnimation // out
	var _cret C.double        // in

	_arg0 = (*C.AdwAnimation)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_animation_get_value(_arg0)
	runtime.KeepAlive(self)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Widget gets the widget self was created for.
//
// The function returns the following values:
//
//   - widget: animation widget.
//
func (self *Animation) Widget() gtk.Widgetter {
	var _arg0 *C.AdwAnimation // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.AdwAnimation)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_animation_get_widget(_arg0)
	runtime.KeepAlive(self)

	var _widget gtk.Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(gtk.Widgetter)
			return ok
		})
		rv, ok := casted.(gtk.Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// Pause pauses a playing animation for self.
//
// Does nothing if the current state of self isn't ADW_ANIMATION_PLAYING.
//
// Sets animation:state to ADW_ANIMATION_PAUSED.
func (self *Animation) Pause() {
	var _arg0 *C.AdwAnimation // out

	_arg0 = (*C.AdwAnimation)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	C.adw_animation_pause(_arg0)
	runtime.KeepAlive(self)
}

// Play starts the animation for self.
//
// If the animation is playing, paused or has been completed, restarts it from
// the beginning. This allows to easily play an animation regardless of whether
// it's already playing or not.
//
// Sets animation:state to ADW_ANIMATION_PLAYING.
//
// The animation will be automatically skipped if animation:widget is unmapped,
// or if gtk.Settings:gtk-enable-animations is FALSE.
//
// As such, it's not guaranteed that the animation will actually run.
// For example, when using glib.IdleAdd() and starting an animation immediately
// afterwards, it's entirely possible that the idle callback will run after the
// animation has already finished, and not while it's playing.
func (self *Animation) Play() {
	var _arg0 *C.AdwAnimation // out

	_arg0 = (*C.AdwAnimation)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	C.adw_animation_play(_arg0)
	runtime.KeepAlive(self)
}

// Reset resets the animation for self.
//
// Sets animation:state to ADW_ANIMATION_IDLE.
func (self *Animation) Reset() {
	var _arg0 *C.AdwAnimation // out

	_arg0 = (*C.AdwAnimation)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	C.adw_animation_reset(_arg0)
	runtime.KeepAlive(self)
}

// Resume resumes a paused animation for self.
//
// This function must only be used if the animation has been paused with
// animation.Pause.
//
// Sets animation:state to ADW_ANIMATION_PLAYING.
func (self *Animation) Resume() {
	var _arg0 *C.AdwAnimation // out

	_arg0 = (*C.AdwAnimation)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	C.adw_animation_resume(_arg0)
	runtime.KeepAlive(self)
}

// Skip skips the animation for self.
//
// If the animation hasn't been started yet, is playing, or is paused, instantly
// skips the animation to the end and causes animation::done to be emitted.
//
// Sets animation:state to ADW_ANIMATION_FINISHED.
func (self *Animation) Skip() {
	var _arg0 *C.AdwAnimation // out

	_arg0 = (*C.AdwAnimation)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	C.adw_animation_skip(_arg0)
	runtime.KeepAlive(self)
}
