package demo

import (
	"errors"
	"fmt"
)

// MyError 自定义error 嵌套 嵌套到自定义结构体中的错误不能被unwrap
type MyError struct {
	errorCode int
	errorMsg  string
	err       error
}

func (m *MyError) Error() string {
	return m.err.Error() + "----" + m.errorMsg
}

// error 类型断言

func typeError(err error) {
	if err, ok := err.(*MyError); ok {
		fmt.Println(err)
	}
}

// 错误嵌套

func wrapErr(err error) error {
	return fmt.Errorf("嵌套一个原始错误：%w", err)
}

// 错误解引

func Unwrap(err error) error {
	fmt.Println(err)
	e := errors.Unwrap(err)
	fmt.Println(e)
	return e
}

// 因为嵌套后的error直接用 != 返回的是false ,所以为了判断B是否被A嵌套了，可以用下面的方法：

func errIs(A, B error) bool {
	// A B是同一个变量时，返回true，B被A嵌套了也会返回true
	return errors.Is(A, B)
}

// 嵌套后的类型断言是不能用的，所以可以用下面的方法

func errAs(err error) bool {
	var e *MyError
	// 判断err 是不是MyError类型的错误
	return errors.As(err, &e)
}

func errAsTwo(err error) bool {
	var e *MyError
	// 判断err 是不是MyError类型的错误
	return errors.As(err, &e)
}
