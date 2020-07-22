# 接続中のConnectionが全て閉じるのを待ってからHttpServerをShutdownする

**Graceful Shutdown** については以下を参照

https://golang.org/pkg/net/http/#Server.Shutdown

今回の実装では、 `interrupt signal` を受け取ったら `http.Server` を `Shutdown` する
