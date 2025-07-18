package main

import (
	"fmt"
	"runtime"
	"time"
)

// Helper function to print memory stats in a readable format (KB, MB, GB)
func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %s", formatBytes(m.Alloc))
	fmt.Printf("\tTotalAlloc = %s", formatBytes(m.TotalAlloc))
	fmt.Printf("\tSys = %s", formatBytes(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

// formatBytes converts a uint64 of bytes to a formatted string.
func formatBytes(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}

func MemUsageExe() {
	fmt.Println("Initial memory usage:")
	printMemUsage()

	// Allocate a large slice to demonstrate memory usage change
	var bigSlice []int
	for i := 0; i < 20_000_000; i++ {
		bigSlice = append(bigSlice, i)
	}

	fmt.Println("\nMemory usage after allocating a large slice:")
	printMemUsage()

	// Force a garbage collection to see its effect
	runtime.GC()
	fmt.Println("\nMemory usage after garbage collection:")
	printMemUsage()

	// Keep the program running for a moment to observe
	time.Sleep(2 * time.Second)
	bigSlice = nil // Release the slice
	runtime.GC()
	fmt.Println("\nMemory usage after releasing slice and running GC:")
	printMemUsage()
}
