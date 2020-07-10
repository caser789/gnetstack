package main

import "log"
import "time"
import "fmt"
import "syscall"
import "unsafe"
import "math"

// Open opens the specified TUN device, sets it to non-blocking mode, and
// returns its file descriptor.
func Open(name string) (int, error) {
    fd, err := syscall.Open("/dev/net/tun", syscall.O_RDWR, 0)
    if err != nil {
        return -1, err
    }

    var ifr struct {
        name [16]byte
        flags uint16
        _ [22]byte
    }

    copy(ifr.name[:], name)
    ifr.flags = syscall.IFF_TUN | syscall.IFF_NO_PI
    _, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), syscall.TUNSETIFF, uintptr(unsafe.Pointer(&ifr)))
    if errno != 0 {
        syscall.Close(fd)
        return -1, errno
    }

    if err = syscall.SetNonblock(fd, true); err != nil {
        syscall.Close(fd)
        return -1, err
    }

    return fd, nil
}

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

func main() {
    log.Println("start")
    name := "tun0"
    fd, err := Open(name)
    if err != nil {
        log.Printf("error = %q", err)
    }

    log.Printf("opened fd is %d", fd)

    mtu, err := GetMTU("tun0")
    if err != nil {
        log.Printf("error = %q", err)
    }
    log.Printf("mtu fd is %d", mtu)

    b := make([]byte, 20)
    go BlockingRead(fd, b)
    go NonBlockingWrite(fd, []byte{'h', 'e', 'l', 'l', 'o'})

    time.Sleep(10 * 1000 * time.Millisecond)
    log.Println(b)
}
