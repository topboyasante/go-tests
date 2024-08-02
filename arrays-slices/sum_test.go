package main

import (
	"reflect"
	"testing"
)

func TestIntegers(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, expected %v", &got, want)
		}
	}

	t.Run("an array of any length of numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d, expected %d", &got, want)
		}
	})

	t.Run("gives the sum of a number of slices", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{1, 2, 3, 4, 5})
		want := []int{6, 15}

		checkSums(t, got, want)

	})

	t.Run("gives the sum of all the tails in a slice", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("gives the sum of all the tails in a slice(with some slices empty)", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, expected %v", &got, want)
		}
	})
}

func BenchmarkSum(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}
