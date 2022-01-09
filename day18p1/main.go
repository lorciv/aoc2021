package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type NumberType int

const (
	TypeError NumberType = iota
	TypeValue
	TypePair
)

type Number struct {
	Type        NumberType
	Parent      *Number
	Val         int     // only for values
	Left, Right *Number // only for pairs
}

func (n *Number) Mag() int {
	switch n.Type {
	case TypeValue:
		return n.Val
	case TypePair:
		return 3*n.Left.Mag() + 2*n.Right.Mag()
	}
	return 0
}

func (n *Number) String() string {
	switch n.Type {
	case TypeValue:
		return strconv.Itoa(n.Val)
	case TypePair:
		return fmt.Sprintf("[%s,%s]", n.Left, n.Right)
	}
	return "err"
}

var input string

func peek() string {
	return input[:1]
}

func shift(n int) string {
	head := input[:n]
	input = input[n:]
	return head
}

func MustParse(s string) *Number {
	n, err := Parse(s)
	if err != nil {
		panic(err)
	}
	return n
}

func Parse(s string) (*Number, error) {
	input = s
	return parse(nil)
}

func parse(parent *Number) (*Number, error) {
	if peek() == "[" {
		return parsePair(parent)
	}
	return parseValue(parent)
}

func parsePair(parent *Number) (*Number, error) {
	p := &Number{
		Type:   TypePair,
		Parent: parent,
	}

	if shift(1) != "[" {
		return nil, errors.New("could not parse pair: missing [")
	}

	var err error
	p.Left, err = parse(p)
	if err != nil {
		return nil, err
	}

	if shift(1) != "," {
		return nil, errors.New("could not parse pair: missing ,")
	}

	p.Right, err = parse(p)
	if err != nil {
		return nil, err
	}
	if shift(1) != "]" {
		return nil, errors.New("could not parse pair: missing ]")
	}

	return p, nil
}

func parseValue(parent *Number) (*Number, error) {
	buf := strings.Builder{}
	for len(input) > 0 && strings.Contains("0123456789", peek()) {
		buf.WriteString(shift(1))
	}

	val, err := strconv.Atoi(buf.String())
	if err != nil {
		return nil, fmt.Errorf("could not parse number: %v", err)
	}
	return &Number{
		Type:   TypeValue,
		Parent: parent,
		Val:    val,
	}, nil
}

func Add(a, b *Number) *Number {
	pair := &Number{
		Type:  TypePair,
		Left:  a,
		Right: b,
	}
	a.Parent = pair
	b.Parent = pair
	return pair
}

func Split(n *Number) bool {
	if n.Type == TypeValue && n.Val > 9 {
		n.Type = TypePair
		n.Left = &Number{
			Type:   TypeValue,
			Parent: n,
			Val:    n.Val / 2,
		}
		n.Right = &Number{
			Type:   TypeValue,
			Parent: n,
			Val:    n.Val - n.Val/2,
		}
		n.Val = 0
		return true
	}
	if n.Type == TypePair {
		if Split(n.Left) {
			return true
		}
		return Split(n.Right)
	}
	return false
}

func Explode(n *Number) bool {
	return explode(n, 0)
}

func explode(n *Number, depth int) bool {
	if n.Type == TypePair && depth == 4 {
		cur := n
		for cur.Parent != nil && cur == cur.Parent.Left {
			cur = cur.Parent
		}
		if cur.Parent != nil {
			cur = cur.Parent.Left
			for cur.Type != TypeValue {
				cur = cur.Right
			}
			cur.Val += n.Left.Val
		}

		cur = n
		for cur.Parent != nil && cur == cur.Parent.Right {
			cur = cur.Parent
		}
		if cur.Parent != nil {
			cur = cur.Parent.Right
			for cur.Type != TypeValue {
				cur = cur.Left
			}
			cur.Val += n.Right.Val
		}

		n.Type = TypeValue
		n.Val = 0
		n.Left = nil
		n.Right = nil

		return true
	}
	if n.Type == TypePair {
		if explode(n.Left, depth+1) {
			return true
		}
		return explode(n.Right, depth+1)
	}
	return false
}

func Reduce(n *Number) {
	if Explode(n) {
		Reduce(n)
	}
	if Split(n) {
		Reduce(n)
	}
}

func main() {
	var num *Number

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		n, err := Parse(scan.Text())
		if err != nil {
			log.Fatal(err)
		}
		if num == nil {
			num = n
			continue
		}
		num = Add(num, n)
		Reduce(num)
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(num, "mag", num.Mag())
}
