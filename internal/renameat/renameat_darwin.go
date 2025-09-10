package renameat

import (
	"syscall"
	"unsafe"
)

const SYS_RENAMEATX_NP = 488

func renameat(olddirfd int, oldpath string, newdirfd int, newpath string, flags uint) error {
	oldpathCString, err := syscall.BytePtrFromString(oldpath)
	if err != nil {
		return err
	}
	newpathCString, err := syscall.BytePtrFromString(newpath)
	if err != nil {
		return err
	}

	_, _, errno := syscall.Syscall6(
		SYS_RENAMEATX_NP,
		uintptr(olddirfd),
		uintptr(unsafe.Pointer(oldpathCString)),
		uintptr(newdirfd),
		uintptr(unsafe.Pointer(newpathCString)),
		uintptr(flags),
		0,
	)

	if errno != 0 {
		return errno
	}
	return nil
}
