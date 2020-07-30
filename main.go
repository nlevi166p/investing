package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	mongoClient := newMongoDBClient()
	httpServer := newServer(mongoClient)
	httpServer.startServer()
	<-sigs
}
