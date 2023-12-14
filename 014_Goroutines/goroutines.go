package goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func GoroutinesTest() {
	fmt.Println("Starting Goroutine")

	go fiveTimes("This goroutine will no always end")
	fiveTimes("This message will be always shown")
	fmt.Println("End of the program")
	// When the main goroutine ends, all the goroutines ends, even if they have not started

	// You can control the excecution stoping the main goroutine using a WaitGroup
	// The main goroutine will continue when the inner counter of the WaitGroup were 0
	waitGroupTest()

	waitGroupRaceCondition()

	mutexNoRaceCondition()
}

func fiveTimes(msg string) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d of 5\t%s\n", i, msg)		
	}
}

// Can Produce Race Conditions
func waitGroupTest(){
	wg := sync.WaitGroup{}

	for i := 0; i <= 3; i++{
		job := i
		wg.Add(1)
		go func ()  {
			defer wg.Done()
			fmt.Println("Job Number:", job)	
		}()
	}

	wg.Wait()
	fmt.Println("All jobs are done")
}

func add(nums []int) int{
	result := 0

	for _, val := range nums{
		result += val
	}

	return result
}

func waitGroupRaceCondition(){
	maxTasks := 4
	v := []int{0,1,3,1,0,7,8,9,3,3,0,2}

	wg := sync.WaitGroup{}
	wg.Add(maxTasks)

	total := 0

	for job := 0; job < maxTasks; job++ {
		i := job

		go func ()  {
			defer wg.Done()
			start := i * len(v)/maxTasks
			end := (i + 1) * len(v)/maxTasks
			result := add(v[start:end])
			total += result
		}()
	}

	wg.Wait()

	if (total != 37){
		fmt.Println("Race condition:", total)
	}else {
		fmt.Println("No race condition:", total)
	}
}

// We can avoid race Conditions locking the goroutines to do a certain action
// lock goroutines its posible using sync.Mutex

func mutexNoRaceCondition(){
	maxTasks := 4
	v := []int{0,1,3,1,0,7,8,9,3,3,0,2}

	mt := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(maxTasks)

	total := 0

	for job := 0; job < maxTasks; job++ {
		i := job

		go func ()  {
			defer wg.Done()
			start := i * len(v)/maxTasks
			end := (i + 1) * len(v)/maxTasks
			result := add(v[start:end])

			// Locking goroutines
			mt.Lock()
			total += result
			// After we lock the gorutines, is important to unlock
			mt.Unlock()
		}()
	}

	wg.Wait()

	if (total != 37){
		fmt.Println("Race condition:", total)
	}else {
		fmt.Println("No race condition:", total)
	}
}

// ? If the used value will be Read Only, its better to use RWMutex, that contains RLock and RUnlock
// ? to allow more goroutines to read a certain data


// Other way is using Atomic, that provides certain funtions to do some kind of operations
func atomicNoRaceCondition(){
	maxTasks := 4
	v := []int{0,1,3,1,0,7,8,9,3,3,0,2}

	wg := sync.WaitGroup{}
	wg.Add(maxTasks)

	total := int64(0)

	for job := 0; job < maxTasks; job++ {
		i := job

		go func ()  {
			defer wg.Done()
			start := i * len(v)/maxTasks
			end := (i + 1) * len(v)/maxTasks
			result := int64( add(v[start:end]) )

			atomic.AddInt64(&total, result)
		}()
	}

	wg.Wait()

	if (total != 37){
		fmt.Println("Race condition:", total)
	}else {
		fmt.Println("No race condition:", total)
	}
}