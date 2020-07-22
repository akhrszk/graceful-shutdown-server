# graceful-shutdown-server

## 概要

接続中のConnectionが全て閉じるのを待ってからHttpServerをShutdownする

**Graceful Shutdown** については以下を参照

https://golang.org/pkg/net/http/#Server.Shutdown

`interrupt signal` を受け取ったら `http.Server` を `Shutdown` する

