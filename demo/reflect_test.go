package demo

import "testing"

func Test_ref(t *testing.T) {
	type args struct {
		input any
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "整型", args: args{input: 1}},
		{name: "字符串", args: args{input: "abc"}},
		{name: "字符", args: args{input: 'a'}},
		{name: "结构体", args: args{input: args{input: 1}}},
		{name: "bool", args: args{input: true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ref(tt.args.input)
		})
	}
}
