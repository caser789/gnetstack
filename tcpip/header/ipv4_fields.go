package header

import "github.com/caser789/netstack/tcpip"

const (
	versIHL  = 0
	tos      = 1
	totalLen = 2
	id       = 4
	flagsFO  = 6
	ttl      = 8
	protocol = 9
	checksum = 10
	srcAddr  = 12
	dstAddr  = 16
)

// IPv4Fields contains the fields of an IPv4 packet. It is used to describe the
// fields of a packet that needs to be encoded.
type IPv4Fields struct {
    // IHL is the "internet header length" field of an IPv4 packet.
    IHL uint8

    // TOS is the "type of service" field of an IPv4 packet.
    TOS uint8

    // TotalLength is the "total length" field of an IPv4 packet.
    TotalLength uint16

    // ID is the "identification" field of an IPv4 packet.
    ID uint16

    // Flags is the "flags" field of an IPv4 packet.
    Flags uint8

    // FragmentOffset is the "fragment offset" field of an IPv4 packet.
    FragmentOffset uint16

    // TTL is the "time to live" field of an IPv4 packet.
    TTL uint8

    // Protocol is the "protocol" field of an IPv4 packet.
    Protocol uint8

    // Checksum is the "checksum" field of an IPv4 packet.
    Checksum uint16

    // SrcAddr is the "source ip address" of an IPv4 packet.
    SrcAddr tcpip.Address

    // DstAddr is the "destination ip address" of an IPv4 packet.
    DstAddr tcpip.Address
}
