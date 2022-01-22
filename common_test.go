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
