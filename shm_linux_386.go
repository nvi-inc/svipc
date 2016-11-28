package svipc

import (
	"unsafe"

	"golang.org/x/sys/unix"
)

func Shmget(key IpcKey, size uintptr, flag int) (shmid uintptr, err error) {
	shmid, _, errno := unix.Syscall6(unix.SYS_IPC, IPCOP_SHMGET, uintptr(key), size, uintptr(flag), 0, 0)
	if errno != 0 {
		err = errno
	}
	return
}

func Shmat(id uintptr, addr uintptr, flag uintptr) (ptr unsafe.Pointer, err error) {
	_, _, errno := unix.Syscall6(unix.SYS_IPC, IPCOP_SHMAT, id, flag, uintptr(unsafe.Pointer(&addr)), addr, 0)
	if errno != 0 {
		err = errno
	}
	ptr = unsafe.Pointer(addr)
	return
}

func Shmctl(id uintptr, cmd uintptr, shmid *ShmId) (err error) {
	ptr := uintptr(unsafe.Pointer(shmid))
	_, _, errno := unix.Syscall6(unix.SYS_IPC, IPCOP_SHMCTL, id, cmd|IPC_64, 0, ptr, 0)
	if errno != 0 {
		err = errno
	}
	return
}

func Shmdt(addr unsafe.Pointer) (err error) {
	_, _, errno := unix.Syscall6(unix.SYS_IPC, IPCOP_SHMDT, 0, 0, 0, uintptr(addr), 0)
	if errno != 0 {
		err = errno
	}
	return
}
