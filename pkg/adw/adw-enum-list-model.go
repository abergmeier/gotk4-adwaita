// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: libadwaita-1
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <adwaita.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.adw_enum_list_item_get_type()), F: marshalEnumListItemmer},
		{T: externglib.Type(C.adw_enum_list_model_get_type()), F: marshalEnumListModeller},
	})
}

// EnumListItem: AdwEnumListItem is the type of items in a enumlistmodel.
type EnumListItem struct {
	*externglib.Object
}

var (
	_ externglib.Objector = (*EnumListItem)(nil)
)

func wrapEnumListItem(obj *externglib.Object) *EnumListItem {
	return &EnumListItem{
		Object: obj,
	}
}

func marshalEnumListItemmer(p uintptr) (interface{}, error) {
	return wrapEnumListItem(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Name gets the enum value name.
func (self *EnumListItem) Name() string {
	var _arg0 *C.AdwEnumListItem // out
	var _cret *C.char            // in

	_arg0 = (*C.AdwEnumListItem)(unsafe.Pointer(self.Native()))

	_cret = C.adw_enum_list_item_get_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Nick gets the enum value nick.
func (self *EnumListItem) Nick() string {
	var _arg0 *C.AdwEnumListItem // out
	var _cret *C.char            // in

	_arg0 = (*C.AdwEnumListItem)(unsafe.Pointer(self.Native()))

	_cret = C.adw_enum_list_item_get_nick(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Value gets the enum value.
func (self *EnumListItem) Value() int {
	var _arg0 *C.AdwEnumListItem // out
	var _cret C.int              // in

	_arg0 = (*C.AdwEnumListItem)(unsafe.Pointer(self.Native()))

	_cret = C.adw_enum_list_item_get_value(_arg0)
	runtime.KeepAlive(self)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// EnumListModel: gio.ListModel representing values of a given enum.
//
// AdwEnumListModel contains objects of type enumlistitem.
type EnumListModel struct {
	*externglib.Object

	gio.ListModel
}

var (
	_ externglib.Objector = (*EnumListModel)(nil)
)

func wrapEnumListModel(obj *externglib.Object) *EnumListModel {
	return &EnumListModel{
		Object: obj,
		ListModel: gio.ListModel{
			Object: obj,
		},
	}
}

func marshalEnumListModeller(p uintptr) (interface{}, error) {
	return wrapEnumListModel(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewEnumListModel creates a new AdwEnumListModel for enum_type.
//
// The function takes the following parameters:
//
//    - enumType: type of the enum to construct the model from.
//
func NewEnumListModel(enumType externglib.Type) *EnumListModel {
	var _arg1 C.GType             // out
	var _cret *C.AdwEnumListModel // in

	_arg1 = C.GType(enumType)

	_cret = C.adw_enum_list_model_new(_arg1)
	runtime.KeepAlive(enumType)

	var _enumListModel *EnumListModel // out

	_enumListModel = wrapEnumListModel(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _enumListModel
}

// FindPosition finds the position of a given enum value in self.
//
// The function takes the following parameters:
//
//    - value: enum value.
//
func (self *EnumListModel) FindPosition(value int) uint {
	var _arg0 *C.AdwEnumListModel // out
	var _arg1 C.int               // out
	var _cret C.guint             // in

	_arg0 = (*C.AdwEnumListModel)(unsafe.Pointer(self.Native()))
	_arg1 = C.int(value)

	_cret = C.adw_enum_list_model_find_position(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// EnumType gets the type of the enum represented by self.
func (self *EnumListModel) EnumType() externglib.Type {
	var _arg0 *C.AdwEnumListModel // out
	var _cret C.GType             // in

	_arg0 = (*C.AdwEnumListModel)(unsafe.Pointer(self.Native()))

	_cret = C.adw_enum_list_model_get_enum_type(_arg0)
	runtime.KeepAlive(self)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}
