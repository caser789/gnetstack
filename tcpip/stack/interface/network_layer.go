package stack

import "github.com/caser789/netstack/tcpip"

// NetworkDispatcher contains the methods used by the network stack to deliver
// packets to the appropriate network endpoint after it has been handled by
// the data link layer
type NetworkDispatcher interface {
    // DeliverNetworkPacket finds the appropriate network protocol
    // endpoint and hands the packet over for further processing.
    DeliverNetworkPacket(protocol tcpip.NetworkProtocolNumber, v []byte)
}

// NetworkEndpoint is the interface that needs to be implemented by endpoints
// of network layer protocols (e.g., ipv4, ipv6).
type NetworkEndpoint interface {
	// MTU is the maximum transmission unit for this endpoint. This is
	// generally calculated as the MTU of the underlying data link endpoint
	// minus the network endpoint max header length.
	MTU() uint32

	// MaxHeaderLength returns the maximum size the network (and lower
	// level layers combined) headers can have. Higher levels use this
	// information to reserve space in the front of the packets they're
	// building.
	MaxHeaderLength() uint16

	// WritePacket writes a packet to the given destination address and
	// protocol.
	WritePacket(hdr *buffer.Prependable, payload buffer.View, protocol tcpip.TransportProtocolNumber) error

	// HandlePacket is called by the link layer when new packets arrive to
	// this network endpoint.
	HandlePacket(v buffer.View)
}

