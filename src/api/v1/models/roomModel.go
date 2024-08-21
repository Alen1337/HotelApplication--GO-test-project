package api_v1_models

type RoomStatus int

const (
	Dirty = iota
	Clean
	Inspected
)

type Room struct {
	ID     int
	Status RoomStatus
}
