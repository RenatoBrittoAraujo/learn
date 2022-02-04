package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(1, 2)
	if total != 3 {
		t.Errorf("Soma(1,2) = %d; expected 3", total)
	}
}
