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

func (p philo) eat(index int) {
	for i := 0; i < 3; i++ {
		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Println("Starting to Eat", index)
		time.Sleep(1 * time.Second)
		fmt.Println("Finished Eating", index)
		p.leftCS.Unlock()
		p.rightCS.Unlock()
	}
	wg.Done()
}

func dine(index int, ch chan int) {

}

func main() {
	chopSticks := make([]*chopStick, 5)
	for i := 0; i < 5; i++ {
		chopSticks[i] = new(chopStick)
	}

	philos := make([]*philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &philo{chopSticks[i], chopSticks[(i+1)%5]}
	}

	// ch := make(chan int)

	for i := 0; i < 5; i++ {
		wg.Add(2)
		go philos[i].eat(i)
		go philos[(i+2)%5].eat((i + 2) % 5)
		wg.Wait()
	}
}
