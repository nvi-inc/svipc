// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs types.go

package svipc

type IpcKey int32

type IpcPerm struct {
	Key  IpcKey
	Uid  uint32
	Gid  uint32
	Cuid uint32
	Cgid uint32
	Mode uint16
	_    uint16
	Seq  uint16
	_    uint16
	_    uint32
	_    uint32
}

type ShmId struct {
	Perm   IpcPerm
	Segsz  uint32
	Atime  int32
	_      uint32
	Dtime  int32
	_      uint32
	Ctime  int32
	_      uint32
	Cpid   int32
	Lpid   int32
	Nattch uint32
	_      uint32
	_      uint32
}

type ShmInfo struct {
	Shmmax uint32
	Shmmin uint32
	Shmmni uint32
	Shmseg uint32
	Shmall uint32
	_      uint32
	_      uint32
	_      uint32
	_      uint32
}

type SemId struct {
	Perm  IpcPerm
	Otime int32
	_     uint32
	Ctime int32
	_     uint32
	Nsems uint32
	_     uint32
	_     uint32
}
