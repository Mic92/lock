package main

import (
	"fmt"
	"github.com/Mic92/lock/filter"
	"math"
	"runtime"
	"sync"
	"sync/atomic"
)

const (
	THREADS = 32
	ROUNDS  = 1e7
	SAMPLES = 100
)

func main() {
	result := make(chan int)
	mutex := sync.Mutex{}
	for procs := 1; procs <= runtime.NumCPU(); procs++ {
		runtime.GOMAXPROCS(procs)
		for threads := int64(1); threads <= THREADS; threads++ {
			var positionCounter int64 = 0
			f := filter.New(threads)
			threadPositions := make([]int64, threads)
			for id := int64(0); id < threads; id++ {
				go func(id int64) {
					overtaken := 0
					for i := 0; i < ROUNDS; i++ {
						mutex.Lock()
						threadPositions[id] = atomic.AddInt64(&positionCounter, 1)
						mutex.Unlock()

						f.Lock(id)
						ownPos := threadPositions[id]
						threadPositions[id] = math.MaxInt64
						for otherId, pos := range threadPositions {
							if id != int64(otherId) && ownPos > pos {
								overtaken += 1
							}
						}
						f.Unlock(id)
					}
					result <- overtaken
				}(id)
			}
			sum := 0
			for j := int64(0); j < threads; j++ {
				sum += <-result
			}
			// OS-Threads, Goroutines, Overtakens
			fmt.Printf("%d,%d,%d,%d,%f\n", procs, threads, sum, int64(sum)/threads,
				float64(sum)/float64(ROUNDS)/float64(threads))
		}
	}
}
