package snailfishmath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumber(t *testing.T) {
	n := New("[1,2]")
	expect := &number{
		literal: 0,
		parent:  nil,
	}
	expect.left = &number{
		literal: 1,
		parent:  expect}
	expect.right = &number{
		literal: 2,
		parent:  expect}
	assert.Nil(t, n.parent)
	assert.Equal(t, expect, n)
}

func TestMagnitude(t *testing.T) {
	n := New("[9,1]")
	assert.Equal(t, 29, n.Magnitude())
	n = New("[1,9]")
	assert.Equal(t, 21, n.Magnitude())
	n = New("[[9,1],[1,9]]")
	assert.Equal(t, 129, n.Magnitude())
	n = New("[[1,2],[[3,4],5]]")
	assert.Equal(t, 143, n.Magnitude())
	n = New("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]")
	assert.Equal(t, 1384, n.Magnitude())
	n = New("[[[[1,1],[2,2]],[3,3]],[4,4]]")
	assert.Equal(t, 445, n.Magnitude())
	n = New("[[[[3,0],[5,3]],[4,4]],[5,5]]")
	assert.Equal(t, 791, n.Magnitude())
	n = New("[[[[5,0],[7,4]],[5,5]],[6,6]]")
	assert.Equal(t, 1137, n.Magnitude())
	n = New("[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]")
	assert.Equal(t, 3488, n.Magnitude())
}

func TestAdd(t *testing.T) {
	n1 := New("[1,2]")
	n2 := New("[[3,4],5]")
	expect := New("[[1,2],[[3,4],5]]")
	assert.Equal(t, expect, n1.Add(n2))

	n := New("[1,1")
	n = n.Add(New("[2,2]"))
	assert.Equal(t, "[[1,1],[2,2]]", n.String())
	n = n.Add(New("[3,3]"))
	assert.Equal(t, "[[[1,1],[2,2]],[3,3]]", n.String())
	n = n.Add(New("[4,4]"))
	assert.Equal(t, "[[[[1,1],[2,2]],[3,3]],[4,4]]", n.String())
	n = n.Add(New("[5,5]"))
	assert.Equal(t, "[[[[3,0],[5,3]],[4,4]],[5,5]]", n.String())
	n = n.Add(New("[6,6]"))
	assert.Equal(t, "[[[[5,0],[7,4]],[5,5]],[6,6]]", n.String())

	n = New("[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]")
	n = n.Add(New("[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]"))
	assert.Equal(t, "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]", n.String())
	n = n.Add(New("[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]"))
	assert.Equal(t, "[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]", n.String())
	n = n.Add(New("[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]"))
	assert.Equal(t, "[[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]]", n.String())
	n = n.Add(New("[7,[5,[[3,8],[1,4]]]]"))
	assert.Equal(t, "[[[[7,7],[7,8]],[[9,5],[8,7]]],[[[6,8],[0,8]],[[9,9],[9,0]]]]", n.String())
	n = n.Add(New("[[2,[2,2]],[8,[8,1]]]"))
	assert.Equal(t, "[[[[6,6],[6,6]],[[6,0],[6,7]]],[[[7,7],[8,9]],[8,[8,1]]]]", n.String())
	n = n.Add(New("[2,9]"))
	assert.Equal(t, "[[[[6,6],[7,7]],[[0,7],[7,7]]],[[[5,5],[5,6]],9]]", n.String())
	n = n.Add(New("[1,[[[9,3],9],[[9,0],[0,7]]]]"))
	assert.Equal(t, "[[[[7,8],[6,7]],[[6,8],[0,8]]],[[[7,7],[5,0]],[[5,5],[5,6]]]]", n.String())
	n = n.Add(New("[[[5,[7,4]],7],1]"))
	assert.Equal(t, "[[[[7,7],[7,7]],[[8,7],[8,7]]],[[[7,0],[7,7]],9]]", n.String())
	n = n.Add(New("[[[[4,2],2],6],[8,7]]"))
	assert.Equal(t, "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", n.String())

}

func TestString(t *testing.T) {
	s := "[[9,1],[1,9]]"
	n := New(s)
	assert.Equal(t, s, n.String())

	s = "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]"
	n = New(s)
	assert.Equal(t, s, n.String())
}
