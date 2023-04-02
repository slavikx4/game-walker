package main

import (
	"game-walker/src"
	"strings"
)

var players []*src.Player
var rooms src.Room

func initGame() {
	rooms = src.NewRoom()
}

func addPlayer(newPlayer *src.Player) {
	rooms.Mu.Lock()
	players = append(players, newPlayer)
	rooms.Mu.Unlock()
	go func(p *src.Player) {
		var command string
		for {
			select {
			case command = <-p.ChannelInput:
				answer := handleCommand(command, newPlayer)
				if answer != "" {
					p.HandleOutput(answer)
				}
			}
		}
	}(newPlayer)
	rooms.Mu.Lock()
	rooms.KitchenRoom.InRoom = append(rooms.KitchenRoom.InRoom, newPlayer)
	rooms.Mu.Unlock()
}

func handleCommand(command string, player *src.Player) string {
	commands := strings.Fields(command)
	f, ok := player.Do[commands[0]]
	if !ok {
		return "неизвестная команда"
	}
	answer := f(commands, player, &rooms)
	return answer
}

func main() {

}
