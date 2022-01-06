package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type packet interface {
	checksum() int
}

type header struct {
	vers, typ int
}

type literal struct {
	header
	val int
}

func (l literal) checksum() int {
	return l.vers
}

func (l literal) String() string {
	return fmt.Sprintf("lit{%d %d - %d}", l.vers, l.typ, l.val)
}

type operator struct {
	header
	children []packet
}

func (o operator) checksum() int {
	sum := o.vers
	for _, subPkt := range o.children {
		sum += subPkt.checksum()
	}
	return sum
}

func (o operator) String() string {
	return fmt.Sprintf("oper{%d %d - %v}", o.vers, o.typ, o.children)
}

func btoi(s string) (int, error) {
	val, err := strconv.ParseInt(s, 2, 0)
	if err != nil {
		return 0, err
	}
	return int(val), err
}

var bits string

func shift(n int) string {
	head := bits[:n]
	bits = bits[n:]
	return head
}

func parse() (packet, error) {
	head, err := parseHeader()
	if err != nil {
		return nil, fmt.Errorf("could not parse packet: %v", err)
	}

	if head.typ == 4 {
		lit, err := parseLiteral(head)
		if err != nil {
			return nil, fmt.Errorf("could not parse packet: %v", err)
		}
		return lit, nil
	}

	op, err := parseOperator(head)
	if err != nil {
		return nil, fmt.Errorf("could not parse packet: %v", err)
	}
	return op, err
}

func parseOperator(head header) (operator, error) {
	op := operator{
		header: head,
	}

	switch lenTyp := shift(1); lenTyp {
	case "0":
		length, err := btoi(shift(15))
		if err != nil {
			return operator{}, fmt.Errorf("could not parse operator: invalid sub-packet bit length: %v", err)
		}
		// fmt.Println("sub-packet bit length", length)

		initLength := len(bits)
		for {
			subPkt, err := parse()
			if err != nil {
				return operator{}, err
			}
			op.children = append(op.children, subPkt)

			if len(bits) == initLength-length {
				break
			}
		}
	case "1":
		count, err := btoi(shift(11))
		if err != nil {
			return operator{}, fmt.Errorf("could not parse operator: invalid sub-packet count: %v", err)
		}
		// fmt.Println("sub-packet count", count)

		for i := 0; i < count; i++ {
			subPkt, err := parse()
			if err != nil {
				return operator{}, err
			}
			op.children = append(op.children, subPkt)
		}
	default:
		return operator{}, fmt.Errorf("could not parse operator: invalid lenght type id %s", lenTyp)
	}

	return op, nil
}

func parseLiteral(head header) (literal, error) {
	buf := strings.Builder{}
	for shift(1) == "1" {
		buf.WriteString(shift(4))
	}
	buf.WriteString(shift(4))

	val, err := btoi(buf.String())
	if err != nil {
		return literal{}, fmt.Errorf("could not parse literal: %v", err)
	}

	return literal{
		header: head,
		val:    val,
	}, nil
}

func parseHeader() (header, error) {
	vers, err := btoi(shift(3))
	if err != nil {
		return header{}, fmt.Errorf("could not parse header: invalid version: %v", err)
	}
	typ, err := btoi(shift(3))
	if err != nil {
		return header{}, fmt.Errorf("could not parse header: invalid type: %v", err)
	}
	return header{
		vers: vers,
		typ:  typ,
	}, nil
}

func main() {
	buf := bytes.Buffer{}
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanRunes)
	for scan.Scan() {
		val, err := strconv.ParseUint(scan.Text(), 16, 64)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(&buf, "%04b", val)
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}
	bits = buf.String()

	fmt.Println("bits", bits)

	root, err := parse()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(root)

	fmt.Println("checksum", root.checksum())
}
