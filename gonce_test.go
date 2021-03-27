package gonce

import (
	"testing"
)

var instance int
var once Once

func getNumber(eq int) int {
	once.Do(func() {
		instance = eq
	})
	return instance
}

func TestOnce(t *testing.T) {
	val1, val2 := getNumber(1), getNumber(2)
	if val1 != instance || val1 != val2 {
		t.Error("Numbers are not equal")
	}
}
