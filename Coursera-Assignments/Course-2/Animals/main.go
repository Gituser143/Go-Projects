package main

import (
	"fmt"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func (a *Animal) init(food, locomotion, noise string) {
	a.food = food
	a.locomotion = locomotion
	a.noise = noise
}

func main() {
	cow := Animal{}
	cow.init("grass", "walk", "moo")

	bird := Animal{}
	bird.init("worms", "fly", "peep")

	snake := Animal{}
	snake.init("mice", "slither", "hsss")

	for {
		var animal, action string
		fmt.Print("> ")
		fmt.Scanf("%s %s", &animal, &action)

		switch animal {
		case "cow":
			switch action {
			case "eat":
				cow.Eat()
			case "move":
				cow.Move()
			case "speak":
				cow.Speak()
			default:
				fmt.Println("Invalid Action")
			}
		case "bird":
			switch action {
			case "eat":
				bird.Eat()
			case "move":
				bird.Move()
			case "speak":
				bird.Speak()
			default:
				fmt.Println("Invalid Action")
			}
		case "snake":
			switch action {
			case "eat":
				snake.Eat()
			case "move":
				snake.Move()
			case "speak":
				snake.Speak()
			default:
				fmt.Println("Invalid Action")
			}
		default:
			fmt.Println("Invalid Animal")
		}
	}

}
