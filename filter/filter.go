package filter

import (
	"reflect"
	"runtime"
	"sync/atomic"
	"unsafe"
)

type Filter struct {
	levels  []int64
	victims []int64
}

func New(n int64) *Filter {
	levels := make([]int64, n)
	for i := range levels {
		levels[i] = -1
	}
	victims := make([]int64, n)
	for i := range victims {
		victims[i] = -1
	}
	return &Filter{
		levels:  levels,
		victims: victims,
	}
}

func StoreSliceAtomic(s []int64, idx int64, val int64) {
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&s))
	// RIP memory safety
	addr := header.Data + uintptr(idx)*unsafe.Sizeof(idx)
	atomic.StoreInt64((*int64)(unsafe.Pointer(addr)), val)
}

func loadSliceAtomic(s []int64, idx int64) int64 {
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&s))
	addr := header.Data + uintptr(idx)*unsafe.Sizeof(idx)
	return atomic.LoadInt64((*int64)(unsafe.Pointer(addr)))
}

func (f *Filter) Lock(id int64) {
	levels := int64(len(f.levels))
	for level := int64(1); level < levels; level++ {
		StoreSliceAtomic(f.levels, id, level)
		StoreSliceAtomic(f.victims, level, id)
		for k := range f.levels {
			for int64(k) != id &&
				loadSliceAtomic(f.levels, int64(k)) >= level &&
				loadSliceAtomic(f.victims, int64(level)) == id {
				runtime.Gosched()
			}
		}
	}
}

func (f *Filter) Unlock(id int64) {
	f.levels[id] = -1
}
