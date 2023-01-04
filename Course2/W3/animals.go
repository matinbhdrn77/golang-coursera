package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food, locomotion, noise string
}

var (
	cow = Animal{
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	}

	bird = Animal{
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	}

	snake = Animal{
		food:       "mice",
		locomotion: "slither",
		noise:      "hss",
	}

	animalMap = map[string]Animal{
		"cow":   cow,
		"bird":  bird,
		"snake": snake,
	}
)

func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}

func main() {

	for {
		fmt.Printf(">")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()

		commands := strings.Split(line, " ")
		if len(commands) != 2 {
			continue
		}

		animal := animalMap[commands[0]]

		switch commands[1] {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		}

	}

}
