package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

const maxLength = 20

func (n *Name) Set(firstName string, lastName string) {
	n.fname = firstName
	n.lname = lastName

	if len(firstName) > maxLength {
		n.fname = firstName[:maxLength]
	}

	if len(lastName) > maxLength {
		n.lname = lastName[:maxLength]
	}
}

func main() {
	var persons []Name
	var fileName string

	fmt.Printf("Enter file name: ")
	fmt.Scan(&fileName)
	readFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	var person Name
	for fileScanner.Scan() {
		line := fileScanner.Text()
		lineSlice := strings.Split(line, " ")

		person.Set(lineSlice[0], lineSlice[1])
		persons = append(persons, person)
	}

	readFile.Close()

	for _, person := range persons {
		fmt.Printf("Name: %v, Lastname: %v\n", person.fname, person.lname)
	}
}
