package main

import (
    "fmt"
    "golang.org/x/net/proxy"
    "io"
    "log"
    "net"
)

func main() {

    // 读取 client.json 配置文件
    config := NewConfig("127.0.0.1", 8000, 9999, "127.0.0.1", "tcp")

    // 连接到socks5代理服务器
    dialer, err := proxy.SOCKS5(config.Proto, config.GetProxyServerAddr(), nil, proxy.Direct)

    if err != nil {
        log.Fatal(err)
    }

    locals := []LocalServer{
        {
            ServerAddr: "127.0.0.1",
            ServerPort: 9083,
            Target: TargetServer{
                ServerAddr: "127.0.0.1",
                ServerPort: 8083,
            },
        },
        {
            ServerAddr: "127.0.0.1",
            ServerPort: 9084,
            Target: TargetServer{
                ServerAddr: "127.0.0.1",
                ServerPort: 8084,
            },
        },
        {
            ServerAddr: "127.0.0.1",
            ServerPort: 9085,
            Target: TargetServer{
                ServerAddr: "127.0.0.1",
                ServerPort: 8085,
            },
        },
    }

    for _, local := range locals {
        listener, err := net.Listen(config.Proto, local.GetLocalAddr())
        if err != nil {
            log.Fatal(err)
        }
        go func(local LocalServer) {
            for {
                // 有新的连接进来时 接受连接
                fmt.Println("Waiting for connection...")
                conn, err := listener.Accept()
                if err != nil {
                    log.Fatal(err)
                }
                // 连接socks5服务器
                fmt.Println("Connected", conn)
                sock5_conn, err := dialer.Dial("tcp", local.Target.GetServerAddr())
                if err != nil {
                    log.Fatal(err)
                }
                go communicate(conn, sock5_conn)
            }
        }(local)
    }

    select {}

}

func communicate(conn net.Conn, sock5_conn net.Conn) {
    defer conn.Close()
    defer sock5_conn.Close()
    // 使用一个通道来通知协程何时完成任务
    done := make(chan struct{})
    done2 := make(chan struct{})
    // 连接成功后 开始转发数据
    // 读取conn的数据 写入sock5_conn
    fmt.Println("Start transfering...")
    // 读取conn的数据 写入sock5_conn
    go transferData(conn, sock5_conn, done)
    // 读取sock5_conn 的数据 写入conn
    go transferData(sock5_conn, conn, done2)

    select {
    case <-done:
        fmt.Println("Source Transfer finished")
    case <-done2:
        fmt.Println("Target Transfer finished")
    }
}

func transferData(source net.Conn, target net.Conn, chanDone chan<- struct{}) {
    defer close(chanDone)
    _, err := io.Copy(target, source)
    if err != nil {
        if err != io.EOF {
            // log.Printf("Error transferring data: %v", err)
        }
    }
}
