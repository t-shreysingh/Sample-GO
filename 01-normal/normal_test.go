// package normal

// import "testing"

// func TestSum(t *testing.T) {

// 	want := 8
// 	got := Sum(3, 5)

// 	if got != want {
// 		t.Errorf("Test fail! want: '%d', got: '%d'", want, got)
// 	}

// }

// func TestSum2(t *testing.T) {

// 	want := 10
// 	got := Sum(7, 2)

// 	if got != want {
// 		t.Errorf("Test fail! want: '%d', got: '%d'", want, got)
// 	}

// }

package main

import "testing"

func Test1(t *testing.T) {
	tests := []struct {
		number int
		want   int
	}{
		{number: 10, want: 100},
		{number: 5, want: 25},
		{number: 7, want: 49},
	}

	for i, tc := range tests {
		result := square(tc.number)
		if result != tc.want {
			t.Fatalf("Test %d failed — Expected %d, got %d", i+1, tc.want, result)
		}
	}
}

func Test11(t *testing.T) {
	tests := []struct {
		number int
		want   int
	}{
		{number: 9, want: 362880},
		{number: 5, want: 120},
		{number: 0, want: 1},
	}

	for i, tc := range tests {
		result := factorial(tc.number)
		if result != tc.want {
			t.Fatalf("Test %d failed — Expected %d, got %d", i+1, tc.want, result)
		}
	}
}
