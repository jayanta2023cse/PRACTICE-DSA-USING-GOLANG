package programs

import (
	"fmt"
	"time"
)

func UnbufferedChannel() {
	// Phone Call (Unbuffered)
	call := make(chan string) // unbuffered

	go func() {
		fmt.Println("Caller: Trying to call friend...")
		call <- "Hello, are you there?" // blocked until receiver picks up
		fmt.Println("Caller: Message delivered (friend picked up).")
	}()

	go func() {
		time.Sleep(4 * time.Second) // simulate delay in picking up
		msg := <-call
		fmt.Println("Friend: Picked up the call and heard:", msg)
	}()

	time.Sleep(5 * time.Second) // give enough time for goroutines
	fmt.Println("Call Finished")
}

func BufferedChannel() {
	// Vouce Mail (Buffered)
	voicemail := make(chan string, 1) // buffered with space for 2 messages

	go func() {
		fmt.Println("\nCaller: Leaving 3 voicemails...")

		voicemail <- "Hey, call me back!"
		fmt.Println("1st Voicemail Sent")

		voicemail <- "Don’t forget the meeting tomorrow."
		fmt.Println("2nd Voicemail Sent")

		voicemail <- "Hope ypu recived my messages"
		fmt.Println("3rd Voicemail Sent")

		fmt.Println("Caller: Done leaving messages.")
	}()

	go func() {
		time.Sleep(2 * time.Second) // friend is busy, checks later
		fmt.Println("Friend: Checking voicemail now...")
		for len(voicemail) > 0 {
			msg := <-voicemail
			fmt.Println("Friend heard:", msg)
		}
	}()

	time.Sleep(4 * time.Second) // give enough time for goroutines

}

func BufferedChannelDeadLock() {
	// Create a buffered channel with a capacity of 2.
	ss := make(chan string, 2)
	ss <- "Scaler"
	ss <- "Golang channels"
	ss <- "hell" // deadlock
	fmt.Println(<-ss)
	fmt.Println(<-ss)
	fmt.Println(<-ss)
}
