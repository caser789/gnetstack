// Package ipv4 contains the implementation of the ipv4 network protocol. To use
// it in the networking stack, this package must be added to the project, and
// activated on the stack by passing ipv4.ProtocolName (or "ipv4") as one of the
// network protocols when calling stack.New(). Then endpoints can be created
// by passing ipv4.ProtocolNumber as the network protocol number when calling
// Stack.NewEndpoint().
package ipv4

import "github.com/caser789/netstack/tcpip"
import "github.com/caser789/netstack/tcpip/stack"
import "github.com/caser789/netstack/tcpip/buffer"
import "github.com/caser789/netstack/tcpip/header"

// maxTotalSize is maximum size that can be encoded in the 16-bit
// TotalLength field of the ipv4 header.
const maxTotalSize = 0xffff

// ProtocolNumber is the ipv4 protocol number
const ProtocolNumber = header.IPv4ProtocolNumber

type address [header.IPv4AddressSize]byte

type endpoint struct {
    linkEP stack.LinkEndpoint
    address address
}

func NewEndpoint(addr tcpip.Address, linkEP stack.LinkEndpoint) *endpoint{
    e := &endpoint{
        linkEP: linkEP,
    }
    copy(e.address[:], addr)

    return e
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

func (e *endpoint) WritePacket(hdr *buffer.Prependable, payload buffer.View, protocol tcpip.TransportProtocolNumber) error {
    ip := header.IPv4(hdr.Prepend(header.IPv4MinimumSize))
    length := uint16(hdr.UsedLength() + len(payload))
    id := uint32(0)
    if length > header.IPv4MaximumHeaderSize + 8 {
        // Packet of 68 bytes or less are required by RFC 791 to not be
        // fragmented, so we only assign ids to larger packets.
    }
    ip.Encode(&header.IPv4Fields{
        IHL: header.IPv4MinimumSize,
        TotalLength: length,
        ID: uint16(id),
        TTL: 65,
        Protocol: uint8(protocol),
        SrcAddr: tcpip.Address(e.address[:]),
        // DstAddr: r.RemoteAddress,
        DstAddr: nil,
    })
    ip.SetChecksum(^ip.CalculateChecksum())

    return e.linkEP.WritePacket(hdr, payload, ProtocolNumber)
}
