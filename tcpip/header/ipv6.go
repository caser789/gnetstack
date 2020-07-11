package header

import "github.com/caser789/netstack/tcpip"


// IPv6Version is the version of the ipv6 protocol.
const IPv6Version = 6

// IPv6ProtocolNumber is IPv6's network protocol number.
const IPv6ProtocolNumber tcpip.NetworkProtocolNumber = 0x86dd

// IPv6MinimumSize is the minimum size of a valid IPv6 packet.
const IPv6MinimumSize = 40
