package unique

import "math"

func SimplePseudoEncrypt(val int) int {
	var l1 int = (val >> 16) & 65535
	var r1 int = val & 65535
	var l2 int
	var r2 int
	var i int = 0

	for i < 3 {
		l2 = r1
		r2 = l1 ^ int((float64((1366*r1+150889)%714025)/714025.0)*32767)
		l1 = l2
		r1 = r2
		i = i + 1
	}
	return ((r1 << 16) + l1)
}

func StaplePseudoEncrypt(val int, len int) int {
	if len <= 0 {
		return 0
	}
	origin := SimplePseudoEncrypt(val)
	divisor := int(math.Pow10(len))
	reminder := origin % divisor
	return reminder
}
