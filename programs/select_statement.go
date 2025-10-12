package programs

import (
	"log"
	"time"
)

// Rules of select
// 1. A select waits until at least one case is ready.
// 2. If multiple cases are ready, one is chosen randomly.
// 3. If no cases are ready:
// It blocks (waits), unless…
// There is a default case, in which case it executes that immediately.

func SelectStatmentExample() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	// chan1 waits for 3 seconds and then return
	go func() {
		time.Sleep(3 * time.Second)
		chan1 <- "chan1"
	}()

	// chan2 waits for 1 second and then return
	go func() {
		time.Sleep(1 * time.Second)
		chan2 <- "chan2"
	}()

	// Without using select statement
	// WithoutSelectStatement(chan1, chan2)

	// Using select statement
	WithSelectStatement(chan1, chan2)
}

// Here you first wait on chan1 → that blocks for 3s until chan1 sends.
// Then you wait on chan2 → but by now, chan2’s message (after 1s) is still buffered in the channel, so you receive it immediately.

func WithoutSelectStatement(chan1, chan2 chan string) {
	msg1 := <-chan1
	log.Println("Message from ", msg1)
	msg2 := <-chan2
	log.Println("Message from ", msg2)
}

// Explanation:
// Second select waits → chan2 is ready after 1 second → prints "Message from chan2".
// Then chan1 becomes ready after 2 more seconds → prints "Message from chan1".

func WithSelectStatement(chan1, chan2 chan string) {
	for {
		select {
		case msg1 := <-chan1:
			log.Println("Message from ", msg1)
		case msg2 := <-chan2:
			log.Println("Message from ", msg2)
		case <-time.After(3 * time.Second):
			log.Println("No more messages, exiting...")
			return
		}
	}
}
