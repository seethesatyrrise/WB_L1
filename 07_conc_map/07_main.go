package main

import (
	"fmt"
	"os"
	"sync"
)

type sMap struct {
	m   map[string]string
	mtx *sync.Mutex
}

func main() {
	var key, value string

	sm := &sMap{m: make(map[string]string), mtx: &sync.Mutex{}}

	for {
		fmt.Fscanln(os.Stdin, &key, &value)
		go sm.putInMap(key, value)
	}
}

func (sm *sMap) putInMap(key, value string) {
	sm.mtx.Lock()
	sm.m[key] = value
	sm.mtx.Unlock()
}
