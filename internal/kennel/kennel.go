package kennel

import (
	"fmt"

	"github.com/lakrsv/go-kennel/internal/pkg/animal"
	"github.com/lakrsv/go-kennel/internal/pkg/kennel"
	"github.com/lakrsv/go-kennel/internal/pkg/race"
	"github.com/lakrsv/go-kennel/internal/pkg/track"
)

func Simulate() {
	kennel := kennel.NewKennel("Dogmatic Racers", 10)
	fmt.Println(kennel)

	// Add the animals
	kennel.Add(animal.NewAnimal("Barky", "Barks a lot", 5, 300.14, 0.0125))
	kennel.Add(animal.NewAnimal("Hoofsworth", "Has a tendency to stampede", 10, 700.14, 0.065))
	kennel.Add(animal.NewAnimal("Purraloaf", "Shorter whiskers than most", 5, 1040.124, 0.0325))
	kennel.Add(animal.NewAnimal("Bitey", "Teeth have dulled over time", 15, 600, 0.045))

	// Setup the track
	track := track.NewTrack(10)
	for _, animal := range kennel.Animals() {
		track.Add(animal)
	}

	// Start the race
	race.Describe(track)
	race.Race(track)
}
