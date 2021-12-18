package snailfishmath

import (
	"fmt"
	"math"
	"strconv"
)

type number struct {
	literal int
	left    *number
	right   *number
	parent  *number
}

func New(s string) *number {
	n, _ := parse(s, 0, nil)
	return n
}

func parse(s string, pos int, parent *number) (*number, int) {
	ret := &number{
		parent: parent,
	}
	if s[pos] == '[' {
		l, npc := parse(s, pos+1, ret)
		r, npc := parse(s, npc+1, ret)
		ret.left = l
		ret.right = r
		return ret, npc + 1
	}
	lit, _ := strconv.Atoi(string(s[pos]))
	ret.literal = lit

	return ret, pos + 1
}

// for testing
func (n *number) String() string {
	if n.left == nil && n.right == nil {
		return strconv.Itoa(n.literal)
	}
	return fmt.Sprintf("[%s,%s]", n.left.String(), n.right.String())
}

func (n *number) Magnitude() int {
	if n.left == nil && n.right == nil {
		return n.literal
	}
	return n.left.Magnitude()*3 + n.right.Magnitude()*2
}

func (n *number) Add(n2 *number) *number {
	sum := &number{
		left:  n,
		right: n2,
		//value: 0,
	}
	sum.left.parent = sum
	sum.right.parent = sum
	changes := 0
	for {
		changes = sum.reduceExplode(sum, 0)
		if changes != 0 {
			continue
		}
		changes = sum.reduceSplit(sum)
		if changes == 0 {
			break
		}
	}
	return sum
}

func (n *number) reduceExplode(head *number, depth int) int {
	if depth == 5 {
		n.parent.explode()
		return 1
	}
	if n.parent == nil && depth != 0 {
		return 0
	}
	if n.left == nil {
		return 0
	}
	changes := n.left.reduceExplode(head, depth+1)
	if n.right == nil || changes != 0 {
		return changes
	}
	return n.right.reduceExplode(head, depth+1)
}

func (n *number) reduceSplit(head *number) int {
	if n.left == nil && n.right == nil {
		if n.literal > 9 {
			n.split()
			return 1
		}
		return 0
	}
	if n.left == nil {
		return 0
	}
	changes := n.left.reduceSplit(head)
	if n.right == nil || changes != 0 {
		return changes
	}

	return n.right.reduceSplit(head)
}

func (n *number) split() {
	n.left = &number{
		left:    nil,
		right:   nil,
		parent:  n,
		literal: int(math.Floor(float64(n.literal) / 2)),
	}
	n.right = &number{
		left:    nil,
		right:   nil,
		parent:  n,
		literal: int(math.Ceil(float64(n.literal) / 2)),
	}
	n.literal = 0
}

func (n *number) explode() {
	left := n.findLeft()
	if left != nil {
		left.literal += n.left.literal
	}
	right := n.findRight()
	if right != nil {
		right.literal += n.right.literal
	}
	n.left.parent = nil
	n.left = nil
	n.right.parent = nil
	n.right = nil
	n.literal = 0
}

func (n *number) findLeft() *number {
	prev := n
	for number := n.parent; number != nil; {
		if number.right == prev {
			prev = number
			number = number.left
			continue
		}
		if number.left == prev {
			if n.parent == nil {
				break
			}
			prev = number
			number = number.parent
			continue
		}
		if number.left == nil && number.right == nil {
			return number
		}
		number = number.right
	}

	return nil
}

func (n *number) findRight() *number {
	prev := n
	for number := n.parent; number != nil; {
		if number.left == prev {
			prev = number
			number = number.right
			continue
		}
		if number.right == prev {
			if n.parent == nil {
				break
			}
			prev = number
			number = number.parent
			continue
		}
		if number.left == nil && number.right == nil {
			return number
		}
		number = number.left
	}

	return nil
}
