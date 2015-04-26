package flag

import (
	"runtime"
	"sync/atomic"
)

type Flag struct {
	Me    *uint32
	Other *uint32
}

func New() (a *Flag, b *Flag) {
	a = &Flag{Me: new(uint32)}
	b = &Flag{Me: new(uint32)}
	a.Other = b.Me
	b.Other = a.Me

	return a, b
}

func (f *Flag) Lock() {
	atomic.StoreUint32(f.Me, 1)
	for atomic.LoadUint32(f.Other) == 1 {
		atomic.StoreUint32(f.Me, 0)
		runtime.Gosched()
		atomic.StoreUint32(f.Me, 1)
	}
}

func (f *Flag) Unlock() {
	atomic.StoreUint32(f.Me, 0)
}
