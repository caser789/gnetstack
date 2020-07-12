package main

import "log"
import "github.com/caser789/netstack/tcpip/link/tun"
import "github.com/caser789/netstack/tcpip/buffer"
import "github.com/caser789/netstack/tcpip/link/rawfile"
import "github.com/caser789/netstack/tcpip/link/fdbased"
import "github.com/caser789/netstack/tcpip"
import "github.com/caser789/netstack/tcpip/network/ipv4"


func main() {
    tunName := "tun0"

	mtu, err := rawfile.GetMTU(tunName)
	if err != nil {
		log.Fatal(err)
	}

	fd, err := tun.Open(tunName)
	if err != nil {
		log.Fatal(err)
	}

    ep := fdbased.NewEndpoint(fd, mtu, nil)

    log.Printf("fdbased mtu is %d", ep.MTU())
    log.Printf("fdbased MaxHeaderLength is %d", ep.MaxHeaderLength())

    iep := ipv4.NewEndpoint("1.1.1.1", ep)

    log.Printf("ipv4 ep mtu is %d", iep.MTU())
    log.Printf("ipv4 ep MaxHeaderLength is %d", iep.MaxHeaderLength())

    hdr := buffer.NewPrependable(60)
    payload := make([]byte, 10)
    iep.WritePacket(&hdr, buffer.View(payload), tcpip.TransportProtocolNumber(1))

    v := make([]byte, 60)
    iep.HandlePacket(buffer.View(v))
}
