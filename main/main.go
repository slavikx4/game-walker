package main

import (
	"game-walker/src"
	"game-walker/telegram"
	"log"
)

func main() {

	src.StartWorld()
	log.Println("initGame: OK")

	log.Println("Start work bot")
	telegram.StartBot()
}
