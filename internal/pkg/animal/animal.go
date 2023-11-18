package animal

type Animal struct {
	name        string
	description string
	age         int
	speed       float32
	speedDecay  float32
}

func NewAnimal(name string, description string, age int, speed float32, speedDecay float32) *Animal {
	return &Animal{name, description, age, speed, speedDecay}
}

func (a *Animal) Name() string {
	return a.name
}

func (a *Animal) Description() string {
	return a.description
}

func (a *Animal) Age() int {
	return a.age
}

func (a *Animal) Speed() float32 {
	return a.speed
}

func (a *Animal) SpeedDecay() float32 {
	return a.speedDecay
}
