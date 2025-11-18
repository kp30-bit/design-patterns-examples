package builder

import "fmt"

type Elevator struct {
	ID           string
	Floors       int
	Speed        float64
	HasFireAlarm bool
	HasCamera    bool
	Mode         string
}

type ElevatorBuilder struct {
	id           string
	floors       int
	speed        float64
	hasFireAlarm bool
	hasCamera    bool
	mode         string
}

// func NewElevator(id string, floors int, speed float64, fireAlarm bool, camera bool, mode string) *Elevator { //Incorect way to construct an object with too many params
// 	return &Elevator{id, floors, speed, fireAlarm, camera, mode}
// }

func NewElevatorBuilder(id string, floors int) *ElevatorBuilder {
	return &ElevatorBuilder{id: id, floors: floors}
}

func (b *ElevatorBuilder) Speed(speed float64) *ElevatorBuilder {
	b.speed = speed
	return b
}

func (b *ElevatorBuilder) WithFireAlarm() *ElevatorBuilder {
	b.hasFireAlarm = true
	return b
}

func (b *ElevatorBuilder) WithCamera() *ElevatorBuilder {
	b.hasCamera = true
	return b
}

func (b *ElevatorBuilder) Mode(mode string) *ElevatorBuilder {
	b.mode = mode
	return b
}

func (b *ElevatorBuilder) Build() (*Elevator, error) {
	if b.id == "" || b.floors <= 0 {
		return nil, fmt.Errorf("invalid ID or floors")
	}

	return &Elevator{
		ID:           b.id,
		Floors:       b.floors,
		Speed:        b.speed,
		HasFireAlarm: b.hasFireAlarm,
		HasCamera:    b.hasCamera,
		Mode:         b.mode,
	}, nil
}
