package unique

import (
	"fmt"
	"math"
)

var adjectiveSet []string = []string{"able", "active", "adaptable", "adroit", "aggressive"}
var statusSet []string = []string{"Duke", "Marquis", "Count", "Viscount", "Baron", "Lord"}
var nameSet []string = []string{"jack", "tom", "tim"}
var cartesian int = 0

func init() {
	//adjectiveSet = config.adjective
	//statusSet = config.status
	//nameSet = config.name
	cartesian = len(adjectiveSet) * len(statusSet) * len(nameSet)
}

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

func GenerateFullName(val int) string {

	generateKey := SimplePseudoEncrypt(val) % cartesian
	nameChoice := generateKey / (len(adjectiveSet) * len(statusSet))
	statusChoice := (generateKey % (len(adjectiveSet) * len(statusSet))) / len(adjectiveSet)
	adjectiveChoice := (generateKey % (len(adjectiveSet) * len(statusSet))) % len(adjectiveSet)
	fullName := fmt.Sprintf("%s %s %s", adjectiveSet[adjectiveChoice], statusSet[statusChoice], nameSet[nameChoice])
	return fullName
}
