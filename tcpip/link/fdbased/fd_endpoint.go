package fdbased

import "syscall"
import "log"
import "github.com/caser789/netstack/tcpip/link/rawfile"
import "github.com/caser789/netstack/tcpip"
import "github.com/caser789/netstack/tcpip/header"

type dispatcher func(tcpip.NetworkProtocolNumber, []byte)


type endpoint struct {
    // fd is the file descriptor used to send and receive packets.
    fd int

    // mtu (maximum transimmision uint) is the maximum size of a packet.
    mtu int

    // closed is a function to be called when the FD's peer (if any) closes
    // its end of the communication pipe.
    closed func(error)
}

// New creates a new fd-based endpoint
func NewEndpoint(fd int, mtu int, closed func(error)) *endpoint {
    syscall.SetNonblock(fd, true)

    return &endpoint{
        fd: fd,
        mtu: mtu,
        closed: closed,
    }
}

// MTU implements stack.LinkEndpoint.MTU. It returns the value initialized
// during construction.
func (e *endpoint) MTU() uint32 {
    return uint32(e.mtu)
}

// MaxHeaderLength returns the maximum size of the header. Given that it
// doesn't have a header, it just returns 0.
func (e *endpoint) MaxHeaderLength() uint16 {
    return 0
}

// WritePacket writes outbound packets to the file descriptor. If it is not
// currently writable, the packet is dropped.
func (e *endpoint) WritePacket(hdr, payload []byte) error {
    if payload == nil {
        return rawfile.NonBlockingWrite(e.fd, hdr)
    }

    return rawfile.NonBlockingWrite2(e.fd, hdr, payload)
}

func (e *endpoint) Dispatch(deliver dispatcher, largeV []byte) (bool, error) {
    n, err := rawfile.BlockingRead(e.fd, largeV)
    if err != nil {
        return false, err
    }

    if n <= 0 {
        return false, err
    }

    v := make([]byte, len(largeV))
    copy(v, largeV)

    // We don't get any indication of what the packet is, so try to guess
    // if it's an IPv4 or IPv6 packet.
    var p tcpip.NetworkProtocolNumber
    switch header.IPVersion(v) {
    case header.IPv4Version:
        p = header.IPv4ProtocolNumber
    case header.IPv6Version:
        p = header.IPv6ProtocolNumber
    default:
        log.Printf("unknown protocol to dispatch %q", header.IPVersion(v))
        return true, nil
    }

    deliver(p, v)
    return true, nil
}

// dispatchLoop reads packets from the file descriptor in a loop and dispatches
// them to the network stack.
func (e *endpoint) dispatchLoop(d dispatcher) error {
	v := buffer.NewView(header.MaxIPPacketSize)
	for {
		cont, err := e.dispatch(d, v)
		if err != nil || !cont {
			if e.closed != nil {
				e.closed(err)
			}
			return err
		}
	}
}

// Attach launches the goroutine that reads packets from the file descriptor and
// dispatches them via the provided dispatcher.
func (e *endpoint) Attach(d dispatcher) {
    go e.dispatchLoop(d)
}
