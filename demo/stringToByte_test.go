package demo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStringToByte(t *testing.T) {
	type args struct {
		str *string
	}
	s := "acd"
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "字符串转字节切片", args: args{str: &s}, want: []byte{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToByte(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkName(b *testing.B) {
	// 相对更高效
	s := "字符串转字节切片"
	for i := 0; i < b.N; i++ {
		fmt.Println(StringToByte(&s))
		//BenchmarkName-8   	  179287	      5612 ns/op
	}
}

func BenchmarkByteToString(b *testing.B) {
	s := "字符串转字节切片"
	for i := 0; i < b.N; i++ {
		fmt.Println([]byte(s))
		// BenchmarkByteToString-8   	  206289	      7103 ns/op
	}
}

func BenchmarkName1(b *testing.B) {
	// 相对更高效
	bc := []byte{229, 173, 151, 231, 172, 166, 228, 184, 178, 232, 189, 172, 229, 173, 151, 232, 138, 130, 229, 136, 135, 231, 137, 135}
	for i := 0; i < b.N; i++ {
		fmt.Println(ByteToString(bc))
		//BenchmarkName1-8   	  357932	      3022 ns/op
	}
}

func BenchmarkByteToString2(b *testing.B) {
	bc := []byte{229, 173, 151, 231, 172, 166, 228, 184, 178, 232, 189, 172, 229, 173, 151, 232, 138, 130, 229, 136, 135, 231, 137, 135}
	for i := 0; i < b.N; i++ {
		fmt.Println(string(bc))
		// BenchmarkByteToString2-8   	  331899	      4501 ns/op
	}
}
