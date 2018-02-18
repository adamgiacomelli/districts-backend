package main

import (
	"log"

	"github.com/adamgiacomelli/districts-backend/lib/wsServer"
)

func main() {
	//restServer.InitServer()
	log.Printf("Districts 0.1\n")

	wsServer.InitWsServer()
}
