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
	// wg.Done()
}

func dine(p []*philo, ch chan int) {
	index := <-ch
	// fmt.Println(index)
	p[index].eat(index)
	wg.Done()
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

	ch := make(chan int, 1)

	for i := 0; i < 5; i++ { // Each philo.eat() is called twice
		wg.Add(1)
		ch <- i
		go dine(philos, ch)
		ch <- (i + 2) % 5
		go dine(philos, ch)
		wg.Wait()
	}
}
