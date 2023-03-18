package demo

import (
	"errors"
	"fmt"
	"testing"
)

func Test_typeError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "类型断言 MyError", args: args{
			err: &MyError{
				errorMsg:  "自定义错误",
				err:       errors.New("原始错误"),
				errorCode: 400,
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			typeError(tt.args.err)
		})
	}
}

func Test_wrapErr(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "嵌套错误", args: args{err: errors.New(":原始错误")}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := wrapErr(tt.args.err); (err != nil) != tt.wantErr {
				t.Errorf("wrapErr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUnwrap(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "嵌套错误", args: args{err: &MyError{errorMsg: "嵌套错误", err: errors.New("原始错误")}}, wantErr: true},
		{name: "嵌套错误", args: args{err: fmt.Errorf("嵌套错误 %w", errors.New("原始错误"))}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Unwrap(tt.args.err); (err != nil) != tt.wantErr {
				t.Errorf("Unwrap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_errAs(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "类型判断",
			args: args{
				err: &MyError{
					err: errors.New("原始错误"),
				},
			},
			want: true,
		},
		{
			name: "类型判断",
			args: args{
				err: errors.New("原始错误"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := errAs(tt.args.err); got != tt.want {
				t.Errorf("errAs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errAsTwo(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "类型判断",
			args: args{
				err: &MyError{
					err: errors.New("原始错误"),
				},
			},
			want: true,
		},
		{
			name: "类型判断",
			args: args{
				err: errors.New("原始错误"),
			},
			want: false,
		},
		{
			name: "类型嵌套",
			args: args{
				err: fmt.Errorf("错误嵌套 %w", &MyError{
					err: errors.New("原始错误"),
				}),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := errAsTwo(tt.args.err); got != tt.want {
				t.Errorf("errAsTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
