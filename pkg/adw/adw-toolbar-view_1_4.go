// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"fmt"
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
	GTypeToolbarStyle = coreglib.Type(C.adw_toolbar_style_get_type())
	GTypeToolbarView  = coreglib.Type(C.adw_toolbar_view_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeToolbarStyle, F: marshalToolbarStyle},
		coreglib.TypeMarshaler{T: GTypeToolbarView, F: marshalToolbarView},
	})
}

// ToolbarStyle describes the possible top or bottom bar styles in an
// toolbarview widget.
//
// ADW_TOOLBAR_FLAT is suitable for simple content, such as statuspage or
// preferencespage, where the background at the top and bottom parts of the page
// is uniform. Additionally, windows with sidebars should always use this style.
//
// <picture style="min-width: 33%; display: inline-block;"> <source
// srcset="toolbar-view-flat-1-dark.png" media="(prefers-color-scheme:
// dark)"> <img src="toolbar-view-flat-1.png" alt="toolbar-view-flat-1">
// </picture> <picture style="min-width: 33%; display: inline-block;"> <source
// srcset="toolbar-view-flat-2-dark.png" media="(prefers-color-scheme: dark)">
// <img src="toolbar-view-flat-2.png" alt="toolbar-view-flat-2"> </picture>
//
// ADW_TOOLBAR_RAISED style is suitable for content such as utility panes
// (https://developer.gnome.org/hig/patterns/containers/utility-panes.html),
// where some elements are directly adjacent to the top/bottom bars, or tabview,
// where each page can have a different background.
//
// ADW_TOOLBAR_RAISED_BORDER style is similar to ADW_TOOLBAR_RAISED, but with
// the shadow replaced with a more subtle border. It's intended to be used in
// applications like image viewers, where a shadow over the content might be
// undesired.
//
// <picture style="min-width: 33%; display: inline-block;"> <source
// srcset="toolbar-view-raised-dark.png" media="(prefers-color-scheme:
// dark)"> <img src="toolbar-view-raised.png" alt="toolbar-view-raised">
// </picture> <picture style="min-width: 33%; display: inline-block;"> <source
// srcset="toolbar-view-raised-border-dark.png" media="(prefers-color-scheme:
// dark)"> <img src="toolbar-view-raised-border.png"
// alt="toolbar-view-raised-border"> </picture>
//
// See toolbarview:top-bar-style and toolbarview:bottom-bar-style.
//
// New values may be added to this enumeration over time.
type ToolbarStyle C.gint

const (
	// ToolbarFlat: no background, shadow only for scrolled content.
	ToolbarFlat ToolbarStyle = iota
	// ToolbarRaised: opaque background with a persistent shadow.
	ToolbarRaised
	// ToolbarRaisedBorder: opaque background with a persistent border.
	ToolbarRaisedBorder
)

