package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	sli := make([]int, 3)
	sliceInt := sli[0:0]

	var usrInput string
	for {
		fmt.Printf("enter number: ")
		_, err := fmt.Scan(&usrInput)
		if err != nil {
			fmt.Println(err)
		} else {
			if usrInput == "X" {
				break
			}
			inputInt, err := strconv.Atoi(usrInput)
			if err != nil {
				fmt.Println(err)
			} else {
				sliceInt = append(sliceInt, inputInt)
				sort.Slice(sliceInt, func(i, j int) bool {
					return sliceInt[i] < sliceInt[j]
				})
				fmt.Println(sliceInt)
			}
		}
	}
}
