//go:build v7

package orm

import "gitee.com/geektime-geekbang/geektime-go/orm/v7/internal/errs"

// 将内部的 sentinel error 暴露出去
var (
	// ErrNoRows 代表没有找到数据
	ErrNoRows = errs.ErrNoRows
)
