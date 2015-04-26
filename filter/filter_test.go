package filter

import (
	"sync/atomic"
	"testing"
)

func lock(t *testing.T, f *Filter, id int64, sanityCounter *int32) {
	for i := 0; i < 1e6; i++ {
		f.Lock(id)

		c := atomic.AddInt32(sanityCounter, 1)
		if c >= 2 {
			t.Fatalf("At least 2 goroutines have enter the critical section: %d, %d\n", c, i)
		}
		atomic.AddInt32(sanityCounter, -1)

		f.Unlock(id)
	}
}

func TestLockUnlock(t *testing.T) {
	var sanityCounter int32 = 0
	f := New(2)
	go lock(t, f, 0, &sanityCounter)
	lock(t, f, 1, &sanityCounter)
}
