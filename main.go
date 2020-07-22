package main

import (
  "context"
  "os"
  "os/signal"
  "fmt"
  "log"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello")
}

func main() {
  var httpServer http.Server

  quit := make(chan os.Signal)
  signal.Notify(quit, os.Interrupt)

  go func() {
    log.Print("HttpServer Up Success!")

    <-quit
    ctx := context.Background()
    log.Print("Shutting down...")
    httpServer.Shutdown(ctx)
  }()

  http.HandleFunc("/", handler)
  httpServer.Addr = ":8888"
  httpServer.ListenAndServe()
}
