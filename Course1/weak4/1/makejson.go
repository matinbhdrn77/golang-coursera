package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	usrInfoMap := make(map[string]string)

	// getting user name
	var usrName string
	fmt.Println("Enter your name: ")
	fmt.Scan(&usrName)
	usrInfoMap["name"] = usrName

	// getting user address
	fmt.Println("Enter your address: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	usrAddr := scanner.Text()
	usrInfoMap["address"] = usrAddr

	if barr, err := json.Marshal(usrInfoMap); err == nil {
		fmt.Println(string(barr))
	} else {
		fmt.Println(err)
	}

}
