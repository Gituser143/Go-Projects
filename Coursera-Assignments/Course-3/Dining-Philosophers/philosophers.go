package main

import (
	"fmt"
	"sync"
	"time"
)

type chopStick struct {
	sync.Mutex
}

type philo struct {
	leftCS, rightCS *chopStick
}

var wg sync.WaitGroup

const num = 5                     // Number of Philosophers
const hunger = 2                  // Number of times each philosopher eats
const timeToEat = 2 * time.Second // Time taken to eat once

func (p philo) eat(index int) {
	for i := 0; i < hunger; i++ {
		if index%2 == 0 {
			p.rightCS.Lock()
			p.leftCS.Lock()
		} else {
			p.leftCS.Lock()
			p.rightCS.Lock()
		}

		fmt.Println("Starting to Eat", index)
		time.Sleep(timeToEat)
		fmt.Println("Finished Eating", index)

		p.leftCS.Unlock()
		p.rightCS.Unlock()
	}
	wg.Done()
}

func main() {

	chopSticks := make([]*chopStick, num)
	for i := 0; i < num; i++ {
		chopSticks[i] = new(chopStick)
	}

	philos := make([]*philo, num)
	for i := 0; i < num; i++ {
		philos[i] = &philo{chopSticks[i], chopSticks[(i+1)%num]}
	}

	wg.Add(num)

	for i := 0; i < num; i++ {
		go philos[i].eat(i)
	}

	wg.Wait()
}
