package main

import (
  "context"
  "os"
  "os/signal"
  "fmt"
  "log"
  "time"
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
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    log.Print("Shutting down...")
    if err := httpServer.Shutdown(ctx); err != nil {
      log.Fatal("Server forced to shutdown:", err)
    }
  }()

  http.HandleFunc("/", handler)
  httpServer.Addr = ":8888"
  httpServer.ListenAndServe()
}
