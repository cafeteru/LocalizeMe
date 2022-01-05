package main

import (
	_ "github.com/joho/godotenv/autoload"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/server"
	"log"
	"os"
	"os/signal"
)

func main() {
	serv, err := createServer()
	go serv.Start()
	waitInterrupt()
	shutdown(err, serv)
}

func createServer() (*server.Server, error) {
	port := os.Getenv("PORT")
	serv, err := server.Create(port)
	if err != nil {
		log.Fatal(err)
	}
	return serv, err
}

func waitInterrupt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func shutdown(err error, serv *server.Server) {
	err = serv.Close()
	if err != nil {
		return
	}
}
