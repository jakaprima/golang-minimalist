package concurency

import (
	"fmt"
	"sync"
	"time"
)

func goroutineExample1(n int){
	fmt.Println("EXECUTE goroutineExample1")
	for i := 0; i < 10; i++ {
        fmt.Println(n, ":", i)
    }
}

func producer(ch chan<- int) { //goroutine sends numbers to the channel.
    for i := 0; i < 10; i++ {
        ch <- i
    }
    close(ch) //  function signals the end of the data stream.
}

func consumer(ch <-chan int) { //goroutine receives numbers from the channel and prints them.
    for {
        val, ok := <-ch
        if !ok {
            break
        }
        fmt.Println(val)
    }
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine finishes

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d finished\n", id)
}

func GorountineExample(){
	fmt.Println("go routiune example")

	// Write concurrent code using goroutines.
	go goroutineExample1(0) // tidak akan terprint karena akan tereksekute secara promise tapi tereksekusi
	goroutineExample1(1)

	// go routine dan channels
	messages := make(chan string)
	go func() {
		messages <- "ping"
	}()
	msg := <-messages
	fmt.Println("msg value: ", msg)

	// Use channels for communication between goroutines.
	ch := make(chan int) // Channels are used for communication and synchronization between goroutines.
    go producer(ch)
    consumer(ch)

	// using wait group
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
			wg.Add(1) // Increment the counter for each goroutine
			go worker(i, &wg)
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All goroutines finished")
	
}

