package Gotem
import (
	"syscall"
	"unsafe"
)

type Point struct {
	x, y int
}

var (
	// lazyDll acts as a sort of singleton.
	lazyDll *syscall.LazyDLL
)



func MoveMouseTo(p Point) {
	if lazyDll == nil {
		lazyDll = syscall.NewLazyDLL("user32.dll")
	}
	proc := lazyDll.NewProc("SetPhysicalCursorPos")
	proc.Call(uintptr(p.x),uintptr(p.y))
}

func GetMousePosition() Point{
	type mosPoint struct {
		x, y int32
	}
	if lazyDll == nil {
		lazyDll = syscall.NewLazyDLL("user32.dll")
	}
	var ran mosPoint
	proc := lazyDll.NewProc("GetPhysicalCursorPos")
	proc.Call(uintptr(unsafe.Pointer(&ran)))
	return Point{int(ran.x),int(ran.y)}

}

func PressKey(key uintptr, isUp bool) {
	if lazyDll == nil {
		lazyDll = syscall.NewLazyDLL("user32.dll")
	}
	// The reason I am using keybd_event instead of SendInput is because I have no idea how to create
	// the Input structure that is needed to be passed in using Go since it uses a union.
	proc := lazyDll.NewProc("keybd_event")
	var dwFlags uintptr
	if isUp {
		dwFlags = 0x02
	} else {
		dwFlags = 0x0
	}
	proc.Call(key, 0, dwFlags, 0)
}

// for some reason I need a main even if it's a library?
func main() {}
