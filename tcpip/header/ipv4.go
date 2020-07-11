package header

import "github.com/caser789/netstack/tcpip"


// IPv4Version is the version of the ipv4 protocol.
const IPv4Version = 4

// IPv4ProtocolNumber is IPv4's network protocol number.
const IPv4ProtocolNumber tcpip.NetworkProtocolNumber = 0x0800

// IPv4MinimumSize is the minimum size of a valid IPv4 packet.
const IPv4MinimumSize = 20
