package svipc

// mksyscall.pl syscall_linux.go syscall_linux_386.go

//sys Shmget(key int, size int, flags int) (id uintptr, err error)
//sys Shmat(id int, targaddr uintptr, flags int) (addr unsafe.Pointer, err error)
//sys Shmctl(id int, cmd uintptr, shmid *ShmId) (err error)
//sys Shmdt(addr unsafe.Pointer) (err error)

//sys Semget(key int, nsems int, flags int) (id uintptr, err error)
//sys Semop(id int, buf unsafe.Pointer, n int) (err error)
//sys Semctl(id uintptr, num int, cmd int) (ret int, err error)
