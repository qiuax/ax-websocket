# ax-websocket
使用golang编写的websocket消息发送接口,并用mongodb存放日志

环境:
  事先要准备golang,mongodb.
启动:
    1.编译:进入ax-websocket,命令启动,go build socketserver.go
    2.运行: go run socketserver.go -mongo="***:27017"
    3.打开static的html页面,测试即可
 
