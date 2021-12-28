package test

import (
	"testing"

	f "github.com/ignaciodopazo/funTools/src/v0"
	"github.com/ignaciodopazo/funTools/test/utils"
)

func TestPartition(t *testing.T) {

	sl := []int{1,2,3,4,5,6,7,8,9}

	t.Run("no element satisfies the predicate", func (t *testing.T) {
		yes, no := f.Partition(func (x int) bool {
			return x > 10
		}, sl)
		expectedYes := []int{}
		expectedNo := []int{1,2,3,4,5,6,7,8,9}
		testutils.ExpectedResultSlices(t, yes, expectedYes)
		testutils.ExpectedResultSlices(t, no, expectedNo)
	})

	t.Run("every element satisfies the predicate", func (t *testing.T) {
		yes, no := f.Partition(func (x int) bool {
			return x < 10
		}, sl)
		expectedYes := []int{1,2,3,4,5,6,7,8,9}
		expectedNo := []int{}
		testutils.ExpectedResultSlices(t, yes, expectedYes)
		testutils.ExpectedResultSlices(t, no, expectedNo)
	})

	t.Run("some elements satisfies the predicate and others not", func (t *testing.T) {
		yes, no := f.Partition(func (x int) bool {
			return x < 5
		}, sl)
		expectedYes := []int{1,2,3,4}
		expectedNo := []int{5,6,7,8,9}
		testutils.ExpectedResultSlices(t, yes, expectedYes)
		testutils.ExpectedResultSlices(t, no, expectedNo)
	})

}
