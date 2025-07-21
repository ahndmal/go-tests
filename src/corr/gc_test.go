package corr

import (
	"math"
	"testing"
	"time"
)

// Allocate an object of size bytes.
// Small objects are allocated from the per-P cache's free lists.
// Large objects (> 32 kB) are allocated straight from the heap.
// https://go.dev/src/runtime/malloc.go

func TestMain(m *testing.M) {
	ballast := make([]byte, 100<<20)
	for i := 0; i < len(ballast)/2; i++ {
		ballast[i] = byte('A')
	}
	<-time.After(time.Duration(math.MaxInt64))
}

//func mallocgc(size uintptr, typ *_type, needzero bool) unsafe.Pointer {
//
//	// Some error checking and debug code omitted here...
//
//	// assistG is the G to charge for this allocation, or nil if
//	// GC is not currently active.
//	var assistG *g
//	if gcBlackenEnabled != 0 {
//		// Charge the current user G for this allocation.
//		assistG = getg()
//		if assistG.m.curg != nil {
//			assistG = assistG.m.curg
//		}
//		// Charge the allocation against the G. We'll account
//		// for internal fragmentation at the end of mallocgc.
//		assistG.gcAssistBytes -= int64(size)
//
//		if assistG.gcAssistBytes < 0 {
//			// This G is in debt. Assist the GC to correct
//			// this before allocating. This must happen
//			// before disabling preemption.
//			gcAssistAlloc(assistG)
//		}
//	}
//
//	// Actual allocation code ommited below...
//}
