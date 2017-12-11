package entity

// An Entity is the interface for anything which
// exists in the world and can move around.
type Entity interface {
	Up()
	Down()
	Left()
	Right()
	GetX() int
	GetY() int
	GetRoomX() int
	GetRoomY() int
	Character() rune
}
