package demo

import (
	"reflect"
	"sync"
	"testing"
)

func TestBytePool_Get(t *testing.T) {
	type fields struct {
		pool chan []byte
		l    int
		c    int
	}
	tests := []struct {
		name   string
		fields fields
		wantB  []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bp := &BytePool{
				pool: tt.fields.pool,
				l:    tt.fields.l,
				c:    tt.fields.c,
			}
			if gotB := bp.Get(); !reflect.DeepEqual(gotB, tt.wantB) {
				t.Errorf("Get() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func BenchmarkBytePool(b *testing.B) {
	bp := NewBytePool(500, 1024, 1024)
	var wg sync.WaitGroup
	wg.Add(500)
	for i := 0; i < 500; i++ {
		go func() {
			defer wg.Done()
			bt := bp.Get()
			defer bp.Put(bt)
			readFile(bt)
		}()
	}
	wg.Wait()
	// BenchmarkBytePool-8   	   10000	  14568543 ns/op
}

func BenchmarkPool(b *testing.B) {
	bp := &sync.Pool{
		New: func() interface{} {
			return make([]byte, 1024, 1024)
		},
	}
	var wg sync.WaitGroup
	wg.Add(500)
	for i := 0; i < 500; i++ {
		go func() {
			defer wg.Done()
			bt := bp.Get().([]byte)
			defer bp.Put(bt)
			readFile(bt)
		}()
	}
	wg.Wait()
}
