package main

import "fmt"

type Config struct {
    Proxy ProxyServer `json:"proxy"`

    Local LocalServer `json:"local"`
    // 监听协议
    Proto string `json:"proto"`
}

type ProxyServer struct {
    // 服务器地址
    ServerAddr string `json:"server_addr"`
    // 服务器端口
    ServerPort int `json:"server_port"`
}

type LocalServer struct {
    // 服务器地址
    ServerAddr string `json:"server_addr"`
    // 服务器端口
    ServerPort int `json:"server_port"`

    Target TargetServer `json:"target"`
}

type TargetServer struct {
    // 目标地址
    ServerAddr string `json:"server_addr"`
    // 目标端口
    ServerPort int `json:"server_port"`
}

func NewConfig(serverAddr string, serverPort int, localPort int, localAddr string, Proto string) *Config {

    return &Config{
        ProxyServer{
            ServerAddr: serverAddr,
            ServerPort: serverPort,
        },
        LocalServer{
            ServerAddr: localAddr,
            ServerPort: localPort,
        },
        Proto,
    }
}

func (c *Config) GetProxyServerAddr() string {
    return fmt.Sprintf("%s:%d", c.Proxy.ServerAddr, c.Proxy.ServerPort)
}

func (c *Config) GetLocalAddr() string {
    return fmt.Sprintf("%s:%d", c.Local.ServerAddr, c.Local.ServerPort)
}

func (local *LocalServer) GetLocalAddr() string {
    return fmt.Sprintf("%s:%d", local.ServerAddr, local.ServerPort)
}

func (t *TargetServer) GetServerAddr() string {
    return fmt.Sprintf("%s:%d", t.ServerAddr, t.ServerPort)
}
