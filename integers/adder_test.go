package main

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	expect := 5

	if sum != expect {
		t.Errorf("Error: Expected %q, Got %q", expect, sum)
	}

}
