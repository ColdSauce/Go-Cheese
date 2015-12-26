package Gotem

import (
	"errors"
	"syscall"
	"unsafe"
)

var (
	user32               *syscall.LazyDLL = syscall.NewLazyDLL("user32.dll")
	setPhysicalCursorPos                  = user32.NewProc("SetPhysicalCursorPos")
	getPhysicalCursorPos                  = lazyDll.NewProc("GetPhysicalCursorPos")
	keybd_event                           = lazyDll.NewProc("keybd_event")
)

func MoveMouseTo(p Point) error {
	ret, _, _ := setPhysicalCursorPos.Call(uintptr(p.X), uintptr(p.Y))
	if ret != 0 {
		return errors.New("Couldn't set cursor position.")
	}
}

func GetMousePosition() (Point, error) {
	type mousePosition struct {
		x, y int32
	}
	var ran mosPoint
	ret, _, _ := getPhysicalCursorPos.Call(uintptr(unsafe.Pointer(&ran)))
	if ret != 0 {
		return errors.New("Couldn't get cursor position.")
	}
	return Point{int(ran.x), int(ran.y)}, nil
}

func PressKey(key uintptr, isUp bool) error {
	var dwFlags uintptr

	if isUp {
		dwFlags = 0x02
	} else {
		dwFlags = 0x0
	}
	ret, _, _ := keybd_event.Call(key, 0, dwFlags, 0)
	if ret != 0 {
		return errors.New("Couldn't press key!.")
	}
}
