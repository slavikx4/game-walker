package src

import "strings"

type Player struct {
	Name          string
	CurrentRoom   string
	Inventory     []string
	NeedToDo      []string
	ChannelOutput chan string
	ChannelInput  chan string
	Do            map[string]func([]string, *Player, *Room) string
}

func NewPlayer(name string) *Player {
	player := Player{
		Name:          name,
		CurrentRoom:   "кухня",
		Inventory:     []string{},
		NeedToDo:      []string{"собрать рюкзак", "идти в универ"},
		ChannelOutput: make(chan string),
		ChannelInput:  make(chan string),
		Do: map[string]func([]string, *Player, *Room) string{
			"осмотреться":    Look,
			"идти":           Walk,
			"одеть":          Dress,
			"взять":          Take,
			"применить":      Apply,
			"сказать":        Speak,
			"сказать_игроку": SpeackTo,
		},
	}
	return &player
}
func (p *Player) GetOutput() chan string {
	return p.ChannelOutput
}

func (p *Player) HandleInput(command string) {
	p.ChannelInput <- command
}

func (p *Player) HandleOutput(answer string) {
	p.ChannelOutput <- answer
}

func Look(commands []string, player *Player, rooms *Room) string {
	var answer string
	switch player.CurrentRoom {
	case "кухня":
		answer += "ты находишься на кухне, "
		if len(rooms.KitchenRoom.Table) > 0 {
			answer += "на столе "
			for _, el := range rooms.KitchenRoom.Table {
				answer += el + ", "
			}
		}
		if len(player.NeedToDo) > 0 {
			answer += "надо "
			for i, el := range player.NeedToDo {
				if i != len(player.NeedToDo)-1 {
					answer += el + " и "
				} else {
					answer += el + ". "
				}
			}
		}
		answer += "можно пройти - "
		for i, el := range rooms.KitchenRoom.Entrance {
			if i != len(rooms.KitchenRoom.Entrance)-1 {
				answer += el + ", "
			} else {
				answer += el
			}
		}
		if len(rooms.KitchenRoom.InRoom) > 1 {
			answer += ". Кроме вас тут ещё "
			for _, el := range rooms.KitchenRoom.InRoom {
				if el.Name == player.Name {
					continue
				}
				answer += el.Name + ", "
			}
			answer = strings.TrimSuffix(answer, ", ")
		}
	case "комната":
		if len(rooms.BedRoom.Table)+len(rooms.BedRoom.Chain) == 0 {
			answer += "пустая комната. "
		} else {

			if len(rooms.BedRoom.Table) > 0 {
				answer += "на столе: "
				for _, el := range rooms.BedRoom.Table {
					answer += el + ", "
				}
			}
			if len(rooms.BedRoom.Chain) > 0 {
				answer += "на стуле - "
				for i, el := range rooms.BedRoom.Chain {
					if i != len(rooms.BedRoom.Chain)-1 {
						answer += el + ", "
					} else {
						answer += el + ". "
					}
				}
			} else {
				answer = strings.TrimSuffix(answer, ", ")
				answer += ". "
			}
		}
		answer += "можно пройти - "
		for i, el := range rooms.BedRoom.Entrance {
			if i != len(rooms.BedRoom.Entrance)-1 {
				answer += el + ", "
			} else {
				answer += el
			}
		}
	case "коридор":
		answer += "ничего интересного. "
		for i, el := range rooms.LineRoom.Entrance {
			if i != len(rooms.LineRoom.Entrance)-1 {
				answer += el + ", "
			} else {
				answer += el
			}
		}
	case "домой":
		answer += "ничего интересного. "
		for i, el := range rooms.HomeRoom.Entrance {
			if i != len(rooms.HomeRoom.Entrance)-1 {
				answer += el + ", "
			} else {
				answer += el
			}
		}
	case "улица":
		answer += "ничего интересного. "
		for i, el := range rooms.StreetRoom.Entrance {
			if i != len(rooms.StreetRoom.Entrance)-1 {
				answer += el + ", "
			} else {
				answer += el
			}
		}
	}
	return answer
}
func Walk(commands []string, player *Player, rooms *Room) string {
	gone(player, rooms)
	var answer string
	switch commands[1] {
	case "кухня":
		if search(player.CurrentRoom, rooms.KitchenRoom.Entrance) {
			rooms.KitchenRoom.InRoom = append(rooms.KitchenRoom.InRoom, player)
			player.CurrentRoom = "кухня"
			answer += "кухня, ничего интересного. "
			answer += "можно пройти - "
			for i, el := range rooms.KitchenRoom.Entrance {
				if i != len(rooms.KitchenRoom.Entrance)-1 {
					answer += el + ", "
				} else {
					answer += el
				}
			}
		} else {
			answer = "нет пути в кухня"
		}
	case "комната":
		if search(player.CurrentRoom, rooms.BedRoom.Entrance) {
			rooms.BedRoom.InRoom = append(rooms.BedRoom.InRoom, player)
			player.CurrentRoom = "комната"
			answer += "ты в своей комнате. "
			answer += "можно пройти - "
			for i, el := range rooms.BedRoom.Entrance {
				if i != len(rooms.BedRoom.Entrance)-1 {
					answer += el + ", "
				} else {
					answer += el
				}
			}
		} else {
			answer = "нет пути в комната"
		}
	case "коридор":
		if search(player.CurrentRoom, rooms.LineRoom.Entrance) {
			rooms.LineRoom.InRoom = append(rooms.LineRoom.InRoom, player)
			player.CurrentRoom = "коридор"
			answer += "ничего интересного. "
			answer += "можно пройти - "
			for i, el := range rooms.LineRoom.Entrance {
				if i != len(rooms.LineRoom.Entrance)-1 {
					answer += el + ", "
				} else {
					answer += el
				}
			}
		} else {
			answer = "нет пути в коридор"
		}
	case "домой":
		if search(player.CurrentRoom, rooms.HomeRoom.Entrance) {
			rooms.HomeRoom.InRoom = append(rooms.HomeRoom.InRoom, player)
			player.CurrentRoom = "домой"
			answer += "ничего интересного. "
			answer += "можно пройти - "
			for i, el := range rooms.HomeRoom.Entrance {
				if i != len(rooms.HomeRoom.Entrance)-1 {
					answer += el + ", "
				} else {
					answer += el
				}
			}
		} else {
			answer = "нет пути в домой"
		}
	case "улица":
		if rooms.StreetRoom.Door {
			if search(player.CurrentRoom, rooms.StreetRoom.Entrance) {
				rooms.StreetRoom.InRoom = append(rooms.StreetRoom.InRoom, player)
				player.CurrentRoom = "улица"
				answer += "на улице весна. "
				answer += "можно пройти - домой"
			} else {
				answer = "нет пути в улица"
			}
		} else {
			answer = "дверь закрыта"
		}
	}
	return answer
}
func Dress(commands []string, player *Player, rooms *Room) string {
	var answer string
	switch commands[1] {
	case "рюкзак":
		if search("рюкзак", rooms.BedRoom.Chain) {
			player.Inventory = append(player.Inventory, "рюкзак")
			removeItem("рюкзак", &rooms.BedRoom.Chain)
			answer = "вы одели: рюкзак"
		} else {
			answer = "нет такого"
		}
	}
	return answer
}
func Take(commands []string, player *Player, rooms *Room) string {
	var answer string
	if search("рюкзак", player.Inventory) {
		switch commands[1] {
		case "ключи":
			if search("ключи", rooms.BedRoom.Table) {
				player.Inventory = append(player.Inventory, "ключи")
				removeItem("ключи", &rooms.BedRoom.Table)
				answer = "предмет добавлен в инвентарь: ключи"
			} else {
				answer = "нет такого"
			}
		case "конспекты":
			if search("конспекты", rooms.BedRoom.Table) {
				player.Inventory = append(player.Inventory, "конспекты")
				removeItem("конспекты", &rooms.BedRoom.Table)
				removeItem("собрать рюкзак", &player.NeedToDo)
				answer = "предмет добавлен в инвентарь: конспекты"
			} else {
				answer = "нет такого"
			}
		default:
			answer = "нет такого"
		}
	} else {
		answer = "некуда класть"
	}
	return answer
}
func Apply(commands []string, player *Player, rooms *Room) string {
	var answer string
	if search(commands[1], player.Inventory) {
		switch commands[2] {
		case "дверь":
			if player.CurrentRoom == "коридор" {
				rooms.StreetRoom.Door = true
				answer = "дверь открыта"
			} else {
				answer = "не к чему применить"
			}
		default:
			answer = "не к чему применить"
		}
	} else {
		answer = "нет предмета в инвентаре - " + commands[1]
	}
	return answer
}

