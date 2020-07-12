package stack

import "github.com/caser789/netstack/tcpip"
import "github.com/caser789/netstack/tcpip/buffer"

// LinkEndpoint is the interface implemented by data link layer protocols (e.g.,
// ethernet, loopback, raw) and used by network layer protocols to send packets
// out through the implementer's data link endpoint.
type LinkEndpoint interface {
    // MTU is the maximum transmission uint for this endpoint. This is
    // usually dictated by the backing physical network; when such a
    // physical network doesn't exist, the limit is generally 64k, which
    // includes the maximum size of an IP packet.
    MTU() uint32

    // MaxHeaderLength returns the maximum size the data link (and
    // lower level layers combined) headers can have. Higher levels use this
    // information to reserve space in the front of the packets they're
    // building.
    MaxHeaderLength() uint16

    // WritePacket writes a packet with the given protocol through the given
    // route.
    WritePacket(hdr *buffer.Prependable, payload buffer.View, protocol tcpip.NetworkProtocolNumber) error

    // Attach attaches the data link layer endpoint to the network-layer
    // dispatcher of the stack.
    Attach(dispatcher NetworkDispatcher)
}
