package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BubbleSort(sli []int) {
	for j := len(sli); j > 0; j-- {
		for i := 0; i < j-1; i++ {
			if sli[i] > sli[i+1] {
				Swap(sli, i)
			}
		}
	}
}

func Swap(sli []int, i int) {
	sli[i], sli[i+1] = sli[i+1], sli[i]
}

func main() {
	var integers []int

	fmt.Println("Enter up to 10 integers sepretad with space: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	strinNumbers := strings.Split(line, " ")
	if len(strinNumbers) > 10 {
		strinNumbers = strinNumbers[0:10]
	}

	for _, v := range strinNumbers {
		if num, err := strconv.Atoi(v); err == nil {
			integers = append(integers, num)
		} else {
			fmt.Printf("Error : %s\n", err)
			os.Exit(1)
		}
	}

	BubbleSort(integers)
	fmt.Println(integers)

}
