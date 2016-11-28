package svipc

import (
	"unsafe"

	"golang.org/x/sys/unix"
)

const (
	SEM_UNDO = 0x1000
	GETPID   = 11
	GETVAL   = 12
	GETALL   = 13
	GETNCNT  = 14
	GETZCNT  = 15
	SETVAL   = 16
	SETALL   = 17

	SEM_STAT = 18
	SEM_INFO = 19
)

func Semget(key IpcKey, nsems uintptr, flags uintptr) (id int, err error) {
	ret, _, errno := unix.Syscall6(unix.SYS_IPC, IPCOP_SEMGET, uintptr(key), nsems, flags, 0, 0)
	id = int(ret)
	if errno != 0 {
		err = errno
	}
	return
}

func Semop(id int, buf unsafe.Pointer, n int) (err error) {
	_, _, errno := unix.Syscall6(unix.SYS_IPC, IPCOP_SEMOP, uintptr(id), uintptr(n), 0, uintptr(buf), 0)
	if errno != 0 {
		err = errno
	}
	return
}

func Semctl(id int, num int, cmd int) (n int, err error) {
	var buf uintptr //currently unused
	ret, _, errno := unix.Syscall6(unix.SYS_IPC, IPCOP_SEMCTL, uintptr(id), uintptr(num), uintptr(cmd|IPC_64), uintptr(unsafe.Pointer(&buf)), 0)
	n = int(ret)
	if errno != 0 {
		err = errno
	}
	return
}
