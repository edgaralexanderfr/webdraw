package mouse

import (
	"syscall"
	"unsafe"
)

// LPPOINT
type Mouse struct {
	x, y int32
}

var (
	_user32       = syscall.NewLazyDLL("user32.dll")
	_GetCursorPos = _user32.NewProc("GetCursorPos")
	_SetCursorPos = _user32.NewProc("SetCursorPos")
)

func New () *Mouse {
	return &Mouse{}
}

func (this *Mouse) Get () uintptr {
	success, _, _ := _GetCursorPos.Call(uintptr(unsafe.Pointer(this)))

	return success
}

func (this *Mouse) Set () uintptr {
	success, _, _ := _SetCursorPos.Call(uintptr(this.x), uintptr(this.y))

	return success
}
