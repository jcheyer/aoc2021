package polymer

import (
	"regexp"
	"strings"

	"github.com/jcheyer/aoc2021/lib"
)

type polymerizer struct {
	Base  string
	rules map[string]string
}

func New(data []string) *polymerizer {
	p := polymerizer{
		rules: make(map[string]string),
	}

	regRules := regexp.MustCompile(`^([A-Z][A-Z]) -> (.?)$`)

	p.Base = data[0]

	for _, line := range data[1:] {
		if match := regRules.FindStringSubmatch(line); match != nil {
			p.rules[match[1]] = match[2]
		}

	}

	return &p
}

func (p *polymerizer) Do(s string) string {

	sb := strings.Builder{}
	l := len(s)

	for i := 0; i < l-1; i++ {
		sb.WriteByte(s[i])
		sb.WriteString(p.rules[s[i:i+2]])

	}
	sb.WriteByte(s[l-1])

	return sb.String()
}

func ComponentCount(s string) map[rune]int {
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}

	return m
}

func HighLow(m map[rune]int) (int, int) {
	high := 0
	low := lib.MaxInt
	for r := range m {
		high = lib.HighInt(high, m[r])
		low = lib.LowInt(low, m[r])
	}

	return high, low

}
