package main

import (
	slog "github.com/go-eden/slf4go"
	_ "github.com/joho/godotenv/autoload"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/server"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/tools"
	"os"
	"os/signal"
)

func main() {
	slog.Debugf("%s: start", tools.GetCurrentFuncName())
	serv := createServer()
	go serv.Start()
	waitInterrupt()
	shutdown(serv)
	slog.Debugf("%s: end", tools.GetCurrentFuncName())
}

func createServer() *server.Server {
	slog.Debugf("%s: start", tools.GetCurrentFuncName())
	port := os.Getenv("PORT")
	serv := server.CreateServer(port)
	slog.Debugf("%s: end", tools.GetCurrentFuncName())
	return serv
}

func waitInterrupt() {
	slog.Debugf("%s: start", tools.GetCurrentFuncName())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	slog.Debugf("%s: end", tools.GetCurrentFuncName())
}

func shutdown(serv *server.Server) {
	slog.Debugf("%s: start", tools.GetCurrentFuncName())
	err := serv.Close()
	if err != nil {
		return
	}
	slog.Debugf("%s: end", tools.GetCurrentFuncName())
}
