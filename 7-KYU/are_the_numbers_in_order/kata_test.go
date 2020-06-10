package main

import (
	"reflect"
	"testing"
)

func TestInAscOrder(t *testing.T) {
	tests := map[string]struct {
		name  string
		input []int
		want  bool
	}{
		"first test":  {input: []int{1, 2, 4, 7, 19}, want: true},
		"second test": {input: []int{1, 2, 3, 4, 5}, want: true},
		"third test":  {input: []int{1, 6, 10, 18, 2, 4, 20}, want: false},
		"forth test":  {input: []int{9, 8, 7, 6, 5, 4, 3, 2, 1}, want: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := InAscOrder(tc.input)

			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected: %v, got:%v", tc.want, got)
			}
		})
	}
}
