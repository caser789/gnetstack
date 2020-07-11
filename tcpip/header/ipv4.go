package header

import "github.com/caser789/netstack/tcpip"


// IPv4Version is the version of the ipv4 protocol.
const IPv4Version = 4

// IPv4ProtocolNumber is IPv4's network protocol number.
const IPv4ProtocolNumber tcpip.NetworkProtocolNumber = 0x0800

// IPv4MinimumSize is the minimum size of a valid IPv4 packet.
const IPv4MinimumSize = 20

// IPv4AddressSize is the size, in bytes, of an IPv4 address.
const IPv4AddressSize = 4

// IPv4 represents an ipv4 header stored in a byte array.
type IPv4 []byte

// Encode encodes all the fields of the ipv4 header.
func (b IPv4) Encode(i *IPv4Fields) {
    b[versIHL] = (4 << 4) | ((i.IHL/4) & 0xf)
    b.SetTOS(i.TOS, nil)
    b.SetTotalLength(i.TotalLength)
    binary.BigEndian.PutUint16(b[id:], i.ID)
    b.SetFlagsFragmentOffset(i.Flags, i.FragmentOffset)
    b[ttl] = i.TTL
    b[protocol] = i.Protocol
    b.SetChecksum(i.Checksum)
    b.SetSourceAddress(i.SrcAddr)
    b.SetDestinationAddress(i.DstAddr)
}

//////////////////////////////////////////////////
//
// Setter
//
//////////////////////////////////////////////////

// SetTotalLength sets the "total length" field of the ipv4 header.
func (b IPv4) SetTotalLength(totalLength uint16) {
    binary.BigEndian.PutUint16(b[totalLen:], totalLength)
}

// SetChecksum sets the checksum field of the ipv4 header.
func (b IPv4) SetChecksum(v uint16) {
    binary.BigEndian.PutUint16(b[checksum:], v)
}

// SetFlagsFragmentOffset sets the "flags" and "fragment offset" fields of the
// ipv4 header.
func (b IPv4) SetFlagsFragmentOffset(flags uint8, offset uint16) {
    v := (uint16(flags) << 13) | (offset >> 3)
    binary.BigEndian.PutUint16(b[flagsFO:], v)
}

// SetTOS sets the "type of service" field of the ipv4 header.
func (b IPv4) SetTOS(v uint8, _ uint32) {
    b[tos] = v
}

// SetSourceAddress sets the "source address" field of the ipv4 header.
func (b IPv4) SetSourceAddress(addr tcpip.Address) {
    copy(b[srcAddr:srcAddr+IPv4AddressSize], addr)
}

// SetDestinationAddress sets the "destination address" field of the ipv4
// header.
func (b IPv4) SetDestinationAddress(addr tcpip.Address) {
	copy(b[dstAddr:dstAddr+IPv4AddressSize], addr)
}

//////////////////////////////////////////////////
//
// Getter
//
//////////////////////////////////////////////////

// HeaderLength returns the value of the "header length" field of the ipv4
// header.
func (b IPv4) HeaderLength() uint8 {
	return (b[versIHL] & 0xf) * 4
}

// ID returns the value of the identifier field of the the ipv4 header.
func (b IPv4) ID() uint16 {
	return binary.BigEndian.Uint16(b[id:])
}

// Protocol returns the value of the protocol field of the the ipv4 header.
func (b IPv4) Protocol() uint8 {
	return b[protocol]
}

// Flags returns the "flags" field of the ipv4 header.
func (b IPv4) Flags() uint8 {
	return uint8(binary.BigEndian.Uint16(b[flagsFO:]) >> 13)
}

// TTL returns the "TTL" field of the ipv4 header.
func (b IPv4) TTL() uint8 {
	return b[ttl]
}

// FragmentOffset returns the "fragment offset" field of the ipv4 header.
func (b IPv4) FragmentOffset() uint16 {
	return binary.BigEndian.Uint16(b[flagsFO:]) << 3
}

// TotalLength returns the "total length" field of the ipv4 header.
func (b IPv4) TotalLength() uint16 {
	return binary.BigEndian.Uint16(b[totalLen:])
}

// Checksum returns the checksum field of the ipv4 header.
func (b IPv4) Checksum() uint16 {
	return binary.BigEndian.Uint16(b[checksum:])
}

// SourceAddress returns the "source address" field of the ipv4 header.
func (b IPv4) SourceAddress() tcpip.Address {
	return tcpip.Address(b[srcAddr : srcAddr+IPv4AddressSize])
}

// DestinationAddress returns the "destination address" field of the ipv4
// header.
func (b IPv4) DestinationAddress() tcpip.Address {
	return tcpip.Address(b[dstAddr : dstAddr+IPv4AddressSize])
}

// TOS returns the "type of service" field of the ipv4 header.
func (b IPv4) TOS() (uint8, uint32) {
	return b[tos], 0
}

//////////////////////////////////////////////////
//
// Other
//
//////////////////////////////////////////////////

// CalculateChecksum calculates the checksum of the ipv4 header.
func (b IPv4) CalculateChecksum() uint16 {
    return Checksum(b[:b.HeaderLength()], 0)
}
