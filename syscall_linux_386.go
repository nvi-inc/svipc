package svipc

import (
	"unsafe"

	"golang.org/x/sys/unix"
)

const (
	IPCOP_SEMOP      = 1
	IPCOP_SEMGET     = 2
	IPCOP_SEMCTL     = 3
	IPCOP_SEMTIMEDOP = 4
	IPCOP_MSGSND     = 11
	IPCOP_MSGRCV     = 12
	IPCOP_MSGGET     = 13
	IPCOP_MSGCTL     = 14
	IPCOP_SHMAT      = 21
	IPCOP_SHMDT      = 22
	IPCOP_SHMGET     = 23
	IPCOP_SHMCTL     = 24

	IPC_64 = 0x0100
)

//SEM

func Semget(key IpcKey, nsems int, flags int) (id int, err error) {
	ret, _, errno := unix.Syscall6(unix.SYS_IPC, IPCOP_SEMGET, uintptr(key), uintptr(nsems), uintptr(flags), 0, 0)
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

//SHM

func Shmget(key IpcKey, size uintptr, flag int) (shmid uintptr, err error) {
	shmid, _, errno := unix.Syscall6(unix.SYS_IPC, IPCOP_SHMGET, uintptr(key), size, uintptr(flag), 0, 0)
	if errno != 0 {
		err = errno
	}
	return
}

func Shmat(id uintptr, addr uintptr, flag int) (ptr unsafe.Pointer, err error) {
	_, _, errno := unix.Syscall6(unix.SYS_IPC, IPCOP_SHMAT, id, uintptr(flag), uintptr(unsafe.Pointer(&addr)), addr, 0)
	if errno != 0 {
		err = errno
	}
	ptr = unsafe.Pointer(addr)
	return
}

func Shmctl(id uintptr, cmd int, shmid *ShmId) (err error) {
	ptr := uintptr(unsafe.Pointer(shmid))
	_, _, errno := unix.Syscall6(unix.SYS_IPC, IPCOP_SHMCTL, id, uintptr(cmd|IPC_64), 0, ptr, 0)
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
