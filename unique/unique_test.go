package unique

import (
	"strconv"
	"testing"
)

func TestSimplePseudoEncrypt(t *testing.T) {

	// test determinacy
	input := []int{-3, -2, -1, 0, 1, 2, 3}
	output1 := []int{}
	output2 := []int{}
	for _, n := range input {
		output1 = append(output1, SimplePseudoEncrypt(n))
	}
	for _, n := range input {
		output2 = append(output2, SimplePseudoEncrypt(n))
	}
	errTimes := 0
	for i := 0; i < len(output1); i++ {
		if output1[i] != output2[i] {
			errTimes = errTimes + 1
		}
	}
	if errTimes > 0 {
		t.Error(`determinacy test fail`)
	}

}

func TestStaplePseudoEncrypt(t *testing.T) {

	// test figures number
	inputValue := 1
	inputLenthList := []int{1, 2, 3}
	errTimes := 0
	for _, l := range inputLenthList {
		output := StaplePseudoEncrypt(inputValue, l)
		if len(strconv.Itoa(output)) != l {
			errTimes = errTimes + 1
		}
	}
	if errTimes > 0 {
		t.Error(`figures number test fail`)
	}

	// test invalid input
	inputValue = 1
	invalidInputLenthList := []int{-1, 0}
	errTimes = 0
	for _, l := range invalidInputLenthList {
		output := StaplePseudoEncrypt(inputValue, l)
		if output != 0 {
			errTimes = errTimes + 1
		}
	}
	if errTimes > 0 {
		t.Error(`invalid input test fail`)
	}
}
