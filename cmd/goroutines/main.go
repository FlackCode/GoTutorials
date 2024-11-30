package main

import (
	"fmt"
	"time"
	"sync"
)

var mutex = sync.RWMutex{} //mutual exclusion
var waitGroup = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i:=0; i<len(dbData); i++ {
		waitGroup.Add(1)
		go dbCall(i)
	}
	waitGroup.Wait()
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\nThe results are %v", results)
}

func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay)*time.Microsecond)
	save(dbData[i])
	log()
	waitGroup.Done()
}

func save(result string) {
	mutex.Lock()
	results = append(results, result)
	mutex.Unlock()
}

func log() {
	mutex.RLock()
	fmt.Printf("\nThe current results are: %v", results)
	mutex.RUnlock()
}
