package snailmath

import "testing"

func TestMag(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"[[1,2],[[3,4],5]]", 143},
		{"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", 1384},
		{"[[[[1,1],[2,2]],[3,3]],[4,4]]", 445},
		{"[[[[3,0],[5,3]],[4,4]],[5,5]]", 791},
		{"[[[[5,0],[7,4]],[5,5]],[6,6]]", 1137},
		{"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", 3488},
	}

	for _, test := range tests {
		n, err := Parse(test.input)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if got := n.Mag(); got != test.want {
			t.Errorf("Mag(%s) = %d, want %d", test.input, got, test.want)
		}
	}
}
