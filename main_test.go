package main

import (
	"bytes"
	"testing"
)

func TestHello(t *testing.T) {
	t.Log("Hello")
}

func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		for _, elem := range elems {
			buf.WriteString(elem)
		}
	}

	b.StopTimer()
}
