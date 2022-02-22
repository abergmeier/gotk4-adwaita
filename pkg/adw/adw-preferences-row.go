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
		{T: externglib.Type(C.adw_preferences_row_get_type()), F: marshalPreferencesRower},
	})
}

// PreferencesRow: gtk.ListBoxRow used to present preferences.
//
// The AdwPreferencesRow widget has a title that preferenceswindow will use to
// let the user look for a preference. It doesn't present the title in any way
// and lets you present the preference as you please.
//
// actionrow and its derivatives are convenient to use as preference rows as
// they take care of presenting the preference's title while letting you compose
// the inputs of the preference around it.
type PreferencesRow struct {
	gtk.ListBoxRow
}

var (
	_ gtk.Widgetter       = (*PreferencesRow)(nil)
	_ externglib.Objector = (*PreferencesRow)(nil)
)

func wrapPreferencesRow(obj *externglib.Object) *PreferencesRow {
	return &PreferencesRow{
		ListBoxRow: gtk.ListBoxRow{
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
			Actionable: gtk.Actionable{
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
			},
			Object: obj,
		},
	}
}

func marshalPreferencesRower(p uintptr) (interface{}, error) {
	return wrapPreferencesRow(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewPreferencesRow creates a new AdwPreferencesRow.
func NewPreferencesRow() *PreferencesRow {
	var _cret *C.GtkWidget // in

	_cret = C.adw_preferences_row_new()

	var _preferencesRow *PreferencesRow // out

	_preferencesRow = wrapPreferencesRow(externglib.Take(unsafe.Pointer(_cret)))

	return _preferencesRow
}

// Title gets the title of the preference represented by self.
func (self *PreferencesRow) Title() string {
	var _arg0 *C.AdwPreferencesRow // out
	var _cret *C.char              // in

	_arg0 = (*C.AdwPreferencesRow)(unsafe.Pointer(self.Native()))

	_cret = C.adw_preferences_row_get_title(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// UseUnderline gets whether an embedded underline in the title indicates a
// mnemonic.
func (self *PreferencesRow) UseUnderline() bool {
	var _arg0 *C.AdwPreferencesRow // out
	var _cret C.gboolean           // in

	_arg0 = (*C.AdwPreferencesRow)(unsafe.Pointer(self.Native()))

	_cret = C.adw_preferences_row_get_use_underline(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetTitle sets the title of the preference represented by self.
//
// The function takes the following parameters:
//
//    - title: title.
//
func (self *PreferencesRow) SetTitle(title string) {
	var _arg0 *C.AdwPreferencesRow // out
	var _arg1 *C.char              // out

	_arg0 = (*C.AdwPreferencesRow)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_preferences_row_set_title(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(title)
}

// SetUseUnderline sets whether an embedded underline in the title indicates a
// mnemonic.
//
// The function takes the following parameters:
//
//    - useUnderline: TRUE if underlines in the text indicate mnemonics.
//
func (self *PreferencesRow) SetUseUnderline(useUnderline bool) {
	var _arg0 *C.AdwPreferencesRow // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.AdwPreferencesRow)(unsafe.Pointer(self.Native()))
	if useUnderline {
		_arg1 = C.TRUE
	}

	C.adw_preferences_row_set_use_underline(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(useUnderline)
}
