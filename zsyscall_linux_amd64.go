// mksyscall.pl syscall_linux.go syscall_linux_amd64.go
// MACHINE GENERATED BY THE COMMAND ABOVE; DO NOT EDIT

package svipc

import (
	"unsafe"

	"golang.org/x/sys/unix"
)

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Shmget(key IpcKey, size uintptr, flags int) (id uintptr, err error) {
	r0, _, e1 := unix.Syscall(unix.SYS_SHMGET, uintptr(key), uintptr(size), uintptr(flags))
	id = uintptr(r0)
	if e1 != 0 {
		err = (e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Shmat(id uintptr, tagaddr uintptr, flags int) (addr unsafe.Pointer, err error) {
	r0, _, e1 := unix.Syscall(unix.SYS_SHMAT, uintptr(id), uintptr(tagaddr), uintptr(flags))
	addr = unsafe.Pointer(r0)
	if e1 != 0 {
		err = (e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Shmctl(id uintptr, cmd uintptr, shmid *ShmId) (err error) {
	_, _, e1 := unix.Syscall(unix.SYS_SHMCTL, uintptr(id), uintptr(cmd), uintptr(unsafe.Pointer(shmid)))
	if e1 != 0 {
		err = (e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Shmdt(addr unsafe.Pointer) (err error) {
	_, _, e1 := unix.Syscall(unix.SYS_SHMDT, uintptr(addr), 0, 0)
	if e1 != 0 {
		err = (e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Semget(key IpcKey, nsems int, flags int) (id uintptr, err error) {
	r0, _, e1 := unix.Syscall(unix.SYS_SEMGET, uintptr(key), uintptr(nsems), uintptr(flags))
	id = uintptr(r0)
	if e1 != 0 {
		err = (e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Semop(id uintptr, buf unsafe.Pointer, n int) (err error) {
	_, _, e1 := unix.Syscall(unix.SYS_SEMOP, uintptr(id), uintptr(buf), uintptr(n))
	if e1 != 0 {
		err = (e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Semctl(id uintptr, num int, cmd int) (ret int, err error) {
	r0, _, e1 := unix.Syscall(unix.SYS_SEMCTL, uintptr(id), uintptr(num), uintptr(cmd))
	ret = int(r0)
	if e1 != 0 {
		err = (e1)
	}
	return
}
