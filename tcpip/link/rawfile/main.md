```go
func main() {
    log.Println("start")
    name := "tun0"
    fd, err := tun.Open(name)
    if err != nil {
        log.Printf("error = %q", err)
    }

    log.Printf("opened fd is %d", fd)

    mtu, err := rawfile.GetMTU("tun0")
    if err != nil {
        log.Printf("error = %q", err)
    }
    log.Printf("mtu fd is %d", mtu)

    b := make([]byte, 20)
    rawfile.BlockingRead(fd, b)

    log.Println(b)

    rawfile.NonBlockingWrite(fd, b)
}
```
