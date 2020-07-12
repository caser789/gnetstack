## test MTU() and MaxHeaderLength()

```go
package main

import "log"
import "github.com/caser789/netstack/tcpip/link/tun"
import "github.com/caser789/netstack/tcpip/link/rawfile"
import "github.com/caser789/netstack/tcpip/link/fdbased"
import "github.com/caser789/netstack/tcpip"
import "github.com/caser789/netstack/tcpip/network/ipv4"

func deliver(p tcpip.NetworkProtocolNumber, b []byte) {
    log.Printf("in deviver %d %q", p, b)
}

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
}
```

```
2020/07/12 04:27:28 fdbased mtu is 1500
2020/07/12 04:27:28 fdbased MaxHeaderLength is 0
2020/07/12 04:27:28 ipv4 ep mtu is 1480
2020/07/12 04:27:28 ipv4 ep MaxHeaderLength is 20
```
