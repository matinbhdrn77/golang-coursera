package main

import "fmt"

func main() {
	var userInput float64
	fmt.Printf("enter a floating point number: ")
	_, err := fmt.Scan(&userInput)
	if err != nil {
		fmt.Println(err)
	} else {
		intiger := int(userInput)
		fmt.Println(intiger)
	}
}
