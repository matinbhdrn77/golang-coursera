package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	fmt.Printf("enter a string: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	line := scanner.Text()

	line = strings.ToLower(line)
	fmt.Println(line)
	matched, err := regexp.MatchString(`^i.*a+.*n$`, line)
	if err != nil {
		fmt.Println(err)
	} else {
		if matched {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}
	}

}
