package fdbased


type endpoint struct {
    // fd is the file descriptor used to send and receive packets.
    fd int

    // mtu (maximum transimmision uint) is the maximum size of a packet.
    mtu int

    // closed is a function to be called when the FD's peer (if any) closes
    // its end of the communication pipe.
    closed func(error)
}

// New creates a new fd-based endpoint. Register it to the stack

// MTU implements stack.LinkEndpoint.MTU. It returns the value initialized
// during construction.
func (e *endpoint) MTU() uint32 {
    return uint32(e.tmu)
}

// MaxHeaderLength returns the maximum size of the header. Given that it
// doesn't have a header, it just returns 0.
func (e *endpoint) MaxHeaderLength() uint16 {
    return 0
}

// WritePacket writes outbound packets to the file descriptor. If it is not
// currently writable, the packet is dropped.
func (e *endpoint) WritePacket(hdr, payload []byte) error {
    if payload == nil {
        return NonBlockingWrite(e.fd, hdr)
    }

    return NonBlockingWrite2(e.fd, hdr, payload)
}
