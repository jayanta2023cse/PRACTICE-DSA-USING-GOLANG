package programs

import (
	"log"
	"sync"
)

func RaceConditionWithoutMutes() {
	var count int = 0
	var wg sync.WaitGroup

	for range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()

			count++
		}()
	}

	wg.Wait()
	log.Println("Final Count", count)
}

func RaceConditionWithMutes() {
	var count int = 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	for range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()

			mu.Lock()
			count++
			mu.Unlock()
		}()
	}

	wg.Wait()
	log.Println("Final Count", count)
}

// A mutex (mutual exclusion lock) ensures that only one goroutine at a time can enter the critical section (the part where count is updated).
// When one goroutine calls mu.Lock(), others must wait until mu.Unlock() is called.
// This guarantees count++ executes atomically (no interleaving of reads/writes).
// Result: At the end, you’ll always see Final Count 100.

// ✅ Summary
// Without mutex → race condition, unpredictable results (final count < 100).
// With mutex → safe concurrency, deterministic result (final count = 100).
