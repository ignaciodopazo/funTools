package test

import (
	"testing"

	f "github.com/ignaciodopazo/funTools/src/v0"
	"github.com/ignaciodopazo/funTools/test/utils"
)

func TestFilter(t *testing.T) {

	isGreaterThan10 := func(x int) bool {
		return x > 10
	}

	t.Run("Filter over an empty slice", func(t *testing.T) {
		got := f.Filter(isGreaterThan10, []int{})
		want := []int{}

		testutils.ExpectedResultSlices(t, got, want)
	})

	t.Run("Filter returns an empty slice", func(t *testing.T) {
		got := f.Filter(isGreaterThan10, []int{1, 2, 3, 4, 9, -1})

		testutils.ExpectedResultSlices(t, got, nil)
		testutils.CheckPredicateOverSlice(t, isGreaterThan10, got)
	})

	t.Run("Filter returns a non-empty slice", func(t *testing.T) {
		got := f.Filter(isGreaterThan10, []int{1, 2, 11, 20, 500, -2})

		testutils.ExpectedResultSlices(t, got, []int{11, 20, 500})
		testutils.CheckPredicateOverSlice(t, isGreaterThan10, got)
	})

}
