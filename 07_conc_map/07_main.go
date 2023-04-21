package main

import (
	"fmt"
	"os"
	"sync"
)

type sMap struct {
	m   map[string]string
	mtx *sync.RWMutex
}

func main() {
	var key, value string

	sm := &sMap{m: make(map[string]string), mtx: &sync.RWMutex{}}

	for {
		fmt.Fscanln(os.Stdin, &key, &value)
		go sm.putInMap(key, value)
	}
}

func (sm *sMap) putInMap(key, value string) {
	sm.mtx.RLock()
	sm.m[key] = value
	sm.mtx.RUnlock()
}
