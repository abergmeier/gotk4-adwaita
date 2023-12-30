// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #include <stdlib.h>
// #include <adwaita.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeColorScheme  = coreglib.Type(C.adw_color_scheme_get_type())
	GTypeStyleManager = coreglib.Type(C.adw_style_manager_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeColorScheme, F: marshalColorScheme},
		coreglib.TypeMarshaler{T: GTypeStyleManager, F: marshalStyleManager},
	})
}

// ColorScheme: application color schemes for stylemanager:color-scheme.
type ColorScheme C.gint

const (
	// ColorSchemeDefault: inherit the parent color-scheme. When set on the
	// AdwStyleManager returned by stylemanager.GetDefault(), it's equivalent to
	// ADW_COLOR_SCHEME_PREFER_LIGHT.
	ColorSchemeDefault ColorScheme = iota
	// ColorSchemeForceLight always use light appearance.
	ColorSchemeForceLight
	// ColorSchemePreferLight: use light appearance unless the system prefers
	// dark colors.
	ColorSchemePreferLight
	// ColorSchemePreferDark: use dark appearance unless the system prefers
	// prefers light colors.
	ColorSchemePreferDark
	// ColorSchemeForceDark always use dark appearance.
	ColorSchemeForceDark
)

func marshalColorScheme(p uintptr) (interface{}, error) {
	return ColorScheme(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for ColorScheme.
func (c ColorScheme) String() string {
	switch c {
	case ColorSchemeDefault:
		return "Default"
	case ColorSchemeForceLight:
		return "ForceLight"
	case ColorSchemePreferLight:
		return "PreferLight"
	case ColorSchemePreferDark:
		return "PreferDark"
	case ColorSchemeForceDark:
		return "ForceDark"
	default:
		return fmt.Sprintf("ColorScheme(%d)", c)
	}
}

// StyleManagerOverrides contains methods that are overridable.
type StyleManagerOverrides struct {
}

func defaultStyleManagerOverrides(v *StyleManager) StyleManagerOverrides {
	return StyleManagerOverrides{}
}

// StyleManager class for managing application-wide styling.
//
// AdwStyleManager provides a way to query and influence the application styles,
// such as whether to use dark or high contrast appearance.
//
// It allows to set the color scheme via the stylemanager:color-scheme property,
// and to query the current appearance, as well as whether a system-wide color
// scheme preference exists.
type StyleManager struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*StyleManager)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*StyleManager, *StyleManagerClass, StyleManagerOverrides](
		GTypeStyleManager,
		initStyleManagerClass,
		wrapStyleManager,
		defaultStyleManagerOverrides,
	)
}

func initStyleManagerClass(gclass unsafe.Pointer, overrides StyleManagerOverrides, classInitFunc func(*StyleManagerClass)) {
	if classInitFunc != nil {
		class := (*StyleManagerClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapStyleManager(obj *coreglib.Object) *StyleManager {
	return &StyleManager{
		Object: obj,
	}
}

func marshalStyleManager(p uintptr) (interface{}, error) {
	return wrapStyleManager(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ColorScheme gets the requested application color scheme.
//
// The function returns the following values:
//
//   - colorScheme: color scheme.
//
func (self *StyleManager) ColorScheme() ColorScheme {
	var _arg0 *C.AdwStyleManager // out
	var _cret C.AdwColorScheme   // in

	_arg0 = (*C.AdwStyleManager)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_style_manager_get_color_scheme(_arg0)
	runtime.KeepAlive(self)

	var _colorScheme ColorScheme // out

	_colorScheme = ColorScheme(_cret)

	return _colorScheme
}

// Dark gets whether the application is using dark appearance.
//
// The function returns the following values:
//
//   - ok: whether the application is using dark appearance.
//
func (self *StyleManager) Dark() bool {
	var _arg0 *C.AdwStyleManager // out
	var _cret C.gboolean         // in

	_arg0 = (*C.AdwStyleManager)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_style_manager_get_dark(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Display gets the display the style manager is associated with.
//
// The display will be NULL for the style manager returned by
// stylemanager.GetDefault().
//
// The function returns the following values:
//
//   - display: (nullable): the display.
//
func (self *StyleManager) Display() *gdk.Display {
	var _arg0 *C.AdwStyleManager // out
	var _cret *C.GdkDisplay      // in

	_arg0 = (*C.AdwStyleManager)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_style_manager_get_display(_arg0)
	runtime.KeepAlive(self)

	var _display *gdk.Display // out

	{
		obj := coreglib.Take(unsafe.Pointer(_cret))
		_display = &gdk.Display{
			Object: obj,
		}
	}

	return _display
}

// HighContrast gets whether the application is using high contrast appearance.
//
// The function returns the following values:
//
//   - ok: whether the application is using high contrast appearance.
//
func (self *StyleManager) HighContrast() bool {
	var _arg0 *C.AdwStyleManager // out
	var _cret C.gboolean         // in

	_arg0 = (*C.AdwStyleManager)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_style_manager_get_high_contrast(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SystemSupportsColorSchemes gets whether the system supports color schemes.
//
// The function returns the following values:
//
//   - ok: whether the system supports color schemes.
//
func (self *StyleManager) SystemSupportsColorSchemes() bool {
	var _arg0 *C.AdwStyleManager // out
	var _cret C.gboolean         // in

	_arg0 = (*C.AdwStyleManager)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_style_manager_get_system_supports_color_schemes(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetColorScheme sets the requested application color scheme.
//
// The effective appearance will be decided based on the application color
// scheme and the system preferred color scheme. The stylemanager:dark property
// can be used to query the current effective appearance.
//
// The function takes the following parameters:
//
//   - colorScheme: color scheme.
//
func (self *StyleManager) SetColorScheme(colorScheme ColorScheme) {
	var _arg0 *C.AdwStyleManager // out
	var _arg1 C.AdwColorScheme   // out

	_arg0 = (*C.AdwStyleManager)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.AdwColorScheme(colorScheme)

	C.adw_style_manager_set_color_scheme(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(colorScheme)
}

// StyleManagerGetDefault gets the default AdwStyleManager instance.
//
// It manages all gdk.Display instances unless the style manager for that
// display has an override.
//
// See stylemanager.GetForDisplay().
//
// The function returns the following values:
//
//   - styleManager: default style manager.
//
func StyleManagerGetDefault() *StyleManager {
	var _cret *C.AdwStyleManager // in

	_cret = C.adw_style_manager_get_default()

	var _styleManager *StyleManager // out

	_styleManager = wrapStyleManager(coreglib.Take(unsafe.Pointer(_cret)))

	return _styleManager
}

// StyleManagerGetForDisplay gets the AdwStyleManager instance managing display.
//
// It can be used to override styles for that specific display instead of the
// whole application.
//
// Most applications should use stylemanager.GetDefault() instead.
//
// The function takes the following parameters:
//
//   - display: GdkDisplay.
//
// The function returns the following values:
//
//   - styleManager: style manager for display.
//
func StyleManagerGetForDisplay(display *gdk.Display) *StyleManager {
	var _arg1 *C.GdkDisplay      // out
	var _cret *C.AdwStyleManager // in

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(coreglib.InternObject(display).Native()))

	_cret = C.adw_style_manager_get_for_display(_arg1)
	runtime.KeepAlive(display)

	var _styleManager *StyleManager // out

	_styleManager = wrapStyleManager(coreglib.Take(unsafe.Pointer(_cret)))

	return _styleManager
}
