package main

import (
	"fmt"
	"os"
)

func GenDisplaceFn(a, v, d float64) func(t float64) float64 {

	return func(t float64) float64 {
		return ((a * t * t) / 2) + (v * t) + d
	}

}

func getFloatValue(text string) float64 {
	var value float64
	fmt.Printf("Enter %s: ", text)
	_, err := fmt.Scan(&value)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return value
}

func main() {

	acc := getFloatValue("acceloration")
	vel := getFloatValue("velocity")
	initDisplace := getFloatValue("displacement")
	time := getFloatValue("time")

	fn := GenDisplaceFn(acc, vel, initDisplace)
	fmt.Println(fn(time))

}
