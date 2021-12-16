package packetdecoder

import (
	"fmt"
	"strconv"
	"strings"
)

type packet struct {
	version      int
	typeID       int
	literalValue int
	sub          []packet
}

type decoder struct {
	Bitstream string
	//pos       int
	//versionSum int

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

	fmt.Printf("Literal Packet at pos %+v version: %+v type: %+v\n", pos, p.version, p.typeID)

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
	//padding := count % 4
	//len := 6 + count + padding - 1
	//d.pos = pos + len
	//d.packets = append(d.packets, p)
	//d.versionSum += p.version

	return p, len
}

func (d *decoder) parseOperator(pos int) (packet, int) {

	start := pos
	p := packet{}
	p.version, p.typeID = d.parseHeader(pos)
	fmt.Printf("Operator Packet at pos %+v version: %+v type: %+v\n", pos, p.version, p.typeID)
	switch d.Bitstream[pos+6] {
	case '0':
		pos += 7
		total, _ := strconv.ParseInt(d.Bitstream[pos:pos+15], 2, 0)
		pos += 15

		fmt.Printf("length mode start: %+v total: %+v\n", start, total)

		//packets := make([]packet, 0)
		p.sub = make([]packet, 0)
		for total > 0 {

			sub, len := d.Parse(pos)

			p.sub = append(p.sub, sub)

			total -= int64(len)

			pos += len
			//if d.isOnlyZero(pos, int(total)) {
			// 	return p, pos - start + int(total)
			//}

		}
		len := pos - start
		//padding := len % 8
		//panic("this is bad")

		return p, len
	case '1':
		pos += 7
		n, _ := strconv.ParseInt(d.Bitstream[pos:pos+11], 2, 0)
		pos += 11

		//fmt.Printf("packet count mode: start: %+v n: %+v\n", start, n)

		packets := make([]packet, 0)
		for i := 0; i < int(n); i++ {
			sp, len := d.Parse(pos)
			packets = append(packets, sp)
			pos += len
		}
		p.sub = packets
		len := pos - start
		//padding := len % 8
		//return p, len + 8 - padding
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

func (d *decoder) isOnlyZero(pos, len int) bool {
	fmt.Printf("comp called !!! pos: %+v len: %+v\n", pos, len)
	comp := strings.Repeat("0", len)
	fmt.Printf("comp := %s bitstreamsub: %s isZero: %+v\n", comp, d.Bitstream[pos:pos+len], d.Bitstream[pos:pos+len] == comp)
	return d.Bitstream[pos:pos+len] == comp
}
