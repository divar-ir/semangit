package testUtils

import "testing"

func AssertPanics(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("Expected to encounter a panic. But no panics occurred.")
		}
	}()
	f()
}
