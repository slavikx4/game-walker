package src

import "sync"

type Room struct {
	KitchenRoom
	BedRoom
	LineRoom
	HomeRoom
	StreetRoom
	Mu sync.Mutex
}

var Rooms Room

func StartWorld() {
	initGame()
}
func initGame() {
	Rooms = NewRoom()
}
func NewRoom() Room {
	return Room{
		KitchenRoom: KitchenRoom{
			Name:     "кухня",
			Entrance: []string{"коридор"},
			Table:    []string{"чай"},
			InRoom:   []*Player{},
		},
		BedRoom: BedRoom{
			Name:     "комната",
			Entrance: []string{"коридор"},
			Table:    []string{"ключи", "конспекты"},
			Chain:    []string{"рюкзак"},
			InRoom:   []*Player{},
		},
		LineRoom: LineRoom{
			Name:     "коридор",
			Entrance: []string{"кухня", "комната", "улица"},
			InRoom:   []*Player{},
		},
		HomeRoom: HomeRoom{
			Name:     "домой",
			Entrance: []string{"коридор"},
			InRoom:   []*Player{},
		},
		StreetRoom: StreetRoom{
			Name:     "улица",
			Entrance: []string{"домой", "коридор"},
			Door:     false,
			InRoom:   []*Player{},
		},
	}
}

type KitchenRoom struct {
	Name     string
	Entrance []string
	Table    []string
	InRoom   []*Player
}

type BedRoom struct {
	Name     string
	Entrance []string
	Table    []string
	Chain    []string
	InRoom   []*Player
}
type LineRoom struct {
	Name     string
	Entrance []string
	InRoom   []*Player
}
type HomeRoom struct {
	Name     string
	Entrance []string
	InRoom   []*Player
}

type StreetRoom struct {
	Name     string
	Entrance []string
	Door     bool
	InRoom   []*Player
}