func Speak(commands []string, player *Player, rooms *Room) string {
	var answer string
	answer = player.Name + " говорит: "
	for i, el := range commands {
		if i == 0 {
			continue
		}
		answer += el + " "
	}
	answer = strings.TrimSuffix(answer, " ")
	switch player.CurrentRoom {
	case "кухня":
		for _, el := range rooms.KitchenRoom.InRoom {
			if el.Name != player.Name {
				el.HandleOutput(answer)
			}
		}
	case "комната":
		for _, el := range rooms.BedRoom.InRoom {
			if el.Name != player.Name {
				el.HandleOutput(answer)
			}
		}
	case "коридор":
		for _, el := range rooms.LineRoom.InRoom {
			if el.Name != player.Name {
				el.HandleOutput(answer)
			}
		}
	case "домой":
		for _, el := range rooms.HomeRoom.InRoom {
			if el.Name != player.Name {
				el.HandleOutput(answer)
			}
		}
	case "улица":
		for _, el := range rooms.StreetRoom.InRoom {
			if el.Name != player.Name {
				el.HandleOutput(answer)
			}
		}
	}
	return answer
}

func SpeackTo(commands []string, player *Player, rooms *Room) string {
	var answer string
	switch player.CurrentRoom {
	case "кухня":
		if searchPlayer(commands[1], rooms.KitchenRoom.InRoom) {
			if len(commands) > 2 {
				answer = player.Name + " говорит вам: "
				for i, el := range commands {
					if i != 0 && i != 1 {
						answer += el + " "
					}
				}
				answer = strings.TrimSuffix(answer, " ")
				for _, el := range rooms.KitchenRoom.InRoom {
					if commands[1] == el.Name {
						el.HandleOutput(answer)
					}
				}
			} else {
				answer = player.Name + " выразительно молчит, смотря на вас"
				for _, el := range rooms.KitchenRoom.InRoom {
					if commands[1] == el.Name {
						el.HandleOutput(answer)
					}
				}
			}
		} else {
			answer = "тут нет такого игрока"
			player.HandleOutput(answer)
		}
	}
	return ""
}

