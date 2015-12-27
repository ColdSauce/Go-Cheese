package Gotem

// #cgo LDFLAGS: -lX11
// #include <X11/Xlib.h>
// #include  <X11/extensions/XTest.h>
import "C"

var( dpy = C.XOpenDisplay(nil))

// MoveMouseTo (X Window System version) moves the user's mouse to a specified point on the screen.
func MoveMouseTo(p Point) error{
	// TODO: Implement this.
}

// GetMousePosition (X Window System version) returns a Point indicating where the cursor is located on the screen.
func GetMousePosition() (Point, error){
	// TODO: Implement this.
	return Point{0, 0}
}

// PressKey (X Window System version) simulates a key press.
func PressKey(key uintptr, isUp bool) error {
	// TODO: Implement this.
}
