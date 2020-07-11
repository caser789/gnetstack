## test dispatch

##### code

```
package main

import "log"
import "github.com/caser789/netstack/tcpip/link/tun"
import "github.com/caser789/netstack/tcpip/link/rawfile"
import "github.com/caser789/netstack/tcpip/link/fdbased"
import "github.com/caser789/netstack/tcpip"

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

    b := make([]byte, 100)
    ep.Dispatch(deliver, b)
}

```

##### result

```
nc 192.168.1.2 1111 hello -q 10


2020/07/11 04:45:50 fdbased mtu is 1500
2020/07/11 04:45:50 fdbased MaxHeaderLength is 0
2020/07/11 04:45:55 in deviver 800 "E\x00\x00<rU@\x00@\x06E\x13\xc0\xa8\x01\x01\xc0\xa8\x01\x02\xdb\x00\x04W\xb4H@/\x00\x00\x00\x00\xa0\x02r\x10K\xe7\x00\x00\x02\x04\x05\xb4\x04\x02\b\n\x00l2z\x00\x00\x00\x00\x01\x03\x03\x06\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"
```
