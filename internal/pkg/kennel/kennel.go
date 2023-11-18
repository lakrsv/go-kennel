package kennel

import (
	"errors"

	"github.com/lakrsv/go-kennel/internal/pkg/animal"
)

type Kennel struct {
	name    string
	animals []*animal.Animal
}

func NewKennel(name string, capacity int) *Kennel {
	return &Kennel{name, make([]*animal.Animal, 0, capacity)}
}

func (k *Kennel) Add(animal *animal.Animal) error {
	if len(k.animals) == cap(k.animals) {
		return errors.New("Animals is full")
	}
	k.animals = append(k.animals, animal)
	return nil
}

func (k *Kennel) Name() string {
	return k.name
}

func (k *Kennel) Animals() []*animal.Animal {
	return k.animals
}
