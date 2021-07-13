package main

import (
	"testing"
	"mathapp/controllers"
)

func TestSum(t *testing.T) {
	if controllers.Add(2, 5) != 7 {
		t.Fail()
	}
	if controllers.Add(2, 100) != 102 {
		t.Fail()
	}
	if controllers.Add(222, 100) != 322 {
		t.Fail()
	}
}

func TestProduct(t *testing.T) {
	if controllers.Multiply(2, 5) != 10 {
		t.Fail()
	}
	if controllers.Multiply(2, 100) != 200 {
		t.Fail()
	}
	if controllers.Multiply(222, 3) != 666 {
		t.Fail()
	}
}