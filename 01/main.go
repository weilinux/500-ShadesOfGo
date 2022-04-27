package main

import "sync"

/*
Type Declarations and Methods
*/

//error
type myMutex sync.Mutex

//inherit from base struct: embedding the original type as an anonymous field
type myLocker struct {
	sync.Mutex
}


func main() {
	var mtx sync.Mutex
	mtx.Lock()
	mtx.Unlock()

/*
	var mtx2 myMutex
	myMutex.Lock()
	myMutex.Unlock()
*/

	var lock myLocker
	lock.Lock()
	lock.Unlock()
}
