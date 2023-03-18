package demo

import (
	"fmt"
	"sync"
	"testing"
)

type DataSore struct {
	mu   sync.RWMutex
	data int
}

func TestRw(t *testing.T) {
	d := DataSore{}
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()

			d.mu.Lock()
			d.data += 1
			d.mu.Unlock()
		}()
	}
	fmt.Println(d.data)
	wg.Wait()
	fmt.Println(d.data)
}

func TestGetStingByInt(t *testing.T) {
	var tests = []struct {
		name  string
		input uint
		want  string
	}{
		{"a", 65, "A"},
		{"b", 66, "B"},
		{"c", 67, "C"},
		{"d", 68, "D"},
		{"e", 69, "E"},
		{"f", 70, "F"},
		{"g", 71, "G"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ans := getStingByInt(test.input)
			if ans != test.want {
				t.Errorf("got %s, want %s", ans, test.want)
			}
		})
	}
}

func Test_getStingByInt(t *testing.T) {
	type args struct {
		n uint
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "65 should be A", args: args{n: 65}, want: "A"},
		{name: "66 should be B", args: args{n: 66}, want: "B"},
		{name: "67 should be C", args: args{n: 67}, want: "C"},
		{name: "68 should be D", args: args{n: 68}, want: "D"},
		{name: "69 should be E", args: args{n: 69}, want: "E"},
		{name: "70 should be F", args: args{n: 70}, want: "F"},
		{name: "71 should be G", args: args{n: 71}, want: "G"},
		{name: "256 should be Ā", args: args{n: 256}, want: "Ā"},
		{name: "256 should be Ā", args: args{n: 10000}, want: "Ā"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStingByInt(tt.args.n); got != tt.want {
				t.Errorf("getStingByInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getIntByChar(t *testing.T) {
	type args struct {
		c rune
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{name: "A should be 65", args: args{c: 'A'}, want: 65},
		{name: "张 should be 24352", args: args{c: '张'}, want: 24352},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntByChar(tt.args.c); got != tt.want {
				t.Errorf("getIntByChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCase(t *testing.T) {
	caseSwitch()
}

func Test_mutipara(t *testing.T) {
	type args struct {
		option []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "输入1 输出1", args: args{option: []int{1, 2, 3, 4, 5}}},
		{name: "输入1 输出1", args: args{option: []int{1, 2, 3, 4, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mutipara(tt.args.option...)
		})
	}
}

func Test_showCond(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "cond"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			showCond()
		})
	}
}

func TestCondT(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "cond"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CondT()
		})
	}
}
