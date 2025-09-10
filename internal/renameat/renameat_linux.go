package renameat

import "golang.org/x/sys/unix"

func renameat(olddirfd int, oldpath string, newdirfd int, newpath string, flags uint) (err error) {
	return unix.Renameat2(olddirfd, oldpath, newdirfd, newpath, flags)
}
