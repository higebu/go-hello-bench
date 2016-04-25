package hello

import (
	"log"
	"runtime"
	"testing"
)

func init() {
	log.Printf("NumCPU: %d\n", runtime.NumCPU())
}

var c1 int
var c2 int

func BenchmarkHello(b *testing.B) {
	log.Printf("c1: %d\n", c1)
	c1++
	for i := 0; i < b.N; i++ {
		c2++
		Hello("test")
	}
	log.Printf("c2: %d\n", c2)
}
