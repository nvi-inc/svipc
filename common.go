package svipc

import (
	"golang.org/x/sys/unix"
)

const (
	SHM_R = 0400
	SHM_W = 0200

	SHM_RDONLY = 010000
	SHM_RND    = 020000
	SHM_REMAP  = 040000
	SHM_EXEC   = 0100000

	SHM_LOCK      = 11
	SHM_UNLOCK    = 12
	SHM_STAT      = 13
	SHM_INFO      = 14
	SHM_DEST      = 01000
	SHM_LOCKED    = 02000
	SHM_HUGETLB   = 04000
	SHM_NORESERVE = 010000

	IPC_CREAT  = 01000
	IPC_EXCL   = 02000
	IPC_NOWAIT = 04000

	IPC_RMID = 0
	IPC_SET  = 1
	IPC_STAT = 2
	IPC_INFO = 3
)

func Ftok(path string, id uint64) (IpcKey, error) {
	var st unix.Stat_t
	err := unix.Stat(path, &st)
	if err != nil {
		return 0, err
	}

	return IpcKey((st.Ino & 0xffff) | ((st.Dev & 0xff) << 16) | ((id & 0xff) << 24)), nil
}
