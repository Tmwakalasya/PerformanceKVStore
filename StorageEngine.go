package main

import "sync"

type DataStore struct {
	key   string
	value string
}

var store = make(map[string]*DataStore)
var mu sync.RWMutex

func get(key string) (*DataStore, bool) {
	mu.RLock()
	defer mu.RUnlock()
	data, ok := store[key]
	return data, ok
}

func set(key string, value string) {
	mu.Lock()
	defer mu.Unlock()
	store[key] = &DataStore{key, value}
}

func RemoveKey(key string) (*DataStore, bool) {
	mu.Lock()
	defer mu.Unlock()
	data, ok := store[key]
	if ok {
		delete(store, key)
	}
	return data, ok
}

func update(key string, value string) {
	mu.Lock()
	defer mu.Unlock()
	store[key] = &DataStore{key, value}
}
