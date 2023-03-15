/**
 * @Author: realpeanut
 * @Date: 2023/3/15 11:52
 */

package demo

// People 接口型函数的应用场景 前提是这个接口只有一个方法，如果有两个以及以上方法，将不能用这种模式
type People interface {
	Do() error
}

// 定义一个函数，他的参数是接口

func DoF(p People) error {
	return p.Do()
}

// 想要调用函数DoF，则必须传入一个实现了People的参数，有以下几种方式
// 1、直接传入一个结构体，这个结构体实现了Do方法

type D struct {
}

func (receiver D) Do() error {
	return nil
}

// DoF(D{})  这样做的缺点是，只能先定义结构体或其他类型，然后再实现接口的方法

// 2、定义函数类型，利用强制类型转换间接实现接口的方法，这样就不用主动去实现接口方法了

// HandFunc 定义函数类型
type HandFunc func() error

// Do 函数类型实行接口方法
func (h HandFunc) Do() error {
	// 执行函数类型本身
	return h()
}
func main() {
	DoF(HandFunc(func() error {
		// 任意实现
		return nil
	}))
}
