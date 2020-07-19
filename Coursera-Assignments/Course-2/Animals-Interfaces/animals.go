package main

import (
	"fmt"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type cow struct {
	food       string
	locomotion string
	noise      string
}

func (a cow) Eat() {
	fmt.Println(a.food)
}

func (a cow) Move() {
	fmt.Println(a.locomotion)
}

func (a cow) Speak() {
	fmt.Println(a.noise)
}

type bird struct {
	food       string
	locomotion string
	noise      string
}

func (a bird) Eat() {
	fmt.Println(a.food)
}

func (a bird) Move() {
	fmt.Println(a.locomotion)
}

func (a bird) Speak() {
	fmt.Println(a.noise)
}

type snake struct {
	food       string
	locomotion string
	noise      string
}

func (a snake) Eat() {
	fmt.Println(a.food)
}

func (a snake) Move() {
	fmt.Println(a.locomotion)
}

func (a snake) Speak() {
	fmt.Println(a.noise)
}

func main() {
	AnimalArr := make(map[string]Animal) //Name:Animal

	var cmd, arg1, arg2 string

	for {
		fmt.Print("> ")
		fmt.Scanf("%s %s %s", &cmd, &arg1, &arg2)
		if strings.Compare(cmd, "newanimal") == 0 {
			name := arg1
			animalType := arg2
			flag := true

			switch animalType {
			case "cow":
				myAnimal := cow{"grass", "walk", "moo"}
				AnimalArr[name] = myAnimal
			case "bird":
				myAnimal := bird{"worms", "fly", "peep"}
				AnimalArr[name] = myAnimal
			case "snake":
				myAnimal := snake{"mice", "slither", "hsss"}
				AnimalArr[name] = myAnimal
			default:
				fmt.Println("Invalid Animal!")
				flag = false
			}
			if flag {
				fmt.Println("Created it!")
			}

		} else if strings.Compare(cmd, "query") == 0 {
			name := arg1
			action := arg2

			if _, ok := AnimalArr[name]; ok {
				switch action {
				case "eat":
					AnimalArr[name].Eat()
				case "move":
					AnimalArr[name].Move()
				case "speak":
					AnimalArr[name].Speak()
				default:
					fmt.Println("Invalid Action!")
				}
			} else {
				fmt.Println("Invalid Animal Name!")
			}

		} else {
			fmt.Println("Invalid Command!")
		}
	}
}
