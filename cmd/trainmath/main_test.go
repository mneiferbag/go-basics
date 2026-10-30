package main

import "testing"

func TestRandomInts(t *testing.T) {
	for i := 0; i < 100; i++ {
		a, b := randomInts()
		if a < 0 || a >= 100 {
			t.Errorf("randomInts() returned a = %d; want in range [0, 100)", a)
		}
		if b < 0 || b >= 100 {
			t.Errorf("randomInts() returned b = %d; want in range [0, 100)", b)
		}
	}
}
