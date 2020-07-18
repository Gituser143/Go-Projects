package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type person struct {
	fname string
	lname string
}

func main() {
	var filename string
	fmt.Println("Enter filename")
	fmt.Scanf("%s", &filename)

	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	l := string(bs)
	lines := strings.Split(l, "\n")
	lines = lines[0 : len(lines)-1]
	personSlice := make([]person, 0)

	for _, line := range lines {
		values := strings.Split(line, " ")
		p := person{values[0], values[1]}
		personSlice = append(personSlice, p)
	}

	for _, p := range personSlice {
		fmt.Println(p.fname, p.lname)
	}
}
