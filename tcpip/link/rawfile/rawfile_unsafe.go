package main

import "fmt"
import "syscall"
import "unsafe"
import "math"


func GetMTU(name string) (int, error) {
    fd, err := syscall.Socket(syscall.AF_UNIX, syscall.SOCK_DGRAM, 0)
    if err != nil {
        return 0, err
    }
    defer syscall.Close(fd)

    var ifreq struct {
        name [16]byte
        mtu int32
        _ [20]byte
    }
    copy(ifreq.name[:], name)
    _, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), syscall.SIOCGIFMTU, uintptr(unsafe.Pointer(&ifreq)))
    if errno != 0 {
        return 0, errno
    }

    return int(ifreq.mtu), nil
}

// BlockingRead reads from a file descriptor that is set up as non-blocking. If
// no data is available, it will block in a poll() syscall until the file
// descriptor becomes readable.
func BlockingRead(fd int, b []byte) (int, error) {
	for {
		n, _, e := syscall.RawSyscall(syscall.SYS_READ, uintptr(fd), uintptr(unsafe.Pointer(&b[0])), uintptr(len(b)))
		if e == 0 {
			return int(n), nil
		}

		event := struct {
			fd      int32
			events  int16
			revents int16
		}{
			fd:     int32(fd),
			events: 1, // POLLIN
		}

		_, _, e = syscall.Syscall(syscall.SYS_POLL, uintptr(unsafe.Pointer(&event)), 1, uintptr(math.MaxUint64))
		if e != 0 && e != syscall.EINTR {
			return 0, e
		}
	}
}

func NonBlockingWrite(fd int, buf []byte) error {
	var ptr unsafe.Pointer
	if len(buf) > 0 {
		ptr = unsafe.Pointer(&buf[0])
	}

	n, _, e := syscall.RawSyscall(syscall.SYS_WRITE, uintptr(fd), uintptr(ptr), uintptr(len(buf)))
	if e != 0 {
		return e
	}

	if n != uintptr(len(buf)) {
		return fmt.Errorf("wrong number of bytes written: expected %d, got %d", len(buf), n)
	}

	return nil
}

// NonBlockingWrite2 writes up to two bytes slices to a file descriptor in a
// single syscall. It fails if partial data is written
func NonBlockingWrite2(fd int, b1, b2 []byte) error {
    // If there is no second buffer, issue a regular write.
    if len(b2) == 0 {
        return NonBlockingWrite(fd, b1)
    }

    // We have two buffers. Build the iovec that represents them and issue
    // a writev syscall.
    iovec := [...]syscall.Iovec{
        {
            Base: (*byte)(unsafe.Pointer(&b1[0])),
            Len: uint64(len(b1)),
        },
        {
            Base: (*byte)(unsafe.Pointer(&b2[0])),
            Len: uint64(len(b2)),
        },
    }

    n, _, e := syscall.RawSyscall(syscall.SYS_WRITEV, uintptr(fd), uintptr(unsafe.Pointer(&iovec[0])), 2)
    if e != 0 {
        return e
    }

    if n != uintptr(len(b1)+len(b2)) {
        return fmt.Errorf("wrong number of bytes written: expected: %d, got %d", len(b1)+len(b2), n)
    }

    return nil
}
