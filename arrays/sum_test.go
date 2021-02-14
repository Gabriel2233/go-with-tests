package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1}

		got := Sum(numbers)
		want := 1

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

// func TestSumAll(t *testing.T) {
// 	got := SumAllTails([]int{1, 2}, []int{3, 4})
// 	want := []int{3, 7}

// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("got %v want %v", got, want)
// 	}
// }

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 5}, []int{3, 4, 10})
		want := []int{7, 14}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 1})
		want := []int{0, 1}
		checkSums(t, got, want)
	})
}
