// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #cgo pkg-config: libadwaita-1
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <adwaita.h>
import "C"

// Init calls adw.Init then restores the user's theme preference. This
// function prevents adw.Init from overriding the user's theme to
// Adwaita.
//
// For more information, see
// https://gitlab.gnome.org/GNOME/libadwaita/-/issues/215.
func InitPreserveTheme() {
	adw.Init()
	settings := gtk.SettingsGetDefault()
	settings.ResetProperty("gtk-theme-name")
}
