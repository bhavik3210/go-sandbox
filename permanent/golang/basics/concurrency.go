package basics

import (
	"fmt"
	"sync"
	"time"
)

func DemoConcurrency() {
	PrintHeader("Concurrency in Go")
	// GoroutineDemo1()
	// ChannelDemo()
	// SelectStatementDemo()
	LoopOverChannelsDemo()
}

func GoroutineDemo1() {
	PrintTitle(" SIMPLE GOROUTINES ")
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		fmt.Println("This happens asynchronously")
		wg.Done()
	}()
	fmt.Println("This is synchronous")
	wg.Wait()
}

// channel operations block until the complementary operation is ready, it make both sides wait until ready
func ChannelDemo() {
	PrintTitle(" CHANNELS ")
	var wg sync.WaitGroup
	ch := make(chan string)
	wg.Add(1)

	go func() {
		ch <- "THE MESSAGE"
	}()

	go func() {
		fmt.Println(<-ch)
		wg.Done()
	}()
	wg.Wait()
}

// In a select statement, if more than one case
// can be acted upon then one case is chosen
// RANDOMLY
func SelectStatementDemo() {
	PrintTitle(" SELECT STATEMENT ")

	ch1, ch2 := make(chan string), make(chan string)

	go func() {
		ch1 <- "Message to channel 1"
	}()
	go func() {
		ch2 <- "Message to channel 2"
	}()
	time.Sleep(10 * time.Millisecond)

	// select statements are not switch statement
	// so it doesn't follow top to bottom order
	// it will select one at random (non-blocking)
	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	default:
		fmt.Println("No Message Available")
	}

}

func LoopOverChannelsDemo() {
	PrintTitle(" LOOP OVER CHANNELS ")
	ch := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch) // no more messages can be sent after closing to this channel, this will stop/break the for loop below that is printing, otherwise it will run into deadlock if this channel is not closed
	}()

	for msg := range ch {
		fmt.Println(msg)
	}
}
