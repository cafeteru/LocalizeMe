package main

import (
	_ "github.com/joho/godotenv/autoload"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/server"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/tools"
	"log"
	"os"
	"os/signal"
)

func main() {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	serv := createServer()
	go serv.Start()
	waitInterrupt()
	shutdown(serv)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

func createServer() *server.Server {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	port := os.Getenv("PORT")
	serv := server.CreateServer(port)
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return serv
}

func waitInterrupt() {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

func shutdown(serv *server.Server) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	err := serv.Close()
	if err != nil {
		return
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
}
