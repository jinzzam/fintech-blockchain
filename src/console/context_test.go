package console

import (
	"reflect"
	"testing"
)

func TestNewBlockchain(t *testing.T) {
	tests := []struct {
		f        interface{}
		params   []interface{}
		expected interface{}
	}{
		{
			f:        func(a, b int) int { return a + b },
			params:   []interface{}{1, 2},
			expected: 3,
		},
		{
			f:        func(a, b, c int) int { return a + b + c },
			params:   []interface{}{1, 2, 3},
			expected: 6,
		},
		{
			f:        func() int { return 0 },
			params:   []interface{}{},
			expected: 0,
		},
		{
			f:        func() {},
			params:   []interface{}{},
			expected: nil,
		},
		{
			f:        func() {},
			params:   []interface{}{1, 2},
			expected: nil,
		},
	}

	for _, test := range tests {
		ctx := Context{
			newBlockchain: reflect.ValueOf(test.f),
		}

		actual := ctx.NewBlockchain(test.params...)

		if actual != test.expected {
			t.Errorf("TestNewBlockchain(%v): expected %v, actual %v\n", test.params, test.expected, actual)
		}
	}
}

func TestNewBlock(t *testing.T) {
	tests := []struct {
		f        interface{}
		params   []interface{}
		expected interface{}
	}{
		{
			f:        func(a, b int) int { return a + b },
			params:   []interface{}{1, 2},
			expected: 3,
		},
		{
			f:        func(a, b, c int) int { return a + b + c },
			params:   []interface{}{1, 2, 3},
			expected: 6,
		},
		{
			f:        func() int { return 0 },
			params:   []interface{}{},
			expected: 0,
		},
		{
			f:        func() {},
			params:   []interface{}{},
			expected: nil,
		},
		{
			f:        func() {},
			params:   []interface{}{1, 2},
			expected: nil,
		},
	}

	for _, test := range tests {
		ctx := Context{
			newBlock: reflect.ValueOf(test.f),
		}

		actual := ctx.NewBlock(test.params...)

		if actual != test.expected {
			t.Errorf("TestNewBlock(%v): expected %v, actual %v\n", test.params, test.expected, actual)
		}
	}
}
