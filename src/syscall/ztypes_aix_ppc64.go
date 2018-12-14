// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs types_aix.go | go run mkpost.go

package syscall

const (
	sizeofPtr      = 0x8
	sizeofShort    = 0x2
	sizeofInt      = 0x4
	sizeofLong     = 0x8
	sizeofLongLong = 0x8
	PathMax        = 0x3ff
)

type (
	_C_short     int16
	_C_int       int32
	_C_long      int64
	_C_long_long int64
)

type Timespec struct {
	Sec  int64
	Nsec int64
}

type Timeval struct {
	Sec       int64
	Usec      int32
	Pad_cgo_0 [4]byte
}

type Timeval32 struct {
	Sec  int32
	Usec int32
}

type Timezone struct {
	Minuteswest int32
	Dsttime     int32
}

type Rusage struct {
	Utime    Timeval
	Stime    Timeval
	Maxrss   int64
	Ixrss    int64
	Idrss    int64
	Isrss    int64
	Minflt   int64
	Majflt   int64
	Nswap    int64
	Inblock  int64
	Oublock  int64
	Msgsnd   int64
	Msgrcv   int64
	Nsignals int64
	Nvcsw    int64
	Nivcsw   int64
}

type Rlimit struct {
	Cur uint64
	Max uint64
}

type _Pid_t int32

type _Gid_t uint32

type Flock_t struct {
	Type   int16
	Whence int16
	Sysid  uint32
	Pid    int32
	Vfs    int32
	Start  int64
	Len    int64
}

type Stat_t struct {
	Dev       uint64
	Ino       uint64
	Mode      uint32
	Nlink     int16
	Flag      uint16
	Uid       uint32
	Gid       uint32
	Rdev      uint64
	Ssize     int32
	Pad_cgo_0 [4]byte
	Atim      StTimespec_t
	Mtim      StTimespec_t
	Ctim      StTimespec_t
	Blksize   int64
	Blocks    int64
	Vfstype   int32
	Vfs       uint32
	Type      uint32
	Gen       uint32
	Reserved  [9]uint32
	Padto_ll  uint32
	Size      int64
}

type Statfs_t struct {
	Version   int32
	Type      int32
	Bsize     uint64
	Blocks    uint64
	Bfree     uint64
	Bavail    uint64
	Files     uint64
	Ffree     uint64
	Fsid      Fsid64_t
	Vfstype   int32
	Pad_cgo_0 [4]byte
	Fsize     uint64
	Vfsnumber int32
	Vfsoff    int32
	Vfslen    int32
	Vfsvers   int32
	Fname     [32]uint8
	Fpack     [32]uint8
	Name_max  int32
	Pad_cgo_1 [4]byte
}

type Fsid64_t struct {
	Val [2]uint64
}

type StTimespec_t struct {
	Sec       int64
	Nsec      int32
	Pad_cgo_0 [4]byte
}

type Dirent struct {
	Offset    uint64
	Ino       uint64
	Reclen    uint16
	Namlen    uint16
	Name      [256]uint8
	Pad_cgo_0 [4]byte
}

type RawSockaddrInet4 struct {
	Len    uint8
	Family uint8
	Port   uint16
	Addr   [4]byte /* in_addr */
	Zero   [8]uint8
}

type RawSockaddrInet6 struct {
	Len      uint8
	Family   uint8
	Port     uint16
	Flowinfo uint32
	Addr     [16]byte /* in6_addr */
	Scope_id uint32
}

type RawSockaddrUnix struct {
	Len    uint8
	Family uint8
	Path   [1023]uint8
}

type RawSockaddr struct {
	Len    uint8
	Family uint8
	Data   [14]uint8
}

type RawSockaddrAny struct {
	Addr RawSockaddr
	Pad  [1012]uint8
}

type _Socklen uint32

type Cmsghdr struct {
	Len   uint32
	Level int32
	Type  int32
}

type ICMPv6Filter struct {
	Filt [8]uint32
}

type Iovec struct {
	Base *byte
	Len  uint64
}

type IPMreq struct {
	Multiaddr [4]byte /* in_addr */
	Interface [4]byte /* in_addr */
}

type IPv6Mreq struct {
	Multiaddr [16]byte /* in6_addr */
	Interface uint32
}

type Linger struct {
	Onoff  int32
	Linger int32
}

type Msghdr struct {
	Name       *byte
	Namelen    uint32
	Pad_cgo_0  [4]byte
	Iov        *Iovec
	Iovlen     int32
	Pad_cgo_1  [4]byte
	Control    *byte
	Controllen uint32
	Flags      int32
}

const (
	SizeofSockaddrInet4 = 0x10
	SizeofSockaddrInet6 = 0x1c
	SizeofSockaddrAny   = 0x404
	SizeofSockaddrUnix  = 0x401
	SizeofLinger        = 0x8
	SizeofIPMreq        = 0x8
	SizeofIPv6Mreq      = 0x14
	SizeofMsghdr        = 0x30
	SizeofCmsghdr       = 0xc
	SizeofICMPv6Filter  = 0x20
)

const (
	PTRACE_TRACEME = 0x0
	PTRACE_CONT    = 0x7
	PTRACE_KILL    = 0x8
)

const (
	SizeofIfMsghdr = 0x10
)

type IfMsgHdr struct {
	Msglen    uint16
	Version   uint8
	Type      uint8
	Addrs     int32
	Flags     int32
	Index     uint16
	Addrlen   uint8
	Pad_cgo_0 [1]byte
}

type Utsname struct {
	Sysname  [32]uint8
	Nodename [32]uint8
	Release  [32]uint8
	Version  [32]uint8
	Machine  [32]uint8
}

const (
	_AT_FDCWD            = -0x2
	_AT_REMOVEDIR        = 0x1
	_AT_SYMLINK_NOFOLLOW = 0x1
)
