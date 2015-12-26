package Gotem

// #cgo LDFLAGS: -lX11
// #include <X11/Xlib.h>
// #include  <X11/extensions/XTest.h>
import "C"

var( dpy = C.XOpenDisplay(nil))

func MoveMouseTo(p Point) {
	// TODO: Implement this.
}

func GetMousePosition() Point {
	// TODO: Implement this.
	return Point{0, 0}
}

func PressKey(key uintptr, isUp bool) {
	// TODO: Implement this.
}
