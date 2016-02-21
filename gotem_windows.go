// Package Gotem implements a simple library for programatically controlling one's keyboard and mouse.
package Gotem

import (
	"errors"
	"syscall"
	"unsafe"
)

var (
	user32               *syscall.LazyDLL = syscall.NewLazyDLL("user32.dll")
	setPhysicalCursorPos                  = user32.NewProc("SetPhysicalCursorPos")
	getPhysicalCursorPos                  = user32.NewProc("GetPhysicalCursorPos")
	w32SendInput 			      = user32.NewProc("SendInput")
)

// MoveMouseTo (Windows version) moves the user's mouse to a specified point on the screen.
func MoveMouseTo(p *Point) error {
	ret, _, _ := setPhysicalCursorPos.Call(uintptr(unsafe.Pointer(&p.X)), uintptr(unsafe.Pointer(&p.Y)))
	if ret == 0 {
		return errors.New("Couldn't set cursor position.")
	} else {
		return nil
	}
}

// GetMousePosition (Windows version) returns a Point indicating where the cursor is located on the screen.
func GetMousePosition() (*Point, error) {
	type mousePosition struct {
		x, y int32
	}
	var ran mousePosition
	ret,_,_:= getPhysicalCursorPos.Call(uintptr(unsafe.Pointer(&ran)))
	if ret == 0 {
		return nil, errors.New("Couldn't get cursor position.")
	}
	return &Point{int(ran.x), int(ran.y)}, nil
}

func sendInput(input Input) uint {
	var someInputArray []Input
	someInputArray = append(someInputArray, input)
	SIZE_OF_C_INPUT := 0x28
	ret, _, _ := w32SendInput.Call(uintptr(len(someInputArray)), uintptr(unsafe.Pointer(&someInputArray[0])), uintptr(SIZE_OF_C_INPUT))
	return uint(ret)
}

// PressKey (Windows version) simulates a key press.
func PressKey(key Key, isUp bool) uint {
	var dwFlagsToUse uint32

	if isUp {
		dwFlagsToUse = 0x02
	} else {
		dwFlagsToUse = 0x0
	}
	s := sendInput(Input {
			type_:KEYBOARD_INPUT,
			ki:Keybdinput{
				wVk: uint16(key),
				dwFlags: dwFlagsToUse,
			},
		},
	)
	return s
}
