package main

import (
	"testing"
)

func TestValidRollDice(t *testing.T) {
	arrOfInt := []int{1, 2, 3, 4, 5, 6}
	result := rollDice(arrOfInt)
	if result < 1 || result > 6 {
		t.Errorf("result must be integer between 1 <= x <= 6, expected : %v, got : %v", arrOfInt, result)
	} else {
		t.Logf("success test rollDice(). expected : %v, got : %v", arrOfInt, result)
	}
}
