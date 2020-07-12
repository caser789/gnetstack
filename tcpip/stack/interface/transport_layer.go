package stack

import "github.com/caser789/netstack/tcpip"
import "github.com/caser789/netstack/tcpip/buffer"

type TransportDispatcher interface {
    // DeliverTransportPacket delivers the packets to the appropriate
    // transport protocol endpoint.
    DeliverTransportPacket(protocol tcpip.TransportProtocolNumber, v buffer.View)
}
