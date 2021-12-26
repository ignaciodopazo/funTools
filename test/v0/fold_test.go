package test

import (
	"strings"
	"testing"

	f "github.com/ignaciodopazo/funTools/src/v0"
	"github.com/ignaciodopazo/funTools/test/utils"
)

func TestFold(t *testing.T) {

	// for folding integers
	sum := func(x, y int) int {
		return x + y
	}
	t.Run("fold integers with +", func(t *testing.T) {
		got := f.Fold(sum, 0, []int{1, 2, 3, 4, 5})
		want := 15

		testutils.ExpectedResult(t, got, want)
	})
	// initial value
	t.Run("fold empty slice with +", func(t *testing.T) {
		got := f.Fold(sum, 0, []int{})
		want := 0

		testutils.ExpectedResult(t, got, want)
	})

	prod := func(x, y int) int {
		return x * y
	}
	t.Run("fold integers with *", func(t *testing.T) {
		got := f.Fold(prod, 1, []int{2, 3, 4})
		want := 24

		testutils.ExpectedResult(t, got, want)
	})
	t.Run("fold empty slice with *", func(t *testing.T) {
		got := f.Fold(prod, 1, []int{})
		want := 1

		testutils.ExpectedResult(t, got, want)
	})

	// folding strings
	appendStr := func(x, y string) string {
		return strings.Join([]string{x, y}, "")
	}
	t.Run("fold strings", func(t *testing.T) {
		got := f.Fold(appendStr, "", []string{"1", "2", "3", "4", "5"})
		want := "12345"

		testutils.ExpectedResult(t, got, want)
	})
	t.Run("fold empty string slice", func(t *testing.T) {
		got := f.Fold(appendStr, "AA", []string{})
		want := "AA"

		testutils.ExpectedResult(t, got, want)
	})

}
