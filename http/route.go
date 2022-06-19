/**
 * @Author: realpeanut
 * @Date: 2022/6/18 22:39
 */
package main

import "repeat/framework"

func registerRouter(core *framework.Core) {
	// 设置控制器
	core.Get("foo", FooControllerHandler)
}