func searchPlayer(name string, m []*Player) bool {
	for _, el := range m {
		if name == el.Name {
			return true
		}
	}
	return false
}

func search(item string, m []string) bool {
	for _, el := range m {
		if item == el {
			return true
		}
	}
	return false
}

func removeItem(item string, slice *[]string) {
	temp := make([]string, len(*slice)-1)
	var indexTemp int
	for i := 0; i < len(*slice); i++ {
		if (*slice)[i] != item {
			temp[indexTemp] = (*slice)[i]
			indexTemp++
		}
	}
	*slice = temp
}

func gone(player *Player, rooms *Room) {
	switch player.CurrentRoom {
	case "кухня":
		for i, el := range rooms.KitchenRoom.InRoom {
			if el.Name == player.Name {
				if i != len(rooms.KitchenRoom.InRoom)-1 {
					rooms.KitchenRoom.InRoom = append(rooms.KitchenRoom.InRoom[:i], rooms.KitchenRoom.InRoom[i+1:]...)
				} else {
					rooms.KitchenRoom.InRoom = rooms.KitchenRoom.InRoom[:len(rooms.KitchenRoom.InRoom)-1]
				}
			}
		}
	case "комната":
		for i, el := range rooms.BedRoom.InRoom {
			if el.Name == player.Name {
				if i != len(rooms.BedRoom.InRoom)-1 {
					rooms.BedRoom.InRoom = append(rooms.BedRoom.InRoom[:i], rooms.BedRoom.InRoom[i+1:]...)
				} else {
					rooms.BedRoom.InRoom = rooms.BedRoom.InRoom[:len(rooms.BedRoom.InRoom)-1]
				}
			}
		}
	case "коридор":
		for i, el := range rooms.LineRoom.InRoom {
			if el.Name == player.Name {
				if i != len(rooms.LineRoom.InRoom)-1 {
					rooms.LineRoom.InRoom = append(rooms.LineRoom.InRoom[:i], rooms.LineRoom.InRoom[i+1:]...)
				} else {
					rooms.LineRoom.InRoom = rooms.LineRoom.InRoom[:len(rooms.LineRoom.InRoom)-1]
				}
			}
		}
	case "домой":
		for i, el := range rooms.HomeRoom.InRoom {
			if el.Name == player.Name {
				if i != len(rooms.HomeRoom.InRoom)-1 {
					rooms.HomeRoom.InRoom = append(rooms.HomeRoom.InRoom[:i], rooms.HomeRoom.InRoom[i+1:]...)
				} else {
					rooms.HomeRoom.InRoom = rooms.HomeRoom.InRoom[:len(rooms.HomeRoom.InRoom)-1]
				}
			}
		}
	case "улица":
		for i, el := range rooms.StreetRoom.InRoom {
			if el.Name == player.Name {
				if i != len(rooms.StreetRoom.InRoom)-1 {
					rooms.StreetRoom.InRoom = append(rooms.StreetRoom.InRoom[:i], rooms.StreetRoom.InRoom[i+1:]...)
				} else {
					rooms.StreetRoom.InRoom = rooms.StreetRoom.InRoom[:len(rooms.StreetRoom.InRoom)-1]
				}
			}
		}
	}
}
