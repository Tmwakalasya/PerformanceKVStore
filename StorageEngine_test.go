package main

import (
	"sync"
	"testing"
)

func TestGet(t *testing.T) {
	dataStore := &DataStore{data: make(map[string]string)}
	dataStore.set("John", "Doe")

	got, ok := dataStore.Get("John")
	if !ok {
		t.Errorf("get() returned false, want true")
	}

	if got != "Doe" {
		t.Errorf("get() = %v, want %v", got, "Doe")
	}
}

func TestSet(t *testing.T) {
	dataStore := &DataStore{data: make(map[string]string)}
	dataStore.set("John", "Doe")
	got, ok := dataStore.Get("John")
	if !ok {
		t.Errorf("get() returned false, want true")
	}
	if got != "Doe" {
		t.Errorf("get() = %v, want %v", got, "Doe")
	}
}

func TestRemoveKey(t *testing.T) {
	dataStore := &DataStore{data: make(map[string]string)}
	dataStore.set("John", "Doe")

	removedValue, ok := dataStore.RemoveKey("John")
	if !ok {
		t.Errorf("remove() returned false, want true for an existing key")
	}

	got, ok := dataStore.Get("John")
	if ok {
		t.Errorf("get() returned true, want false")
	}
	if got != "" {
		t.Errorf("get() returned non-nil value, want nil")
	}

	if removedValue != "Doe" {
		t.Errorf("remove() did not return the correct deleted value")
	}
}

func TestConcurrency(t *testing.T) {
	var wg sync.WaitGroup
	dataStore := &DataStore{data: make(map[string]string)}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			dataStore.Get("key")
			//dataStore.set("key", "value")
			dataStore.RemoveKey("key")
		}()
	}

	wg.Wait()

	// Validate the results of the concurrent operations
	// ...
}
