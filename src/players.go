package src

import "strings"

type Player struct {
	Room
	CurrentRoom string
	Inventory   []string
	Do          map[string]func([]string, *Player) string
}

func Look(commands []string, player *Player) string {
	var answer string
	switch player.CurrentRoom {
	case "кухня":
		answer += "ты находишься на кухне, "
		if len(player.KitchenRoom.Table) > 0 {
			answer += "на столе "
			for _, el := range player.KitchenRoom.Table {
				answer += el + ", "
			}
		}
		if len(player.KitchenRoom.NeedToDo) > 0 {
			answer += "надо "
			for i, el := range player.KitchenRoom.NeedToDo {
				if i != len(player.KitchenRoom.NeedToDo)-1 {
					answer += el + " и "
				} else {
					answer += el + ". "
				}
			}
		}
		answer += "можно пройти - "
		for i, el := range player.KitchenRoom.Entrance {
			if i != len(player.KitchenRoom.Entrance)-1 {
				answer += el + ", "
			} else {
				answer += el
			}
		}
	case "комната":
		if len(player.BedRoom.Table)+len(player.BedRoom.Chain) == 0 {
			answer += "пустая комната. "
		} else {

			if len(player.BedRoom.Table) > 0 {
				answer += "на столе: "
				for _, el := range player.BedRoom.Table {
					answer += el + ", "
				}
			}
			if len(player.BedRoom.Chain) > 0 {
				answer += "на стуле - "
				for i, el := range player.BedRoom.Chain {
					if i != len(player.BedRoom.Chain)-1 {
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
		for i, el := range player.BedRoom.Entrance {
			if i != len(player.BedRoom.Entrance)-1 {
				answer += el + ", "
			} else {
				answer += el
			}
		}
	case "коридор":
		answer += "ничего интересного. "
		for i, el := range player.LineRoom.Entrance {
			if i != len(player.LineRoom.Entrance)-1 {
				answer += el + ", "
			} else {
				answer += el
			}
		}
	case "домой":
		answer += "ничего интересного. "
		for i, el := range player.HomeRoom.Entrance {
			if i != len(player.HomeRoom.Entrance)-1 {
				answer += el + ", "
			} else {
				answer += el
			}
		}
	case "улица":
		answer += "ничего интересного. "
		for i, el := range player.StreetRoom.Entrance {
			if i != len(player.StreetRoom.Entrance)-1 {
				answer += el + ", "
			} else {
				answer += el
			}
		}
	}
	return answer
}
func Walk(commands []string, person *Player) string {
	var answer string
	switch commands[1] {
	case "кухня":
		if search(person.CurrentRoom, person.KitchenRoom.Entrance) {
			person.CurrentRoom = "кухня"
			answer += "кухня, ничего интересного. "
			answer += "можно пройти - "
			for i, el := range person.KitchenRoom.Entrance {
				if i != len(person.KitchenRoom.Entrance)-1 {
					answer += el + ", "
				} else {
					answer += el
				}
			}
		} else {
			answer = "нет пути в кухня"
		}
	case "комната":
		if search(person.CurrentRoom, person.BedRoom.Entrance) {
			person.CurrentRoom = "комната"
			answer += "ты в своей комнате. "
			answer += "можно пройти - "
			for i, el := range person.BedRoom.Entrance {
				if i != len(person.BedRoom.Entrance)-1 {
					answer += el + ", "
				} else {
					answer += el
				}
			}
		} else {
			answer = "нет пути в комната"
		}
	case "коридор":
		if search(person.CurrentRoom, person.LineRoom.Entrance) {
			person.CurrentRoom = "коридор"
			answer += "ничего интересного. "
			answer += "можно пройти - "
			for i, el := range person.LineRoom.Entrance {
				if i != len(person.LineRoom.Entrance)-1 {
					answer += el + ", "
				} else {
					answer += el
				}
			}
		} else {
			answer = "нет пути в коридор"
		}
	case "домой":
		if search(person.CurrentRoom, person.HomeRoom.Entrance) {
			person.CurrentRoom = "домой"
			answer += "ничего интересного. "
			answer += "можно пройти - "
			for i, el := range person.HomeRoom.Entrance {
				if i != len(person.HomeRoom.Entrance)-1 {
					answer += el + ", "
				} else {
					answer += el
				}
			}
		} else {
			answer = "нет пути в домой"
		}
	case "улица":
		if person.StreetRoom.Door {
			if search(person.CurrentRoom, person.StreetRoom.Entrance) {
				person.CurrentRoom = "улица"
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
func Dress(commands []string, person *Player) string {
	var answer string
	switch commands[1] {
	case "рюкзак":
		if search("рюкзак", person.BedRoom.Chain) {
			person.Inventory = append(person.Inventory, "рюкзак")
			removeItem("рюкзак", &person.BedRoom.Chain)
			answer = "вы одели: рюкзак"
		} else {
			answer = "нет такого"
		}
	}
	return answer
}
func Take(commands []string, person *Player) string {
	var answer string
	if search("рюкзак", person.Inventory) {
		switch commands[1] {
		case "ключи":
			if search("ключи", person.BedRoom.Table) {
				person.Inventory = append(person.Inventory, "ключи")
				removeItem("ключи", &person.BedRoom.Table)
				answer = "предмет добавлен в инвентарь: ключи"
			} else {
				answer = "нет такого"
			}
		case "конспекты":
			if search("конспекты", person.BedRoom.Table) {
				person.Inventory = append(person.Inventory, "конспекты")
				removeItem("конспекты", &person.BedRoom.Table)
				removeItem("собрать рюкзак", &person.KitchenRoom.NeedToDo)
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
func Apply(commands []string, person *Player) string {
	var answer string
	if search(commands[1], person.Inventory) {
		switch commands[2] {
		case "дверь":
			if person.CurrentRoom == "коридор" {
				person.StreetRoom.Door = true
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
