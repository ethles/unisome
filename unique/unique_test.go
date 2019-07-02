package unique

import (
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSimplePseudoEncrypt(t *testing.T) {
	Convey("test determinacy", t, func() {
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
		So(errTimes, ShouldEqual, 0)
	})

}

func TestStaplePseudoEncrypt(t *testing.T) {
	Convey("test figures number", t, func() {
		inputValue := 1
		inputLenthList := []int{1, 2, 3}
		errTimes := 0
		for _, l := range inputLenthList {
			output := StaplePseudoEncrypt(inputValue, l)
			if len(strconv.Itoa(output)) != l {
				errTimes = errTimes + 1
			}
		}
		So(errTimes, ShouldEqual, 0)
	})

	Convey("test invalid input", t, func() {
		inputValue := 1
		invalidInputLenthList := []int{-1, 0}
		errTimes := 0
		for _, l := range invalidInputLenthList {
			output := StaplePseudoEncrypt(inputValue, l)
			if output != 0 {
				errTimes = errTimes + 1
			}
		}
		So(errTimes, ShouldEqual, 0)
	})
}
