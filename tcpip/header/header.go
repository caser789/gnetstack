package header

const versIHL = 0

// MaxIPPacketSize is the maximum supported IP packet size, exluding
// jumbograms. The maximum IPv4 packet size is 64k-1 (total size must fit
// in 16 bits). For IPv6, the payload max size (excluding jumbograms) is
// 64k-1 (also needs to fit in 16 bits). So we use 64k - 1 + 2 * m, where
// m is the minimum IPv6 header size; we leave room for some potential
// IP options.
const MaxIPPacketSize = 0xffff + 2*IPv6MinimumSize


// IPVersion returns the version of IP used in the given packet. It returns -1
// if the packet is not large enough to contain the version field
func IPVersion(b []byte) int {
    // Length must be at least offset+length of version field.
    if len(b) < versIHL+1 {
        return -1
    }

    return int(b[versIHL] >> 4)
}
