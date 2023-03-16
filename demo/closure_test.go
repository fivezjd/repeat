/**
 * @Author: realpeanut
 * @Date: 2023/3/15 12:21
 */

// Package demo 闭包
package demo

import (
	"context"
	"fmt"
	"testing"
)

// 函数闭包：匿名函数 + 定义它的上下文

func TestClosure(t *testing.T) {
	t.Log(closure("abc")())
}

func closure(str string) func() string {
	// 匿名函数能访问定义它的上游变量
	return func() string {
		return str
	}
}

// 闭包延时绑定
func Delay() {
	fns := make([]func(), 0, 10)

	for i := 0; i < 10; i++ {
		fns = append(fns, func() {
			fmt.Println(i)
		})
	}

	for _, fn := range fns {
		fn()
	}
}

func TestDelay(t *testing.T) {
	Delay()
}

// 责任链模式的实现

// 方法1
type Filter func(c context.Context)

type FilterBuilder func(next Filter) Filter

type Server struct {
	name string
	hand Filter
}

func makeBuilder(name string) FilterBuilder {
	return func(next Filter) Filter {
		fmt.Println(name)
		return func(c context.Context) {
			next(c)
		}
	}
}

func NewServer(name string, builders ...FilterBuilder) Server {
	// 最开始的一个Filter
	first := func(ctx context.Context) {
		fmt.Println("first")
	}
	for _, builder := range builders {
		first = builder(first)
	}

	return Server{
		name: name,
		hand: first,
	}
}

func twoFilterBuilder(next Filter) Filter {
	return func(c context.Context) {
		fmt.Println("two")
		next(c)
	}
}

func threeFilterBuilder(next Filter) Filter {
	return func(c context.Context) {
		fmt.Println("three")
		fmt.Println(c.Value("name"))
		next(c)
	}
}

func TestFilter(t *testing.T) {
	s := NewServer("abc", twoFilterBuilder, threeFilterBuilder, makeBuilder("oki"))
	fmt.Println(s.name)
	ctx := context.WithValue(context.Background(), "name", s.name)
	s.hand(ctx)
}
