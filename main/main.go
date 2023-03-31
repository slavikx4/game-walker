package main

import (
	"game-walker/src"
	"strings"
)

var players []src.Player
var index = -1

func initGame() {
	player := src.Player{
		Room:        src.NewRoom(),
		CurrentRoom: "кухня",
		Inventory:   []string{},
		Do: map[string]func([]string, *src.Player) string{
			"осмотреться": src.Look,
			"идти":        src.Walk,
			"одеть":       src.Dress,
			"взять":       src.Take,
			"применить":   src.Apply,
		},
	}
	index++
	players = append(players, player)
}

func handleCommand(command string) string {
	commands := strings.Fields(command)
	player := &players[index]
	f, ok := player.Do[commands[0]]
	if !ok {
		return "неизвестная команда"
	}
	answer := f(commands, player)
	return answer
}

func main() {

}
