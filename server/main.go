package main

import (
    "context"
    "fmt"
    "github.com/armon/go-socks5"
    "net"
)

func main() {

    // Create a SOCKS5 server
    conf := &socks5.Config{
        Dial: func(ctx context.Context, network, addr string) (net.Conn, error) {
            fmt.Println(network)
            fmt.Println(addr)
            return net.Dial(network, addr)
        },
    }
    server, err := socks5.New(conf)
    if err != nil {
        panic(err)
    }

    // Create SOCKS5 proxy on localhost port 8000
    if err := server.ListenAndServe("tcp", "127.0.0.1:8000"); err != nil {
        panic(err)
    }
}