func marshalToolbarStyle(p uintptr) (interface{}, error) {
	return ToolbarStyle(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for ToolbarStyle.
func (t ToolbarStyle) String() string {
	switch t {
	case ToolbarFlat:
		return "Flat"
	case ToolbarRaised:
		return "Raised"
	case ToolbarRaisedBorder:
		return "RaisedBorder"
	default:
		return fmt.Sprintf("ToolbarStyle(%d)", t)
	}
}

// ToolbarViewOverrides contains methods that are overridable.
type ToolbarViewOverrides struct {
}

func defaultToolbarViewOverrides(v *ToolbarView) ToolbarViewOverrides {
	return ToolbarViewOverrides{}
}

// ToolbarView: widget containing a page, as well as top and/or bottom bars.
//
// <picture> <source srcset="toolbar-view-dark.png"
// media="(prefers-color-scheme: dark)"> <img src="toolbar-view.png"
// alt="toolbar-view"> </picture>
//
// AdwToolbarView has a single content widget and one or multiple top and bottom
// bars, shown at the top and bottom sides respectively.
//
// Example of an AdwToolbarView UI definition:
//
//    <object class="AdwToolbarView">
//      <child type="top">
//        <object class="AdwHeaderBar"/>
//      </child>
//      <property name="content">
//        <object class="AdwPreferencesPage">
//          <!-- ... -->
//        </object>
//      </property>
//    </object>
//
// The following kinds of top and bottom bars are supported:
//
// - headerbar
//
// - tabbar
//
// - viewswitcherbar
//
// - gtk.ActionBar
//
// - gtk.HeaderBar
//
// - gtk.PopoverMenuBar
//
// - gtk.SearchBar
//
// - Any gtk.Box or a similar widget with the .toolbar
// (style-classes.html#toolbars) style class
//
// By default, top and bottom bars are flat and scrolling content has
// a subtle undershoot shadow, same as when using the .undershoot-top
// (style-classes.html#undershot-indicators) and .undershoot-bottom
// (style-classes.html#undershot-indicators) style classes. This works well in
// most cases, e.g. with statuspage or preferencespage, where the background at
// the top and bottom parts of the page is uniform. Additionally, windows with
// sidebars should always use this style.
//
// toolbarview:top-bar-style and toolbarview:bottom-bar-style properties
// can be used add an opaque background and a persistent shadow to top
// and bottom bars, this can be useful for content such as utility panes
// (https://developer.gnome.org/hig/patterns/containers/utility-panes.html),
// where some elements are adjacent to the top/bottom bars, or tabview, where
// each page can have a different background.
//
// <picture style="min-width: 33%; display: inline-block;"> <source
// srcset="toolbar-view-flat-1-dark.png" media="(prefers-color-scheme:
// dark)"> <img src="toolbar-view-flat-1.png" alt="toolbar-view-flat-1">
// </picture> <picture style="min-width: 33%; display: inline-block;"> <source
// srcset="toolbar-view-flat-2-dark.png" media="(prefers-color-scheme:
// dark)"> <img src="toolbar-view-flat-2.png" alt="toolbar-view-flat-2">
// </picture> <picture style="min-width: 33%; display: inline-block;"> <source
// srcset="toolbar-view-raised-dark.png" media="(prefers-color-scheme: dark)">
// <img src="toolbar-view-raised.png" alt="toolbar-view-raised"> </picture>
//
// AdwToolbarView ensures the top and bottom bars have consistent backdrop
// styles and vertical spacing. For comparison:
//
// <picture style="min-width: 40%; display: inline-block;"> <source
// srcset="toolbar-view-spacing-dark.png" media="(prefers-color-scheme:
// dark)"> <img src="toolbar-view-spacing.png" alt="toolbar-view-spacing">
// </picture> <picture style="min-width: 40%; display:
// inline-block;"> <source srcset="toolbar-view-spacing-box-dark.png"
// media="(prefers-color-scheme: dark)"> <img src="toolbar-view-spacing-box.png"
// alt="toolbar-view-spacing-box"> </picture>
//
// Any top and bottom bars can also be dragged to move the window, equivalent to
// putting them into a gtk.WindowHandle.
//
// Content is typically place between top and bottom bars,
// but can also extend behind them. This is controlled
// with the toolbarview:extend-content-to-top-edge and
// toolbarview:extend-content-to-bottom-edge properties.
//
// Top and bottom bars can be hidden and revealed with an animation using the
// toolbarview:reveal-top-bars and toolbarview:reveal-bottom-bars properties.
//
// # AdwToolbarView as GtkBuildable
//
// The AdwToolbarView implementation of the gtk.Buildable interface supports
// adding a top bar by specifying “top” as the “type” attribute of a <child>
// element, or adding a bottom bar by specifying “bottom”.
//
// # Accessibility
//
// AdwToolbarView uses the GTK_ACCESSIBLE_ROLE_GROUP role.
type ToolbarView struct {
	_ [0]func() // equal guard
	gtk.Widget
}

var (
	_ gtk.Widgetter = (*ToolbarView)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*ToolbarView, *ToolbarViewClass, ToolbarViewOverrides](
		GTypeToolbarView,
		initToolbarViewClass,
		wrapToolbarView,
		defaultToolbarViewOverrides,
	)
}

func initToolbarViewClass(gclass unsafe.Pointer, overrides ToolbarViewOverrides, classInitFunc func(*ToolbarViewClass)) {
	if classInitFunc != nil {
		class := (*ToolbarViewClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapToolbarView(obj *coreglib.Object) *ToolbarView {
	return &ToolbarView{
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
	}
}

func marshalToolbarView(p uintptr) (interface{}, error) {
	return wrapToolbarView(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewToolbarView creates a new AdwToolbarView.
//
// The function returns the following values:
//
//   - toolbarView: newly created AdwToolbarView.
//
func NewToolbarView() *ToolbarView {
	var _cret *C.GtkWidget // in

	_cret = C.adw_toolbar_view_new()

	var _toolbarView *ToolbarView // out

	_toolbarView = wrapToolbarView(coreglib.Take(unsafe.Pointer(_cret)))

	return _toolbarView
}

// AddBottomBar adds a bottom bar to self.
//
// The function takes the following parameters:
//
//   - widget: widget.
//
func (self *ToolbarView) AddBottomBar(widget gtk.Widgetter) {
	var _arg0 *C.AdwToolbarView // out
	var _arg1 *C.GtkWidget      // out

	_arg0 = (*C.AdwToolbarView)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	C.adw_toolbar_view_add_bottom_bar(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(widget)
}

// AddTopBar adds a top bar to self.
//
// The function takes the following parameters:
//
//   - widget: widget.
//
func (self *ToolbarView) AddTopBar(widget gtk.Widgetter) {
	var _arg0 *C.AdwToolbarView // out
	var _arg1 *C.GtkWidget      // out

	_arg0 = (*C.AdwToolbarView)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	C.adw_toolbar_view_add_top_bar(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(widget)
}

// BottomBarHeight gets the current bottom bar height for self.
//
// Bottom bar height does change depending on toolbarview:reveal-bottom-bars,
// including during the transition.
//
// See toolbarview.GetTopBarHeight.
//
// The function returns the following values:
//
//   - gint: current bottom bar height.
//
func (self *ToolbarView) BottomBarHeight() int {
	var _arg0 *C.AdwToolbarView // out
	var _cret C.int             // in

	_arg0 = (*C.AdwToolbarView)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_toolbar_view_get_bottom_bar_height(_arg0)
	runtime.KeepAlive(self)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// BottomBarStyle gets appearance of the botom bars for self.
//
// The function returns the following values:
//
//   - toolbarStyle: bottom bar style.
//
func (self *ToolbarView) BottomBarStyle() ToolbarStyle {
	var _arg0 *C.AdwToolbarView // out
	var _cret C.AdwToolbarStyle // in

	_arg0 = (*C.AdwToolbarView)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_toolbar_view_get_bottom_bar_style(_arg0)
	runtime.KeepAlive(self)

	var _toolbarStyle ToolbarStyle // out

	_toolbarStyle = ToolbarStyle(_cret)

	return _toolbarStyle
}

// Content gets the content widget for self.
//
// The function returns the following values:
//
//   - widget (optional): content widget.
//
func (self *ToolbarView) Content() gtk.Widgetter {
	var _arg0 *C.AdwToolbarView // out
	var _cret *C.GtkWidget      // in

	_arg0 = (*C.AdwToolbarView)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_toolbar_view_get_content(_arg0)
	runtime.KeepAlive(self)

	var _widget gtk.Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

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
	}

	return _widget
}

// ExtendContentToBottomEdge gets whether the content widget can extend behind
// bottom bars.
//
// The function returns the following values:
//
//   - ok: whether content extends behind bottom bars.
//
func (self *ToolbarView) ExtendContentToBottomEdge() bool {
	var _arg0 *C.AdwToolbarView // out
	var _cret C.gboolean        // in

	_arg0 = (*C.AdwToolbarView)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_toolbar_view_get_extend_content_to_bottom_edge(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ExtendContentToTopEdge gets whether the content widget can extend behind top
// bars.
//
// The function returns the following values:
//
//   - ok: whether content extends behind top bars.
//
func (self *ToolbarView) ExtendContentToTopEdge() bool {
	var _arg0 *C.AdwToolbarView // out
	var _cret C.gboolean        // in

	_arg0 = (*C.AdwToolbarView)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_toolbar_view_get_extend_content_to_top_edge(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RevealBottomBars gets whether bottom bars are revealed for self.
//
// The function returns the following values:
//
//   - ok: whether bottom bars are revealed.
//
func (self *ToolbarView) RevealBottomBars() bool {
	var _arg0 *C.AdwToolbarView // out
	var _cret C.gboolean        // in

	_arg0 = (*C.AdwToolbarView)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_toolbar_view_get_reveal_bottom_bars(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RevealTopBars gets whether top bars are revealed for self.
//
// The function returns the following values:
//
//   - ok: whether top bars are revealed.
//
func (self *ToolbarView) RevealTopBars() bool {
	var _arg0 *C.AdwToolbarView // out
	var _cret C.gboolean        // in

	_arg0 = (*C.AdwToolbarView)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_toolbar_view_get_reveal_top_bars(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TopBarHeight gets the current top bar height for self.
//
// Top bar height does change depending on toolbarview:reveal-top-bars,
// including during the transition.
//
// See toolbarview.GetBottomBarHeight.
//
// The function returns the following values:
//
//   - gint: current top bar height.
//
func (self *ToolbarView) TopBarHeight() int {
	var _arg0 *C.AdwToolbarView // out
	var _cret C.int             // in

	_arg0 = (*C.AdwToolbarView)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_toolbar_view_get_top_bar_height(_arg0)
	runtime.KeepAlive(self)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// TopBarStyle gets appearance of the top bars for self.
//
// The function returns the following values:
//
//   - toolbarStyle: top bar style.
//
func (self *ToolbarView) TopBarStyle() ToolbarStyle {
	var _arg0 *C.AdwToolbarView // out
	var _cret C.AdwToolbarStyle // in

	_arg0 = (*C.AdwToolbarView)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_toolbar_view_get_top_bar_style(_arg0)
	runtime.KeepAlive(self)

	var _toolbarStyle ToolbarStyle // out

	_toolbarStyle = ToolbarStyle(_cret)

	return _toolbarStyle
}

// Remove removes a child from self.
//
// The function takes the following parameters:
//
//   - widget: child to be removed.
//
func (self *ToolbarView) Remove(widget gtk.Widgetter) {
	var _arg0 *C.AdwToolbarView // out
	var _arg1 *C.GtkWidget      // out

	_arg0 = (*C.AdwToolbarView)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	C.adw_toolbar_view_remove(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(widget)
}

// SetBottomBarStyle sets appearance of the bottom bars for self.
//
// If set to ADW_TOOLBAR_FLAT, bottom bars are flat and scrolling content has a
// subtle undershoot shadow when touching them, same as the .undershoot-bottom
// (style-classes.html#undershot-indicators) style class. This works well for
// simple content, e.g. statuspage or preferencespage, where the background at
// the bottom of the page is uniform. Additionally, windows with sidebars should
// always use this style.
//
// Undershoot shadow is only present if a bottom bar is
// actually present and visible. It is also never present if
// toolbarview:extend-content-to-bottom-edge is set to TRUE.
//
// If set to ADW_TOOLBAR_RAISED, bottom bars have an opaque background and
// a persistent shadow, this is suitable for content such as utility panes
// (https://developer.gnome.org/hig/patterns/containers/utility-panes.html),
// where some elements are directly adjacent to the bottom bars, or tabview,
// where each page can have a different background.
//
// ADW_TOOLBAR_RAISED_BORDER is similar to ADW_TOOLBAR_RAISED, but the shadow is
// replaced with a more subtle border. This can be useful for applications like
// image viewers.
//
// See also toolbarview.SetTopBarStyle.
//
// The function takes the following parameters:
//
//   - style: bottom bar style.
//
func (self *ToolbarView) SetBottomBarStyle(style ToolbarStyle) {
	var _arg0 *C.AdwToolbarView // out
	var _arg1 C.AdwToolbarStyle // out

	_arg0 = (*C.AdwToolbarView)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.AdwToolbarStyle(style)

	C.adw_toolbar_view_set_bottom_bar_style(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(style)
}

// SetContent sets the content widget for self.
//
// The function takes the following parameters:
//
//   - content (optional) widget.
//
func (self *ToolbarView) SetContent(content gtk.Widgetter) {
	var _arg0 *C.AdwToolbarView // out
	var _arg1 *C.GtkWidget      // out

	_arg0 = (*C.AdwToolbarView)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if content != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(content).Native()))
	}

	C.adw_toolbar_view_set_content(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(content)
}

// SetExtendContentToBottomEdge sets whether the content widget can extend
// behind bottom bars.
//
// This can be used in combination with toolbarview:reveal-bottom-bars to show
// and hide toolbars in fullscreen.
//
// See toolbarview.SetExtendContentToTopEdge.
//
// The function takes the following parameters:
//
//   - extend: whether content extends behind bottom bars.
//
func (self *ToolbarView) SetExtendContentToBottomEdge(extend bool) {
	var _arg0 *C.AdwToolbarView // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.AdwToolbarView)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if extend {
		_arg1 = C.TRUE
	}

	C.adw_toolbar_view_set_extend_content_to_bottom_edge(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(extend)
}

// SetExtendContentToTopEdge sets whether the content widget can extend behind
// top bars.
//
// This can be used in combination with toolbarview:reveal-top-bars to show and
// hide toolbars in fullscreen.
//
// See toolbarview.SetExtendContentToBottomEdge.
//
// The function takes the following parameters:
//
//   - extend: whether content extends behind top bars.
//
func (self *ToolbarView) SetExtendContentToTopEdge(extend bool) {
	var _arg0 *C.AdwToolbarView // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.AdwToolbarView)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if extend {
		_arg1 = C.TRUE
	}

	C.adw_toolbar_view_set_extend_content_to_top_edge(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(extend)
}

// SetRevealBottomBars sets whether bottom bars are revealed for self.
//
// The transition will be animated.
//
// This can be used in combination with
// toolbarview:extend-content-to-bottom-edge to show and hide toolbars in
// fullscreen.
//
// See toolbarview.SetRevealTopBars.
//
// The function takes the following parameters:
//
//   - reveal: whether to reveal bottom bars.
//
func (self *ToolbarView) SetRevealBottomBars(reveal bool) {
	var _arg0 *C.AdwToolbarView // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.AdwToolbarView)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if reveal {
		_arg1 = C.TRUE
	}

	C.adw_toolbar_view_set_reveal_bottom_bars(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(reveal)
}

// SetRevealTopBars sets whether top bars are revealed for self.
//
// The transition will be animated.
//
// This can be used in combination with toolbarview:extend-content-to-top-edge
// to show and hide toolbars in fullscreen.
//
// See toolbarview.SetRevealBottomBars.
//
// The function takes the following parameters:
//
//   - reveal: whether to reveal top bars.
//
func (self *ToolbarView) SetRevealTopBars(reveal bool) {
	var _arg0 *C.AdwToolbarView // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.AdwToolbarView)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if reveal {
		_arg1 = C.TRUE
	}

	C.adw_toolbar_view_set_reveal_top_bars(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(reveal)
}

// SetTopBarStyle sets appearance of the top bars for self.
//
// If set to ADW_TOOLBAR_FLAT, top bars are flat and scrolling content has a
// subtle undershoot shadow when touching them, same as the .undershoot-top
// (style-classes.html#undershot-indicators) style class. This works well for
// simple content, e.g. statuspage or preferencespage, where the background at
// the top of the page is uniform. Additionally, windows with sidebars should
// always use this style.
//
// Undershoot shadow is only present if a top bar is actually present and
// visible. It is also never present if toolbarview:extend-content-to-top-edge
// is set to TRUE.
//
// If set to ADW_TOOLBAR_RAISED, top bars have an opaque background and a
// persistent shadow, this is suitable for content such as utility panes
// (https://developer.gnome.org/hig/patterns/containers/utility-panes.html),
// where some elements are directly adjacent to the top bars, or tabview,
// where each page can have a different background.
//
// ADW_TOOLBAR_RAISED_BORDER is similar to ADW_TOOLBAR_RAISED, but the shadow is
// replaced with a more subtle border. This can be useful for applications like
// image viewers.
//
// See also toolbarview.SetBottomBarStyle.
//
// The function takes the following parameters:
//
//   - style: top bar style.
//
func (self *ToolbarView) SetTopBarStyle(style ToolbarStyle) {
	var _arg0 *C.AdwToolbarView // out
	var _arg1 C.AdwToolbarStyle // out

	_arg0 = (*C.AdwToolbarView)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.AdwToolbarStyle(style)

	C.adw_toolbar_view_set_top_bar_style(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(style)
}
