package mongodb

import (
	"log"

	"ax-websocket/conf"

	"gopkg.in/mgo.v2"
)

// 日志库
var SessionLogs *mgo.Session

func init() {
	session, err := mgo.Dial(conf.MongoAddr)
	if err != nil {
		log.Fatalf("连接日志数据库(mongodb)失败,err:%v", err)
		return
	}

	SessionLogs = session
}
