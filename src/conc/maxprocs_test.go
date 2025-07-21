package conc

import (
	"log"
	"runtime"
	"testing"
)

func TestMaxProcs(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Printf(">> number of CPUs: %v", runtime.NumCPU())
}
