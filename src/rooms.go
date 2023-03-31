package src

type Room struct {
	KitchenRoom
	BedRoom
	LineRoom
	HomeRoom
	StreetRoom
}

func NewRoom() Room {
	return Room{
		KitchenRoom: KitchenRoom{
			Name:     "кухня",
			Entrance: []string{"коридор"},
			Table:    []string{"чай"},
			NeedToDo: []string{"собрать рюкзак", "идти в универ"},
		},
		BedRoom: BedRoom{
			Name:     "комната",
			Entrance: []string{"коридор"},
			Table:    []string{"ключи", "конспекты"},
			Chain:    []string{"рюкзак"},
		},
		LineRoom: LineRoom{
			Name:     "коридор",
			Entrance: []string{"кухня", "комната", "улица"},
		},
		HomeRoom: HomeRoom{
			Name:     "домой",
			Entrance: []string{},
		},
		StreetRoom: StreetRoom{
			Name:     "улица",
			Entrance: []string{"домой", "коридор"},
			Door:     false,
		},
	}
}

type KitchenRoom struct {
	Name     string
	Entrance []string
	Table    []string
	NeedToDo []string
}

type BedRoom struct {
	Name     string
	Entrance []string
	Table    []string
	Chain    []string
}
type LineRoom struct {
	Name     string
	Entrance []string
}
type HomeRoom struct {
	Name     string
	Entrance []string
}

type StreetRoom struct {
	Name     string
	Entrance []string
	Door     bool
}
