package builder

type House struct {
	windowType string
	doorType   string
	floor      int
}

type HouseBuilder struct {
	h House
}

func NewHouseBuilder() *HouseBuilder {
	return &HouseBuilder{}
}

func (b *HouseBuilder) SetWindowType(windowType string) *HouseBuilder {
	b.h.windowType = windowType
	return b
}

func (b *HouseBuilder) SetDoorType(doorType string) *HouseBuilder {
	b.h.doorType = doorType
	return b
}

func (b *HouseBuilder) SetFloor(floor int) *HouseBuilder {
	b.h.floor = floor
	return b
}

func (b *HouseBuilder) Build() House {
	return b.h
}
