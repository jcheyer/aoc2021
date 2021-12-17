package packetdecoder

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/jcheyer/aoc2021/lib"
)

type packet struct {
	version      int
	typeID       int
	literalValue int
	sub          []packet
}

type decoder struct {
	Bitstream string
}

func New(s string) *decoder {
	bs, err := raw2Bin(s)
	if err != nil {
		panic(err)
	}
	d := &decoder{
		Bitstream: bs,
	}
	return d
}

func (d *decoder) Parse(pos int) (packet, int) {
	_, t := d.parseHeader(pos)
	if t == 4 {
		return d.parseLiteral(pos)
	}
	return d.parseOperator(pos)
}

func (d *decoder) parseHeader(pos int) (int, int) {
	version, _ := strconv.ParseInt(d.Bitstream[pos:pos+3], 2, 32)

	typeID, _ := strconv.ParseInt(d.Bitstream[pos+3:pos+6], 2, 32)

	return int(version), int(typeID)
}

func (d *decoder) parseLiteral(pos int) (packet, int) {
	p := packet{}
	p.version, p.typeID = d.parseHeader(pos)

	i := pos + 6
	value := ""
	count := 0
	for {
		count += 5

		c := d.Bitstream[i : i+5]
		value += c[1:5]
		if c[0] == byte('0') {
			break
		}
		i += 5
	}

	lv, _ := strconv.ParseInt(value, 2, 32)
	p.literalValue = int(lv)

	len := 6 + count

	return p, len
}

func (d *decoder) parseOperator(pos int) (packet, int) {

	start := pos
	p := packet{}
	p.version, p.typeID = d.parseHeader(pos)

	switch d.Bitstream[pos+6] {
	case '0':
		pos += 7
		total, _ := strconv.ParseInt(d.Bitstream[pos:pos+15], 2, 0)
		pos += 15

		p.sub = make([]packet, 0)
		for total > 0 {

			sub, len := d.Parse(pos)

			p.sub = append(p.sub, sub)

			total -= int64(len)

			pos += len

		}
		len := pos - start

		return p, len
	case '1':
		pos += 7
		n, _ := strconv.ParseInt(d.Bitstream[pos:pos+11], 2, 0)
		pos += 11

		packets := make([]packet, 0)
		for i := 0; i < int(n); i++ {
			sp, len := d.Parse(pos)
			packets = append(packets, sp)
			pos += len
		}
		p.sub = packets
		len := pos - start

		return p, len
	}
	panic("woong length type ID")
}

func raw2Bin(s string) (string, error) {
	sb := strings.Builder{}
	for _, r := range s {
		value, err := strconv.ParseInt(string(r), 16, 64)
		if err != nil {
			return "", err
		}
		sb.WriteString(fmt.Sprintf("%04b", value))

	}

	return sb.String(), nil
}

func (p *packet) Versions() int {
	result := p.version
	for _, sub := range p.sub {
		result += sub.Versions()
	}
	return result
}

func (p *packet) Eval() int64 {
	switch p.typeID {
	case 0:
		sum := int64(0)
		for _, s := range p.sub {
			sum += s.Eval()
		}
		return sum
	case 1:
		prod := int64(1)
		for _, s := range p.sub {
			prod *= s.Eval()
		}
		return prod
	case 2:
		min := int64(math.MaxInt64)
		for _, s := range p.sub {
			min = lib.LowInt64(min, s.Eval())
		}
		return min
	case 3:
		max := int64(0)
		for _, s := range p.sub {
			max = lib.HighInt64(max, s.Eval())
		}
		return max
	case 4:
		return int64(p.literalValue)
	case 5:
		if p.sub[0].Eval() > p.sub[1].Eval() {
			return 1
		}
		return 0
	case 6:
		if p.sub[0].Eval() < p.sub[1].Eval() {
			return 1
		}
		return 0
	case 7:
		if p.sub[0].Eval() == p.sub[1].Eval() {
			return 1
		}
		return 0
	}

	panic("eval error")
	return 0
}
