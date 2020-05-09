package win32

import (
	"syscall"
	"unsafe"
)

var (
	kernel32DLL          = syscall.NewLazyDLL("kernel32.dll")
	procCreateJobObjectA = kernel32DLL.NewProc("CreateJobObjectA")
)

// CreateJobObject uses the CreateJobObjectA Windows API Call to create and return a Handle to a new JobObject
func CreateJobObject(attr *syscall.SecurityAttributes, name string) (syscall.Handle, error) {
	r1, _, err := procCreateJobObjectA.Call(
		uintptr(unsafe.Pointer(attr)),
		uintptr(unsafe.Pointer(StringToCharPtr(name))),
	)
	if err != syscall.Errno(0) {
		return 0, err
	}
	return syscall.Handle(r1), nil
}
