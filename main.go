package main

import "log"
import "github.com/caser789/netstack/tcpip/link/tun"
import "github.com/caser789/netstack/tcpip/link/rawfile"

func main() {
    log.Println("start")
    name := "tun0"
    fd, err := tun.Open(name)
    if err != nil {
        log.Printf("error = %q", err)
    }

    log.Printf("opened fd is %d", fd)

    mtu, err := rawfile.GetMTU("tun0")
    if err != nil {
        log.Printf("error = %q", err)
    }
    log.Printf("mtu fd is %d", mtu)

    b := make([]byte, 20)
    rawfile.BlockingRead(fd, b)

    log.Println(b)

    rawfile.NonBlockingWrite(fd, b)
}
