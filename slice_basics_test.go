package main

import (
	"reflect"
	"testing"
)

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func BenchmarkDoubleSliceElements(b *testing.B) {
	tests := []struct {
		name     string
		inputt   []int
		expected []int
	}{
		{name: "BenchmarkDoubleSliceElements small", inputt: []int{1, 2, 3, 4, 5}, expected: []int{2, 4, 6, 8, 10}},
		{name: "BenchmarkDoubleSliceElements negative", inputt: []int{10, 20, 30, -40, -50}, expected: []int{20, 40, 60, -80, -100}},
		{name: "BenchmarkDoubleSliceElements empty", inputt: []int{}, expected: []int{}},
		{name: "BenchmarkDoubleSliceElements nil", inputt: nil, expected: []int{}},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				doubleSliceElements(tt.inputt)
			}
		})
	}
}

func TestDoubleSliceElements(t *testing.T) {
	tests := []struct {
		name     string
		inputt   []int
		expected []int
	}{
		{name: "TestDoubleSliceElements small", inputt: []int{1, 2, 3, 4, 5}, expected: []int{2, 4, 6, 8, 10}},
		{name: "TestDoubleSliceElements negative", inputt: []int{10, 20, 30, -40, -50}, expected: []int{20, 40, 60, -80, -100}},
		{name: "TestDoubleSliceElements empty", inputt: []int{}, expected: []int{}},
		{name: "TestDoubleSliceElements nil", inputt: nil, expected: []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := doubleSliceElements(tt.inputt)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Fatalf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
