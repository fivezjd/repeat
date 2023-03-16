/**
 * @Author: realpeanut
 * @Date: 2023/3/15 22:22
 */
package demo

import (
	"context"
	"testing"
	"time"
)

// 父控制子 控制是从上至下的 查找是从下至上的

func TestContext(t *testing.T) {
	// 10秒的超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	rangeTimer := time.NewTimer(time.Second)
	for {
		select {
		case <-ctx.Done():
			t.Log("10秒超时到期")
			return
		case <-rangeTimer.C:
			t.Log("1秒自动重置")
			rangeTimer.Reset(time.Second)
		}
	}
}
