package main

import "log"
import "syscall"
import "unsafe"

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

func main() {
    log.Println("start")
    name := "tun0"
    fd, err := Open(name)
    if err != nil {
        log.Printf("error = %q", err)
    }

    log.Printf("opened fd is %q", fd)

    mtu, err := GetMTU("tun0")
    if err != nil {
        log.Printf("error = %q", err)
    }
    log.Printf("mtu fd is %q", mtu)
}
