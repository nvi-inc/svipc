// +build ignore

package svipc

/*
#include <stddef.h>
#include <sys/types.h>
#include <sys/ipc.h>
#include <sys/shm.h>
#include <sys/sem.h>
*/
import "C"

type IpcKey C.int

type IpcPerm C.struct_ipc_perm

type ShmId C.struct_shmid_ds

//Linux only
type ShmInfo C.struct_shminfo

type SemId C.struct_semid_ds
