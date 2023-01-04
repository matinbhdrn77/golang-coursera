package main

import (
	"bufio"
	"fmt"
	"os"
	_ "strconv"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

type Snake struct{}

func (c Snake) Eat() {
	fmt.Println("mice")
}

func (c Snake) Move() {
	fmt.Println("slither")
}

func (c Snake) Speak() {
	fmt.Println("hsss")
}

type Bird struct{}

func (c Bird) Eat() {
	fmt.Println("worms")
}

func (c Bird) Move() {
	fmt.Println("fly")
}

func (c Bird) Speak() {
	fmt.Println("peep")
}

func main() {
	var animals = make(map[string]Animal)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}
		input = strings.TrimSpace(input)
		request := strings.Split(input, " ")
		switch strings.ToLower(request[0]) {
		case "newanimal":
			switch strings.ToLower(request[2]) {
			case "cow":
				animals[strings.ToLower(request[1])] = Cow{}
			case "bird":
				animals[strings.ToLower(request[1])] = Bird{}
			case "snake":
				animals[strings.ToLower(request[1])] = Snake{}
			default:
				fmt.Println("Invalid animal type")
				continue
			}
		case "query":
			animal, ok := animals[strings.ToLower(request[1])]
			if !ok {
				fmt.Println("Invalid name")
				continue
			}
			switch strings.ToLower(request[2]) {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Invalid request")
				continue
			}
		}
	}
}
