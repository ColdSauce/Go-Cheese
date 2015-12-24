package Gotem

import (
	"syscall"
	"unsafe"
)

var (
	user32               *syscall.LazyDLL = syscall.NewLazyDLL("user32.dll")
	setPhysicalCursorPos                  = user32.NewProc("SetPhysicalCursorPos")
	getPhysicalCursorPos                  = user32.NewProc("GetPhysicalCursorPos")
	keybd_event                           = user32.NewProc("keybd_event")
)

func MoveMouseTo(p Point) error {
	setPhysicalCursorPos.Call(uintptr(p.X), uintptr(p.Y))
}

func GetMousePosition() (Point, error) {
	type mousePosition struct {
		x, y int32
	}
	var ran mousePosition
	getPhysicalCursorPos.Call(uintptr(unsafe.Pointer(&ran)))
	return Point{int(ran.x), int(ran.y)}
}

func PressKey(key uintptr, isUp bool) error {
	var dwFlags uintptr

	if isUp {
		dwFlags = 0x02
	} else {
		dwFlags = 0x0
	}
	keybd_event.Call(key, 0, dwFlags, 0)
}
