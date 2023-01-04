package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var toSplit = 4

func sorting(sli []int, wg *sync.WaitGroup) {

	sort.Slice(sli, func(i, j int) bool {
		return sli[i] < sli[j]
	})

	if wg != nil {
		fmt.Println("sorted chunk: ", sli)
		wg.Done()
	} else {
		fmt.Println("sorted array: ", sli)
	}
}

// function to convert slice of strings to slice of integers
func sliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}

func main() {
	var wg sync.WaitGroup

	// prompt user and receive list on numbers
	fmt.Printf("Enter sequence of numbers seperated with space: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	//convert line to slice of numbers
	stringNums := strings.Split(line, " ")
	sliNums, err := sliceAtoi(stringNums)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//split array and sort each subarray
	sliLen := len(sliNums)
	chunkSize := (sliLen + toSplit - 1) / toSplit
	for i := 0; i < sliLen; i += chunkSize {
		end := i + chunkSize
		if end > sliLen {
			end = sliLen
		}
		wg.Add(1)
		go sorting(sliNums[i:end], &wg)
	}

	wg.Wait()
	sorting(sliNums, nil)

	// fmt.Println("sorted sli: ", sliNums)
}
