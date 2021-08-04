// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #cgo pkg-config: libadwaita-1
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <adwaita.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.adw_tab_bar_get_type()), F: marshalTabBarrer},
	})
}

// TabBar: tab bar for adw.TabView.
//
// The AdwTabBar widget is a tab bar that can be used with conjunction with
// AdwTabView.
//
// AdwTabBar can autohide and can optionally contain action widgets on both
// sides of the tabs.
//
// When there's not enough space to show all the tabs, AdwTabBar will scroll
// them. Pinned tabs always stay visible and aren't a part of the scrollable
// area.
//
//
// CSS nodes
//
// AdwTabBar has a single CSS node with name tabbar.
type TabBar struct {
	gtk.Widget
}

func wrapTabBar(obj *externglib.Object) *TabBar {
	return &TabBar{
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
	}
}

func marshalTabBarrer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTabBar(obj), nil
}

// NewTabBar creates a new AdwTabBar.
func NewTabBar() *TabBar {
	var _cret *C.AdwTabBar // in

	_cret = C.adw_tab_bar_new()

	var _tabBar *TabBar // out

	_tabBar = wrapTabBar(externglib.Take(unsafe.Pointer(_cret)))

	return _tabBar
}

// Autohide gets whether the tabs automatically hide.
func (self *TabBar) Autohide() bool {
	var _arg0 *C.AdwTabBar // out
	var _cret C.gboolean   // in

	_arg0 = (*C.AdwTabBar)(unsafe.Pointer(self.Native()))

	_cret = C.adw_tab_bar_get_autohide(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// EndActionWidget gets the widget shown after the tabs.
func (self *TabBar) EndActionWidget() gtk.Widgetter {
	var _arg0 *C.AdwTabBar // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.AdwTabBar)(unsafe.Pointer(self.Native()))

	_cret = C.adw_tab_bar_get_end_action_widget(_arg0)

	var _widget gtk.Widgetter // out

	if _cret != nil {
		_widget = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(gtk.Widgetter)
	}

	return _widget
}

// ExpandTabs gets whether tabs expand to full width.
func (self *TabBar) ExpandTabs() bool {
	var _arg0 *C.AdwTabBar // out
	var _cret C.gboolean   // in

	_arg0 = (*C.AdwTabBar)(unsafe.Pointer(self.Native()))

	_cret = C.adw_tab_bar_get_expand_tabs(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Inverted gets whether tabs use inverted layout.
func (self *TabBar) Inverted() bool {
	var _arg0 *C.AdwTabBar // out
	var _cret C.gboolean   // in

	_arg0 = (*C.AdwTabBar)(unsafe.Pointer(self.Native()))

	_cret = C.adw_tab_bar_get_inverted(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsOverflowing gets whether self is overflowing.
func (self *TabBar) IsOverflowing() bool {
	var _arg0 *C.AdwTabBar // out
	var _cret C.gboolean   // in

	_arg0 = (*C.AdwTabBar)(unsafe.Pointer(self.Native()))

	_cret = C.adw_tab_bar_get_is_overflowing(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// StartActionWidget gets the widget shown before the tabs.
func (self *TabBar) StartActionWidget() gtk.Widgetter {
	var _arg0 *C.AdwTabBar // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.AdwTabBar)(unsafe.Pointer(self.Native()))

	_cret = C.adw_tab_bar_get_start_action_widget(_arg0)

	var _widget gtk.Widgetter // out

	if _cret != nil {
		_widget = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(gtk.Widgetter)
	}

	return _widget
}

// TabsRevealed gets whether the tabs are currently revealed.
func (self *TabBar) TabsRevealed() bool {
	var _arg0 *C.AdwTabBar // out
	var _cret C.gboolean   // in

	_arg0 = (*C.AdwTabBar)(unsafe.Pointer(self.Native()))

	_cret = C.adw_tab_bar_get_tabs_revealed(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// View gets the tab view self controls.
func (self *TabBar) View() *TabView {
	var _arg0 *C.AdwTabBar  // out
	var _cret *C.AdwTabView // in

	_arg0 = (*C.AdwTabBar)(unsafe.Pointer(self.Native()))

	_cret = C.adw_tab_bar_get_view(_arg0)

	var _tabView *TabView // out

	if _cret != nil {
		_tabView = wrapTabView(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _tabView
}

// SetAutohide sets whether the tabs automatically hide.
func (self *TabBar) SetAutohide(autohide bool) {
	var _arg0 *C.AdwTabBar // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.AdwTabBar)(unsafe.Pointer(self.Native()))
	if autohide {
		_arg1 = C.TRUE
	}

	C.adw_tab_bar_set_autohide(_arg0, _arg1)
}

// SetEndActionWidget sets the widget to show after the tabs.
func (self *TabBar) SetEndActionWidget(widget gtk.Widgetter) {
	var _arg0 *C.AdwTabBar // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.AdwTabBar)(unsafe.Pointer(self.Native()))
	if widget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	}

	C.adw_tab_bar_set_end_action_widget(_arg0, _arg1)
}

// SetExpandTabs sets whether tabs expand to full width.
func (self *TabBar) SetExpandTabs(expandTabs bool) {
	var _arg0 *C.AdwTabBar // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.AdwTabBar)(unsafe.Pointer(self.Native()))
	if expandTabs {
		_arg1 = C.TRUE
	}

	C.adw_tab_bar_set_expand_tabs(_arg0, _arg1)
}

// SetInverted sets whether tabs tabs use inverted layout.
func (self *TabBar) SetInverted(inverted bool) {
	var _arg0 *C.AdwTabBar // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.AdwTabBar)(unsafe.Pointer(self.Native()))
	if inverted {
		_arg1 = C.TRUE
	}

	C.adw_tab_bar_set_inverted(_arg0, _arg1)
}

// SetStartActionWidget sets the widget to show before the tabs.
func (self *TabBar) SetStartActionWidget(widget gtk.Widgetter) {
	var _arg0 *C.AdwTabBar // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.AdwTabBar)(unsafe.Pointer(self.Native()))
	if widget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	}

	C.adw_tab_bar_set_start_action_widget(_arg0, _arg1)
}

// SetView sets the tab view self controls.
func (self *TabBar) SetView(view *TabView) {
	var _arg0 *C.AdwTabBar  // out
	var _arg1 *C.AdwTabView // out

	_arg0 = (*C.AdwTabBar)(unsafe.Pointer(self.Native()))
	if view != nil {
		_arg1 = (*C.AdwTabView)(unsafe.Pointer(view.Native()))
	}

	C.adw_tab_bar_set_view(_arg0, _arg1)
}

// SetupExtraDropTarget sets the supported types for this drop target.
//
// Sets up an extra drop target on tabs.
//
// This allows to drag arbitrary content onto tabs, for example URLs in a web
// browser.
//
// If a tab is hovered for a certain period of time while dragging the content,
// it will be automatically selected.
//
// The adw.TabBar::extra-drag-drop signal can be used to handle the drop.
func (self *TabBar) SetupExtraDropTarget(actions gdk.DragAction, types []externglib.Type) {
	var _arg0 *C.AdwTabBar    // out
	var _arg1 C.GdkDragAction // out
	var _arg2 *C.GType        // out
	var _arg3 C.gsize

	_arg0 = (*C.AdwTabBar)(unsafe.Pointer(self.Native()))
	_arg1 = C.GdkDragAction(actions)
	_arg3 = (C.gsize)(len(types))
	_arg2 = (*C.GType)(C.malloc(C.ulong(len(types)) * C.ulong(C.sizeof_GType)))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice((*C.GType)(_arg2), len(types))
		for i := range types {
			out[i] = C.GType(types[i])
		}
	}

	C.adw_tab_bar_setup_extra_drop_target(_arg0, _arg1, _arg2, _arg3)
}
