package flag

import (
	"sync/atomic"
	"testing"
)

func lock(t *testing.T, f *Flag, sanityCounter *int32) {
	for i := 0; i < 1e6; i++ {
		f.Lock()

		c := atomic.AddInt32(sanityCounter, 1)
		if c >= 2 {
			t.Fatalf("At least 2 goroutines have enter the critical section: %d, %d\n", c, i)
		}
		atomic.AddInt32(sanityCounter, -1)

		f.Unlock()
	}
}

func TestLockUnlock(t *testing.T) {
	var sanityCounter int32 = 0
	a, b := New()
	go lock(t, a, &sanityCounter)
	lock(t, b, &sanityCounter)
}
