package loggingfs

import (
	"fmt"
	fuse "github.com/hanwen/go-fuse/fuse"
	osusers "os/user"
	"time"
)

/*
 Logical write-only file that forwards completed lines to remote logging system
*/

type LogFile struct {
	fuse.DefaultFile

	name string

	pid      uint32
	uid      uint32
	username string
	gid      uint32

	last_write time.Time

	buffer    []byte
	transport Transport
}

func NewLogFile(name string, context *fuse.Context) (fs *LogFile, err error) {
	file := new(LogFile)

	file.name = name

	file.pid = context.Pid
	file.uid = context.Uid
	user, err := osusers.LookupId(fmt.Sprint(file.uid))
	if err == nil {
		file.username = user.Username
	}
	file.gid = context.Gid

	return file, nil
}

func (f *LogFile) String() string {
	// TODO: include log params
	return "LogFile"
}

func (f *LogFile) Read(buf []byte, off int64) (fuse.ReadResult, fuse.Status) {
	// TODO: ? maybe make this a readable buffer or status message
	return nil, fuse.EPERM
}

func (f *LogFile) Write(data []byte, off int64) (uint32, fuse.Status) {
	fmt.Println(string(data))
	return uint32(len(data)), fuse.OK
}

func (f *LogFile) Flush() fuse.Status {
	return fuse.OK
}

// func (f *LogFile) GetAttr(*fuse.Attr) fuse.Status {
// 	return fuse.ENOSYS
// }

func (f *LogFile) Allocate(off uint64, size uint64, mode uint32) fuse.Status {
	return fuse.ENOSYS
}
