package arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {

	t.Run("Should add at same index", func(t *testing.T) {

		got := SumAll([]int{1, 2, 4}, []int{0, 9, 6})
		want := []int{7, 15}

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	})
}
