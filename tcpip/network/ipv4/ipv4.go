// Package ipv4 contains the implementation of the ipv4 network protocol. To use
// it in the networking stack, this package must be added to the project, and
// activated on the stack by passing ipv4.ProtocolName (or "ipv4") as one of the
// network protocols when calling stack.New(). Then endpoints can be created
// by passing ipv4.ProtocolNumber as the network protocol number when calling
// Stack.NewEndpoint().
package ipv4

import "github.com/caser789/netstack/tcpip/stack"

// maxTotalSize is maximum size that can be encoded in the 16-bit
// TotalLength field of the ipv4 header.
const maxTotalSize = 0xffff

type endpoint struct {
    linkEP stack.LinkEndpoint
}

func NewEndpoint(linkEP stack.LinkEndpoint) *endpoint{
    return &endpoint{
        linkEP: linkEP,
    }
}

// MaxHeaderLength returns the maximum length needed by ipv4 headers (and
// underlying protocols).
func (e *endpoint) MaxHeaderLength() uint16 {
    return e.linkEP.MaxHeaderLength() + header.IPv4MinimumSize
}

// MTU implements stack.NetworkEndpoint.MTU. It returns the link-layer MTU minus
// the network layer max header length.
func (e *endpoint) MTU() uint32 {
    lmtu := e.linkEP.MTU()
    if lmtu > maxTotalSize {
        lmtu = maxTotalSize
    }

    return lmtu - uint32(e.MaxHeaderLength())
}
