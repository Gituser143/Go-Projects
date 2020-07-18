package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	jsonmap := make(map[string]string)

	var name, address string
	fmt.Println("Enter Name")
	fmt.Scanf("%s", &name)

	fmt.Println("Enter Address")
	fmt.Scanf("%s", &address)

	jsonmap["name"] = name
	jsonmap["address"] = address

	jsonObj, err := json.Marshal(jsonmap)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nThe JSON object is")
	fmt.Println(string(jsonObj))
}
