package main

import (
	"bytes"
	"testing"
)

func BenchmarkString(b *testing.B) {
	var str string
	for i := 0; i < b.N; i++ {
		str = str + "s"
	}
}

func BenchmarkBufferString(b *testing.B) {
	bb := bytes.NewBufferString("")
	for i := 0; i < b.N; i++ {
		bb.WriteString("s")
	}
}
