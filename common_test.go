package gcf_test

import (
	"testing"

	"github.com/meian/gcf"
	"github.com/stretchr/testify/assert"
)

// testBeforeAndAfter tests Iterator before calling first MoveNext or after gone last MoveNext.
// On bofore calling first, Iterator.Current() should be zero value.
// On after gone last, Iterator.MoveNext() should be false, and Iterator.Current() should be zero value.
func testBeforeAndAfter[T any](t *testing.T, itb gcf.Iterable[T]) {
	t.Helper()
	it := itb.Iterator()
	t.Run("before first MoveNext", func(t *testing.T) {
		assert.Zero(t, it.Current())
	})
	t.Run("after gone last MoveNext", func(t *testing.T) {
		assert := assert.New(t)
		for it.MoveNext() {
		}
		assert.False(it.MoveNext())
		assert.Zero(it.Current())
	})
}

// testEmpties tests Iterable func by empty element case variations.
//
// - emptyIterable chaining
//   - test any func chaining result from emptyIterable is emptyIterable or not.
// - no panic from empty
//   - test that panic does not occurred when any func chaining from empty elements that are not emptyIterable.
func testEmpties(t *testing.T, f func(itb gcf.Iterable[int]) gcf.Iterable[int]) {
	t.Helper()
	t.Run("emptyIterable chaining", func(t *testing.T) {
		itb := gcf.FromSlice([]int{}) // returns emptyIterable
		itb = f(itb)
		assert.True(t, gcf.IsEmptyIterable(itb), "%v", gcf.ToSlice(itb))
	})
	t.Run("no panic from empty", func(t *testing.T) {
		defer func() {
			err := recover()
			if err != nil {
				t.Errorf("%v", err)
			}
		}()
		// make empty elements but not emptyIterable.
		itb := gcf.FromSlice([]int{1})
		itb = gcf.Filter(itb, func(v int) bool { return false })
		itb = f(itb)
		_ = gcf.ToSlice(itb)
	})
}
