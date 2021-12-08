package sevensegments

import (
	"strings"
)

type SevenSegment struct {
	rawInput   string
	digits     []string
	value      []string
	reverseMap map[int]string
	forwardMap map[string]int
	pools      map[int][]string
}

func New(s string) *SevenSegment {
	seg := SevenSegment{
		rawInput:   s,
		digits:     make([]string, 0),
		reverseMap: make(map[int]string),
		forwardMap: make(map[string]int),
		pools:      make(map[int][]string),
	}

	parts := strings.Split(s, "|")

	seg.digits = strings.Fields(parts[0])
	seg.value = strings.Fields(parts[1])

	for _, digit := range seg.digits {
		seg.pools[len(digit)] = append(seg.pools[len(digit)], digit)
	}

	return &seg
}

func (seg *SevenSegment) Solve() int {

	// solve the 1:1 pools
	seg.reverseMap[1] = seg.pools[2][0]
	seg.reverseMap[4] = seg.pools[4][0]
	seg.reverseMap[7] = seg.pools[3][0]
	seg.reverseMap[8] = seg.pools[7][0]

	//
	// solve 6 (segment pool (0 6 9))
	//

	// 9 includes the 4
	remainingFromPool6 := make([]string, 0)

	for _, digit := range seg.pools[6] {
		if containsDigitSegments(digit, seg.reverseMap[4]) {
			seg.reverseMap[9] = digit
		} else {
			remainingFromPool6 = append(remainingFromPool6, digit)
		}
	}

	// 0 includes 7, the other on is 6
	for _, digit := range remainingFromPool6 {
		if containsDigitSegments(digit, seg.reverseMap[7]) {
			seg.reverseMap[0] = digit
		} else {
			seg.reverseMap[6] = digit
		}
	}

	//
	// solve 5 (segment pool (2 3 5))
	//

	remainingFromPool5 := make([]string, 0)

	// 3 includes the 7

	for _, digit := range seg.pools[5] {
		if containsDigitSegments(digit, seg.reverseMap[7]) {
			seg.reverseMap[3] = digit
		} else {
			remainingFromPool5 = append(remainingFromPool5, digit)
		}
	}

	// 5 is fully in 6, the other is 2
	for _, digit := range remainingFromPool5 {
		if containsDigitSegments(seg.reverseMap[6], digit) {
			seg.reverseMap[5] = digit
		} else {
			seg.reverseMap[2] = digit
		}
	}

	for k, v := range seg.reverseMap {
		seg.forwardMap[sort(v)] = k
	}

	result := 0
	for _, digit := range seg.value {
		result = result*10 + seg.forwardMap[sort(digit)]
	}

	return result
}

func sort(s string) string {
	var result string
	for _, r := range []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'} {
		if strings.ContainsRune(s, r) {
			result = result + string(r)
		}
	}
	return result
}

func strDiff(si1, si2 string) string {
	var result string
	s1 := si1
	s2 := si2
	if len(si1) < len(si2) {
		s2 = si1
		s1 = si2
	}

	for _, r := range s1 {
		if !strings.ContainsRune(s2, r) {
			result = result + string(r)
		}
	}
	return result
}

func containsDigitSegments(toCheck, includes string) bool {
	for _, r := range includes {
		if !strings.ContainsRune(toCheck, r) {
			return false
		}
	}
	return true
}
