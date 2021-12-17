package packetdecoder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpecial(t *testing.T) {
	d3 := New("620080001611562C8802118E34")
	p3, _ := d3.Parse(0)
	assert.Equal(t, 12, p3.Versions())
}

func TestVersions(t *testing.T) {

	//D2FE28
	d1 := New("D2FE28")
	p1, _ := d1.Parse(0)
	assert.Equal(t, 6, p1.Versions())

	d2 := New("8A004A801A8002F478")
	p2, _ := d2.Parse(0)
	assert.Equal(t, 16, p2.Versions())

	/*t.Logf("D3 !!!!!!!!!\n")
	d3 := New("620080001611562C8802118E34")
	p3, _ := d3.Parse(0)
	assert.Equal(t, 12, p3.Versions())

	/*d4 := New("C0015000016115A2E0802F182340")
	p4, _ := d4.Parse(0)
	assert.Equal(t, 23, p4.Versions())*/

	d5 := New("A0016C880162017C3686B18A3D4780")
	p5, _ := d5.Parse(0)
	assert.Equal(t, 31, p5.Versions())
}

func TestRaw2Binary(t *testing.T) {
	assert.Equal(t, shim("0001", nil), shim(raw2Bin("1")))
	assert.Equal(t, shim("1010", nil), shim(raw2Bin("A")))
	assert.Equal(t, shim("1111", nil), shim(raw2Bin("F")))
}

func TestDecoderLiteral(t *testing.T) {
	d := New("D2FE28")
	assert.Equal(t, shim(6, 4), shim(d.parseHeader(0)))
	//assert.Len(t, d.packets, 0)
	//assert.Len(t, d.Bitstream, 21)
	p, parsed := d.Parse(0)
	assert.Equal(t, 21, parsed)
	//assert.Len(t, d.packets, 1)
	expectedPacket := packet{
		version:      6,
		typeID:       4,
		literalValue: 2021,
	}
	assert.Equal(t, expectedPacket, p)
	//assert.Equal(t, d.pos, len(d.bitstream))
}

func TestDecoderOperator0(t *testing.T) {
	s := "38006F45291200"
	assert.Len(t, s, 14)
	d := New(s)
	assert.Equal(t, shim(1, 6), shim(d.parseHeader(0)))
	assert.Equal(t, "00111000000000000110111101000101001010010001001000000000", d.Bitstream)
	assert.Len(t, d.Bitstream, 56)
	p, count := d.Parse(0)
	assert.Equal(t, 49, count)
	assert.Len(t, p.sub, 2)
	assert.Equal(t, 10, p.sub[0].literalValue)
	assert.Equal(t, 20, p.sub[1].literalValue)
}

func TestDecoderOperator1(t *testing.T) {
	d := New("EE00D40C823060")
	assert.Equal(t, shim(7, 3), shim(d.parseHeader(0)))
	assert.Equal(t, "11101110000000001101010000001100100000100011000001100000", d.Bitstream)
	//assert.Len(t, d.Bitstream, 56)
	p, c := d.Parse(0)
	assert.Equal(t, 51, c)

	assert.Len(t, p.sub, 3)
	assert.Equal(t, 1, p.sub[0].literalValue)
	assert.Equal(t, 2, p.sub[1].literalValue)
	assert.Equal(t, 3, p.sub[2].literalValue)
}

func TestEval(t *testing.T) {
	d := New("C200B40A82")
	p, _ := d.Parse(0)
	assert.Equal(t, int64(3), p.Eval())

	d = New("04005AC33890")
	p, _ = d.Parse(0)
	assert.Equal(t, int64(54), p.Eval())

	d = New("880086C3E88112")
	p, _ = d.Parse(0)
	assert.Equal(t, int64(7), p.Eval())

	d = New("CE00C43D881120")
	p, _ = d.Parse(0)
	assert.Equal(t, int64(9), p.Eval())

	d = New("D8005AC2A8F0")
	p, _ = d.Parse(0)
	assert.Equal(t, int64(1), p.Eval())

	d = New("F600BC2D8F")
	p, _ = d.Parse(0)
	assert.Equal(t, int64(0), p.Eval())

	d = New("9C005AC2F8F0")
	p, _ = d.Parse(0)
	assert.Equal(t, int64(0), p.Eval())

	d = New("9C0141080250320F1802104A08")
	p, _ = d.Parse(0)
	assert.Equal(t, int64(1), p.Eval())

}

func shim(a, b interface{}) []interface{} {
	return []interface{}{a, b}
}
