package race

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lakrsv/go-kennel/internal/pkg/animal"
	"github.com/lakrsv/go-kennel/internal/pkg/track"
)

var FlavorText = [...]string{
	"What an exciting turn of events!",
	"That was unexpected.",
	"The audience is jumping out of their seats!",
	"Fairly impressive.",
	"Not exactly a fan favorite.",
	"Questionable, but the audience seems excited!",
}

type RaceState struct {
	animal       *animal.Animal
	currentSpeed float32
	travelled    float32
}

type RaceInfo struct {
	animal    *animal.Animal
	message   string
	finished  bool
	timestamp time.Time
}

func Describe(track *track.Track) {
	fmt.Println("Let the games begin!")
	fmt.Println("--------------------")
	fmt.Println("The participants are:")
	for _, animal := range track.Participants() {
		fmt.Printf("%s: %s. It is %d years old.\n", animal.Name(), animal.Description(), animal.Age())
	}
	fmt.Println("--------------------")
	fmt.Printf("In this race, the participants will have to run a total of %.2f kilometers. Good luck everybody!\n", track.Length())
	fmt.Println("--------------------")
}

func Race(track *track.Track) {
	fmt.Println("Ready...")
	time.Sleep(1 * time.Second)
	fmt.Println("Set...")
	time.Sleep(1 * time.Second)
	fmt.Println("GO!")
	fmt.Println("--------------------")

	state := make([]*RaceState, len(track.Participants()))
	for i, animal := range track.Participants() {
		state[i] = &RaceState{animal, animal.Speed(), 0}
	}

	raceChannel := make(chan RaceInfo)
	for _, state := range state {
		go Tick(state, track.Length(), raceChannel)
	}

	finished := 0
	bestTime := time.Unix(1<<63-62135596801, 999999999)
	var winner *animal.Animal
	for info := range raceChannel {
		fmt.Print(info.message)
		if info.finished {
			finished += 1
			if info.timestamp.Before(bestTime) {
				bestTime = info.timestamp
				winner = info.animal
			}
		}
		if finished == len(state) {
			break
		}
	}
	fmt.Println("--------------------")
	fmt.Println("The race is finished!")
	fmt.Println("The winner is...")
	time.Sleep(1 * time.Second)
	fmt.Printf("%s!\n", winner.Name())
}

func Tick(state *RaceState, trackLength float32, raceChannel chan RaceInfo) {
	for {
		state.travelled += state.currentSpeed / 1000
		state.currentSpeed *= 1 - state.animal.SpeedDecay()

		if state.travelled >= trackLength {
			raceChannel <- RaceInfo{state.animal, fmt.Sprintf("%s just passed the finish line! %s\n", state.animal.Name(), FlavorText[rand.Intn(len(FlavorText))]), true, time.Now()}
			return
		} else if rand.Float32() > 0.6 {
			raceChannel <- RaceInfo{state.animal, fmt.Sprintf("%s has travelled %.2f km so far! Only %.2f km to the finish line! %s\n", state.animal.Name(), state.travelled, trackLength-state.travelled, FlavorText[rand.Intn(len(FlavorText))]), false, time.Now()}
		}
		time.Sleep((time.Duration(250 + rand.Intn(250))) * time.Millisecond)
	}
}
