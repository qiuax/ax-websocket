package conf

import (
	"flag"
	"fmt"
)

const (
	// 日志库
	MongoLogsDatabase = "ax-websocket"
)

// MongoAddr 通过参数的形式传进来
var MongoAddr string

func init() {
	addr := flag.String("mongo", "", "芒果数据库地址")
	flag.Parse()
	fmt.Println("数据库地址为-->" + *addr)
	MongoAddr = *addr
}
