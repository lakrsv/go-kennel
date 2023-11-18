package track

import "github.com/lakrsv/go-kennel/internal/pkg/animal"

type Track struct {
	length       float32
	participants []*animal.Animal
}

func NewTrack(length float32) *Track {
	return &Track{length, make([]*animal.Animal, 0)}
}

func (t *Track) Add(animal *animal.Animal) {
	t.participants = append(t.participants, animal)
}

func (t *Track) Length() float32 {
	return t.length
}

func (t *Track) Participants() []*animal.Animal {
	return t.participants
}
