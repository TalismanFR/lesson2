package main

import (
	"bytes"
	"encoding/json"
	"sync"
	"testing"
)

type Sig struct {
	Name string
	Size int
	Sig  []byte
}

var List = []Sig{
	{"name", 50, []byte{50, 60}},
	{"name1", 51, []byte{50, 60, 70}},
	{"name2", 510, []byte{50, 80, 70}},
}

func BenchmarkAllocMem(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			data := bytes.NewBuffer(make([]byte, 0, 64))
			_ = json.NewEncoder(data).Encode(List)
		}
	})
}

var dataPoll = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 64))
	},
}

func BenchmarkAllocPoolMem(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			data := dataPoll.Get().(*bytes.Buffer)
			_ = json.NewEncoder(data).Encode(List)
			data.Reset()
			dataPoll.Put(data)
		}
	})
}
