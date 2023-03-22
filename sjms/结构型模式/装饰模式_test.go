package 结构型模式

import (
	"fmt"
	"testing"
)

// 数据抽象
type Db interface {
	Query(string) string
}

// 装饰器基类
type Decorator struct {
	db Db
}

func (d *Decorator) Query(s string) string {
	return s
}

type SqlLite struct {
}

func (s *SqlLite) Query(s2 string) string {
	return s2
}

type DecoratorOne struct {
	// 继承基类（组合）
	Decorator
}

func (receiver *DecoratorOne) Query(s string) string {
	// 进行一些装饰
	fmt.Println("这是一些装饰")
	return receiver.Decorator.Query(s)
}

// 构造方法
func NewDecoratorOne(db Db) *DecoratorOne {
	return &DecoratorOne{Decorator{db}}
}

func TestDecorator(t *testing.T) {
	d := NewDecoratorOne(&SqlLite{})
	t.Log(d.Query("select id from user"))
}

// 可以参考DecoratorOne 实现其他装饰类型

// 代理模式和装饰模式虽然都属于结构型设计模式，但是它们的目的和实现方式有所不同。
//
// 代理模式旨在为一个对象提供一个代理，以控制对该对象的访问。代理对象通常充当另一个对象的接口，客户端代码只能通过代理访问实际对象。代理可以在不影响原始对象的情况下，执行一些附加的任务，比如：缓存数据、记录日志、验证权限等。
//
// 装饰模式旨在为一个对象动态地添加额外的行为，而不需要修改对象本身。装饰器对象通常与另一个对象组合，以增强该对象的行为。装饰器对象可以在不影响原始对象的情况下，增加新的行为或者修改现有行为，比如：加密数据、压缩数据、验证数据等。
