package pkg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStringify(t *testing.T) {
	tests := map[string]struct {
		input  interface{}
		output string
	}{
		"nil": {
			input:  nil,
			output: "nil",
		},
		"integer": {
			input:  1534,
			output: "1534",
		},
		"float": {
			input:  23.5246,
			output: "23.5246",
		},
		"boolean": {
			input:  true,
			output: "true",
		},
		"array of values": {
			input:  [5]int{1, 2, 3, 4, 5},
			output: "[5]int[1, 2, 3, 4, 5]",
		},
		"slice": {
			input:  []int{1, 2, 3, 4, 5},
			output: "[]int[1, 2, 3, 4, 5]",
		},
		"map": {
			input:  map[string]int{"a": 1, "b": 2, "c": 3},
			output: "{\"a\": 1, \"b\": 2, \"c\": 3}",
		},
		"struct": {
			input:  struct{ A int64 }{A: int64(1)},
			output: "struct{A: 1}",
		},
		"pointer": {
			input:  new(int),
			output: "*int(0)",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			require.Equal(t, test.output, Stringify(test.input))
		})
	}
}
