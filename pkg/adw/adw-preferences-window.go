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
		{T: externglib.Type(C.adw_preferences_window_get_type()), F: marshalPreferencesWindower},
	})
}

// PreferencesWindow: window to present an application's preferences.
//
// <picture> <source srcset="preferences-window-dark.png"
// media="(prefers-color-scheme: dark)"> <img src="preferences-window.png"
// alt="preferences-window"> </picture>
//
// The AdwPreferencesWindow widget presents an application's preferences
// gathered into pages and groups. The preferences are searchable by the user.
//
//
// CSS nodes
//
// AdwPreferencesWindow has a main CSS node with the name window and the style
// class .preferences.
type PreferencesWindow struct {
	Window
}

var (
	_ gtk.Widgetter       = (*PreferencesWindow)(nil)
	_ externglib.Objector = (*PreferencesWindow)(nil)
)

func wrapPreferencesWindow(obj *externglib.Object) *PreferencesWindow {
	return &PreferencesWindow{
		Window: Window{
			Window: gtk.Window{
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
				Root: gtk.Root{
					NativeSurface: gtk.NativeSurface{
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
				},
				ShortcutManager: gtk.ShortcutManager{
					Object: obj,
				},
				Object: obj,
			},
		},
	}
}

func marshalPreferencesWindower(p uintptr) (interface{}, error) {
	return wrapPreferencesWindow(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewPreferencesWindow creates a new AdwPreferencesWindow.
func NewPreferencesWindow() *PreferencesWindow {
	var _cret *C.GtkWidget // in

	_cret = C.adw_preferences_window_new()

	var _preferencesWindow *PreferencesWindow // out

	_preferencesWindow = wrapPreferencesWindow(externglib.Take(unsafe.Pointer(_cret)))

	return _preferencesWindow
}

// Add adds a preferences page to self.
//
// The function takes the following parameters:
//
//    - page to add.
//
func (self *PreferencesWindow) Add(page *PreferencesPage) {
	var _arg0 *C.AdwPreferencesWindow // out
	var _arg1 *C.AdwPreferencesPage   // out

	_arg0 = (*C.AdwPreferencesWindow)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.AdwPreferencesPage)(unsafe.Pointer(page.Native()))

	C.adw_preferences_window_add(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(page)
}

// AddToast displays toast.
//
// See toastoverlay.AddToast.
//
// The function takes the following parameters:
//
//    - toast: toast.
//
func (self *PreferencesWindow) AddToast(toast *Toast) {
	var _arg0 *C.AdwPreferencesWindow // out
	var _arg1 *C.AdwToast             // out

	_arg0 = (*C.AdwPreferencesWindow)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.AdwToast)(unsafe.Pointer(toast.Native()))
	C.g_object_ref(C.gpointer(toast.Native()))

	C.adw_preferences_window_add_toast(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(toast)
}

// CloseSubpage closes the current subpage.
//
// If there is no presented subpage, this does nothing.
//
// See preferenceswindow.CloseSubpage.
func (self *PreferencesWindow) CloseSubpage() {
	var _arg0 *C.AdwPreferencesWindow // out

	_arg0 = (*C.AdwPreferencesWindow)(unsafe.Pointer(self.Native()))

	C.adw_preferences_window_close_subpage(_arg0)
	runtime.KeepAlive(self)
}

// CanNavigateBack gets whether gestures and shortcuts for closing subpages are
// enabled.
func (self *PreferencesWindow) CanNavigateBack() bool {
	var _arg0 *C.AdwPreferencesWindow // out
	var _cret C.gboolean              // in

	_arg0 = (*C.AdwPreferencesWindow)(unsafe.Pointer(self.Native()))

	_cret = C.adw_preferences_window_get_can_navigate_back(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SearchEnabled gets whether search is enabled for self.
func (self *PreferencesWindow) SearchEnabled() bool {
	var _arg0 *C.AdwPreferencesWindow // out
	var _cret C.gboolean              // in

	_arg0 = (*C.AdwPreferencesWindow)(unsafe.Pointer(self.Native()))

	_cret = C.adw_preferences_window_get_search_enabled(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// VisiblePage gets the currently visible page of self.
func (self *PreferencesWindow) VisiblePage() *PreferencesPage {
	var _arg0 *C.AdwPreferencesWindow // out
	var _cret *C.AdwPreferencesPage   // in

	_arg0 = (*C.AdwPreferencesWindow)(unsafe.Pointer(self.Native()))

	_cret = C.adw_preferences_window_get_visible_page(_arg0)
	runtime.KeepAlive(self)

	var _preferencesPage *PreferencesPage // out

	if _cret != nil {
		_preferencesPage = wrapPreferencesPage(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _preferencesPage
}

// VisiblePageName gets the name of currently visible page of self.
func (self *PreferencesWindow) VisiblePageName() string {
	var _arg0 *C.AdwPreferencesWindow // out
	var _cret *C.char                 // in

	_arg0 = (*C.AdwPreferencesWindow)(unsafe.Pointer(self.Native()))

	_cret = C.adw_preferences_window_get_visible_page_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// PresentSubpage sets subpage as the window's subpage and opens it.
//
// The transition can be cancelled by the user, in which case visible child will
// change back to the previously visible child.
//
// The function takes the following parameters:
//
//    - subpage: subpage.
//
func (self *PreferencesWindow) PresentSubpage(subpage gtk.Widgetter) {
	var _arg0 *C.AdwPreferencesWindow // out
	var _arg1 *C.GtkWidget            // out

	_arg0 = (*C.AdwPreferencesWindow)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(subpage.Native()))

	C.adw_preferences_window_present_subpage(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(subpage)
}

// Remove removes a page from self.
//
// The function takes the following parameters:
//
//    - page to remove.
//
func (self *PreferencesWindow) Remove(page *PreferencesPage) {
	var _arg0 *C.AdwPreferencesWindow // out
	var _arg1 *C.AdwPreferencesPage   // out

	_arg0 = (*C.AdwPreferencesWindow)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.AdwPreferencesPage)(unsafe.Pointer(page.Native()))

	C.adw_preferences_window_remove(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(page)
}

// SetCanNavigateBack sets whether gestures and shortcuts for closing subpages
// are enabled.
//
// The function takes the following parameters:
//
//    - canNavigateBack: new value.
//
func (self *PreferencesWindow) SetCanNavigateBack(canNavigateBack bool) {
	var _arg0 *C.AdwPreferencesWindow // out
	var _arg1 C.gboolean              // out

	_arg0 = (*C.AdwPreferencesWindow)(unsafe.Pointer(self.Native()))
	if canNavigateBack {
		_arg1 = C.TRUE
	}

	C.adw_preferences_window_set_can_navigate_back(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(canNavigateBack)
}

// SetSearchEnabled sets whether search is enabled for self.
//
// The function takes the following parameters:
//
//    - searchEnabled: whether search is enabled.
//
func (self *PreferencesWindow) SetSearchEnabled(searchEnabled bool) {
	var _arg0 *C.AdwPreferencesWindow // out
	var _arg1 C.gboolean              // out

	_arg0 = (*C.AdwPreferencesWindow)(unsafe.Pointer(self.Native()))
	if searchEnabled {
		_arg1 = C.TRUE
	}

	C.adw_preferences_window_set_search_enabled(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(searchEnabled)
}

// SetVisiblePage makes page the visible page of self.
//
// The function takes the following parameters:
//
//    - page of self.
//
func (self *PreferencesWindow) SetVisiblePage(page *PreferencesPage) {
	var _arg0 *C.AdwPreferencesWindow // out
	var _arg1 *C.AdwPreferencesPage   // out

	_arg0 = (*C.AdwPreferencesWindow)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.AdwPreferencesPage)(unsafe.Pointer(page.Native()))

	C.adw_preferences_window_set_visible_page(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(page)
}

// SetVisiblePageName makes the page with the given name visible.
//
// The function takes the following parameters:
//
//    - name of the page to make visible.
//
func (self *PreferencesWindow) SetVisiblePageName(name string) {
	var _arg0 *C.AdwPreferencesWindow // out
	var _arg1 *C.char                 // out

	_arg0 = (*C.AdwPreferencesWindow)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_preferences_window_set_visible_page_name(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(name)
}
