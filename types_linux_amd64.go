// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs types.go

package svipc

type IpcKey int32

type IpcPerm struct {
	Key  int32
	Uid  uint32
	Gid  uint32
	Cuid uint32
	Cgid uint32
	Mode uint16
	_    uint16
	Seq  uint16
	_    uint16
	_    [4]byte
	_    uint64
	_    uint64
}

type ShmId struct {
	Perm   IpcPerm
	Segsz  uint64
	Atime  int64
	Dtime  int64
	Ctime  int64
	Cpid   int32
	Lpid   int32
	Nattch uint64
	_      uint64
	_      uint64
}

type ShmInfo struct {
	Shmmax uint64
	Shmmin uint64
	Shmmni uint64
	Shmseg uint64
	Shmall uint64
	_      uint64
	_      uint64
	_      uint64
	_      uint64
}

type SemId struct {
	Perm  IpcPerm
	Otime int64
	_     uint64
	Ctime int64
	_     uint64
	Nsems uint64
	_     uint64
	_     uint64
}
