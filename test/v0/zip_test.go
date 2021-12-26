package test

import (
	"testing"

	f "github.com/ignaciodopazo/funTools/src/v0"
	"github.com/ignaciodopazo/funTools/test/utils"
)

func TestZip(t *testing.T) {

	t.Run("left slice is empty, so result must be empty", func (t *testing.T) {
		got := f.Zip([]float32 {}, []float64 {1.0,5.0})
		want := []f.Pair[float32, float64] {}
		testutils.ExpectedResultSlices(t, got, want)
	})

	t.Run("right slice is empty, so result must be empty", func (t *testing.T) {
		got := f.Zip([]int {1,2,6,0}, []string {})
		want := []f.Pair[int, string] {}
		testutils.ExpectedResultSlices(t, got, want)
	})

	t.Run("non-empty same-length slices Zip", func (t *testing.T) {
		got := f.Zip([]int {1,2,-1,-2}, []bool {true,true,false,false})
		want := []f.Pair[int, bool] { {1,true}, {2,true}, {-1,false}, {-2,false} }
		testutils.ExpectedResultSlices(t, got, want)
	})

	t.Run("non-empty different-length slices Zip", func (t *testing.T) {
		got := f.Zip([]int {1,2}, []bool {true,true,false,false})
		want := []f.Pair[int, bool] { {1,true}, {2,true} }
		testutils.ExpectedResultSlices(t, got, want)
	})

	t.Run("non-empty different-length slices Zip", func (t *testing.T) {
		got := f.Zip([]int {1,2,3,4}, []bool {true,true})
		want := []f.Pair[int, bool] { {1,true}, {2,true} }
		testutils.ExpectedResultSlices(t, got, want)
	})

}
