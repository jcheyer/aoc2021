package polymer

import (
	"regexp"
	"strings"

	"github.com/jcheyer/aoc2021/lib"
)

type polymerizer struct {
	poly  map[string]int
	Base  string
	Rules map[string]string
}

func New(data []string) *polymerizer {
	p := polymerizer{
		Rules: make(map[string]string),
		poly:  make(map[string]int),
	}

	regRules := regexp.MustCompile(`^([A-Z][A-Z]) -> (.?)$`)

	p.Base = data[0]

	for _, line := range data[1:] {
		if match := regRules.FindStringSubmatch(line); match != nil {
			p.Rules[match[1]] = match[2]
		}

	}
	p.poly = p.stringToPoly(p.Base)

	return &p
}

func (p *polymerizer) stringToPoly(s string) map[string]int {
	m := make(map[string]int)
	l := len(s)
	for i := 0; i < l-1; i++ {
		m[s[i:i+2]]++
	}
	return m
}

func (p *polymerizer) Do(c int) (int, int) {
	for i := 0; i < c; i++ {
		p.do()
	}

	m := map[rune]int{}
	for k, v := range p.poly {
		m[rune(k[0])] += v
		m[rune(k[1])] += v
	}
	// First and Last rune of Base are not double counted, so add them
	m[rune(p.Base[0])]++
	m[rune(p.Base[len(p.Base)-1])]++
	h, l := HighLow(m)
	return h / 2, l / 2
}

func (p *polymerizer) do() {
	newPoly := make(map[string]int)
	for poly, c := range p.poly {
		newPoly[string(poly[0])+p.Rules[poly]] += c
		newPoly[p.Rules[poly]+string(poly[1])] += c
	}
	p.poly = newPoly
}

func (p *polymerizer) DontDo(s string) string {

	sb := strings.Builder{}
	l := len(s)

	for i := 0; i < l-1; i++ {
		sb.WriteByte(s[i])
		sb.WriteString(p.Rules[s[i:i+2]])

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
