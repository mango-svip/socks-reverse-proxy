package main

type Config struct {
    // 服务器地址
    ServerAddr string `json:"server_addr"`
    // 服务器端口
    ServerPort int `json:"server_port"`
    // 服务器名称
    ServerName string `json:"server_name"`
    // 服务器描述
    ServerDesc string `json:"server_desc"`
    // 服务器版本
    ServerVersion string `json:"server_version"`
    // 服务器类型
}

func NewConfig(serverAddr string, serverPort int, serverName string, serverDesc string, serverVersion string) *Config {
    return &Config{
        ServerAddr:    serverAddr,
        ServerPort:    serverPort,
        ServerName:    serverName,
        ServerDesc:    serverDesc,
        ServerVersion: serverVersion,
    }
}
