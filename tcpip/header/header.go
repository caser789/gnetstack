package header

const versIHL = 0


// IPVersion returns the version of IP used in the given packet. It returns -1
// if the packet is not large enough to contain the version field
func IPVersion(b []byte) int {
    // Length must be at least offset+length of version field.
    if len(b) < versIHL+1 {
        return -1
    }

    return int(b[versIHL] >> 4)
}
