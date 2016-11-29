package svipc

// mksyscall.pl syscall_linux.go syscall_linux_386.go

//sys Shmget(key IpcKey, size uintptr, flags int) (id uintptr, err error)
//sys Shmat(id uintptr, targaddr uintptr, flags int) (addr unsafe.Pointer, err error)
//sys Shmctl(id uintptr, cmd int, shmid *ShmId) (err error)
//sys Shmdt(addr unsafe.Pointer) (err error)

//sys Semget(key IpcKey, nsems int, flags int) (id uintptr, err error)
//sys Semop(id uintptr, buf unsafe.Pointer, n int) (err error)
//sys Semctl(id uintptr, num int, cmd int) (ret int, err error)
