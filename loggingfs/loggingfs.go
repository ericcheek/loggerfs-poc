package loggingfs

import (
	//"bytes"
	//"errors"
	"fmt"
	fuse "github.com/hanwen/go-fuse/fuse"
	"os"
	"time"
)

type LoggingFs struct {
	fuse.PathNodeFs
	fuse.DefaultFileSystem

	// path -> buffer
	config map[string]LogConfig

	transports []*Transport
}

// Does not return until
func SetupAndMount(filename string) error {
	config, err := load_config(filename)

	if err != nil {
		fmt.Printf("Failed loading configuration: %v\n", err)
	}

	var finalFs fuse.FileSystem
	fs, err := NewLoggingFs(nil)
	finalFs = fs

	pathFs := fuse.NewPathNodeFs(finalFs, nil)

	if err != nil {
		fmt.Printf("Failed creating fs: %v\n", err)
	}

	conn := fuse.NewFileSystemConnector(pathFs, nil)
	state := fuse.NewMountState(conn)

	// TODO: cut
	os.Exit(2)
	mount_err := state.Mount(config.Mount_point, nil)

	if mount_err != nil {
		fmt.Printf("Mount fail: %v\n", mount_err)
		return mount_err
	}

	fmt.Println("Mounted!")
	state.Loop()

	return nil
}

func NewLoggingFs(transport *Transport) (fs *LoggingFs, err error) {
	loggerfs := new(LoggingFs)

	return loggerfs, nil
}

func (fs *LoggingFs) OnMount(nodeFS *fuse.PathNodeFs) {
}

func (fs *LoggingFs) OnUnmount() {
	fmt.Printf("Unmounting apparently...")
}

func (fs *LoggingFs) String() string {
	return "LoggingFs"
}

func (fs *LoggingFs) OpenDir(name string, context *fuse.Context) (stream []fuse.DirEntry, code fuse.Status) {

	//stream = make([]fuse.DirEntry, 0, len(fs.files)+2)
	stream = make([]fuse.DirEntry, 0, 3)

	default_log := fuse.DirEntry{Name: "test.log"}
	default_log.Mode = uint32(os.O_WRONLY | os.O_APPEND)
	stream = append(stream, default_log)

	return stream, fuse.OK
}

func (fs *LoggingFs) Open(name string, flags uint32, context *fuse.Context) (fuseFile fuse.File, code fuse.Status) {
	// TODO: restrict reads
	if flags&fuse.O_ANYWRITE == 0 {
		return nil, fuse.EPERM
	}

	file, _ := NewLogFile(name, context)

	return file, fuse.OK
}

func (fs *LoggingFs) Chmod(path string, mode uint32, context *fuse.Context) (code fuse.Status) {
	return fuse.ENOSYS
}

func (fs *LoggingFs) Chown(path string, uid uint32, gid uint32, context *fuse.Context) (code fuse.Status) {
	return fuse.ENOSYS
}

func (fs *LoggingFs) GetAttr(name string, context *fuse.Context) (a *fuse.Attr, code fuse.Status) {
	a = &fuse.Attr{}
	if name == "" {
		a.Mode = fuse.S_IFDIR | 0777
		return a, fuse.OK
	}

	a.Mode = fuse.S_IFREG | 0220
	return a, fuse.OK
}

func (fs *LoggingFs) Truncate(path string, offset uint64, context *fuse.Context) (code fuse.Status) {
	return fuse.OK
}

func (fs *LoggingFs) Utimens(path string, Atime *time.Time, Mtime *time.Time, context *fuse.Context) (code fuse.Status) {
	return fuse.ENOSYS
}

func (fs *LoggingFs) Readlink(name string, context *fuse.Context) (out string, code fuse.Status) {
	return "", fuse.ENOSYS
}

func (fs *LoggingFs) Mknod(name string, mode uint32, dev uint32, context *fuse.Context) (code fuse.Status) {
	return fuse.ENOSYS
}

func (fs *LoggingFs) Mkdir(path string, mode uint32, context *fuse.Context) (code fuse.Status) {
	return fuse.ENOSYS
}

func (fs *LoggingFs) Unlink(name string, context *fuse.Context) (code fuse.Status) {
	return fuse.ENOSYS
}

func (fs *LoggingFs) Rmdir(name string, context *fuse.Context) (code fuse.Status) {
	return fuse.ENOSYS
}

func (fs *LoggingFs) Symlink(pointedTo string, linkName string, context *fuse.Context) (code fuse.Status) {
	return fuse.ENOSYS
}

func (fs *LoggingFs) Rename(oldPath string, newPath string, context *fuse.Context) (code fuse.Status) {
	return fuse.ENOSYS
}

func (fs *LoggingFs) Link(orig string, newName string, context *fuse.Context) (code fuse.Status) {
	return fuse.ENOSYS
}

func (fs *LoggingFs) Access(name string, mode uint32, context *fuse.Context) (code fuse.Status) {
	return fuse.ENOSYS
}

func (fs *LoggingFs) Create(path string, flags uint32, mode uint32, context *fuse.Context) (fuseFile fuse.File, code fuse.Status) {
	return nil, fuse.ENOSYS
}
