package main

import (
	"ax-websocket/lib"
	log "ax-websocket/modules/logger"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsHandle(w http.ResponseWriter, r *http.Request) {
	logFields := log.Fields{"read or write message": "message", "sub_category": "接收和发送消息 wsHandle()"}
	var wsConn *websocket.Conn

	wsConn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Errorf(logFields, "Upgrade 失败,error :%v", err)
		return
	}

	conn := lib.NewConn(wsConn)
	for {
		data, err := conn.ReadMessage()
		if err != nil {
			goto ERR
		}
		log.Infof(logFields, "接收的消息是 : %s", string(data))
		if err := conn.WriteMessage(data); err != nil {
			goto ERR
		}
	}

ERR:
	conn.Close()

}
func main() {
	http.HandleFunc("/ws", wsHandle)
	http.ListenAndServe(":80", nil)
}
