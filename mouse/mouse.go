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

func (this *Mouse) GetX () int32 {
	return this.x
}

func (this *Mouse) GetY () int32 {
	return this.y
}

func (this *Mouse) GetXY () (int32, int32) {
	return this.x, this.y
}

func (this *Mouse) SetXY (x, y int32) {
	this.x = x
	this.y = y
}

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
