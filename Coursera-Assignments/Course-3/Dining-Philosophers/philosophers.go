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

func dine(p []*philo, ch chan int) {
	for {
		index := <-ch
		switch index {
		case -1:
			return
		default:
			go p[index].eat(index)
		}
	}
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

	ch := make(chan int, 6)

	go dine(philos, ch)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		ch <- i
	}
	ch <- -1
	wg.Wait()
}
