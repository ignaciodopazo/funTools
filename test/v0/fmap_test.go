package test

import (
	"strings"
	"testing"

	f "github.com/ignaciodopazo/funTools/src/v0"
	"github.com/ignaciodopazo/funTools/test/utils"
)

func TestFmap(t *testing.T) {

	t.Run("Fmap over an empty slice", func(t *testing.T) {
		got := f.Fmap(strings.ToLower, []string{})
		want := []string{}

		testutils.ExpectedResultSlices(t, got, want)
	})

	add1 := func(x int) int {
		return x + 1
	}
	t.Run("Add 1 to every element", func(t *testing.T) {
		got := f.Fmap(add1, []int{0, 1, 2, 3})
		want := []int{1, 2, 3, 4}

		testutils.ExpectedResultSlices(t, got, want)
	})

	t.Run(".toUpper every element", func(t *testing.T) {
		got := f.Fmap(strings.ToUpper, []string{"a", "b", "c"})
		want := []string{"A", "B", "C"}

		testutils.ExpectedResultSlices(t, got, want)
	})

}
