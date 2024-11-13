package main

import "testing"

func TestGet(t *testing.T) {

	set("John", "Doe")

	got, ok := get("John")

	if !ok {
		t.Errorf("get() returned false, want true")
	}

	if got.value != "Doe" {
		t.Errorf("get() = %v, want %v", got.value, "Doe")
	}
}

func TestSet(t *testing.T) {
	set("John", "Doe")
	got, ok := get("John")
	if !ok {
		t.Errorf("get() returned false, want true")
	}
	if got.value != "Doe" {
		t.Errorf("get() = %v, want %v", got.value, "Doe")
	}

}
func TestRemoveKey(t *testing.T) {
	set("John", "Doe")

	removedValue, ok := RemoveKey("John")
	if !ok {
		t.Errorf("remove() returned false, want true for an existing key")
	}

	got, ok := get("John")
	if ok {
		t.Errorf("get() returned true, want false")
	}
	if got != nil {
		t.Errorf("get() returned non-nil value, want nil")
	}

	if removedValue == nil || removedValue.value != "Doe" {
		t.Errorf("remove() did not return the correct deleted value")
	}
}
