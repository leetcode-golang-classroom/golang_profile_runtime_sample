package main

import (
	"fmt"
	"runtime"
)

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MB\n", bToMb(m.Alloc))
	fmt.Printf("TotalAlloc = %v MB\n", bToMb(m.TotalAlloc))
	fmt.Printf("Sys = %v MB\n", bToMb(m.Sys))
	fmt.Printf("NumGC = %v \n", m.NumGC)
}
func bToMb(b uint64) uint64 {
	return b / 1000 / 1000
}
func main() {
	fmt.Println("Mem stats before:")
	printMemStats()
	s := make([]int, 10_000_000)
	for i := range s {
		s[i] = i
	}
	fmt.Println("Mem stats after:")
	printMemStats()
}
