package helper

import (
	"log"
	"net"
)

// 本机ip地址
var LocalIp string = ""

func init() {
	LocalIp = GetLocalIp()
}

// 获取本机ip地址
func GetLocalIp() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		log.Printf("获取本机ip地址出错: %s\n", err)
		return ""
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String() // 只取第1个
			}
		}
	}

	return ""
}
