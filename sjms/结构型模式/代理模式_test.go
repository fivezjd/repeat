package 结构型模式

import (
	"log"
	"strings"
	"testing"
)

// Query 抽象接口
type Query interface {
	Exec(sql string) interface{}
}

// ProxyQuery 代理层
type ProxyQuery struct {
	query Query
}

func (p *ProxyQuery) Exec(sql string) interface{} {
	// 在代理层做一些其他的事情
	log.Println("代理层执行query之前做一些事情")
	// 执行入参的方法
	return p.query.Exec(sql)
}

// MySql 具体执行层
type MySql struct {
}

func (m *MySql) Exec(sql string) interface{} {
	s := strings.Builder{}
	s.WriteString(sql)
	s.WriteString(":")
	s.WriteString("mysql")
	return s.String()
}

// Mongo 具体执行层
type Mongo struct {
}

func (m *Mongo) Exec(sql string) interface{} {
	s := strings.Builder{}
	s.WriteString(sql)
	s.WriteString(":")
	s.WriteString("mongo")
	return s.String()
}

// 如果想在不修改执行层的前提下，拓展执行层的能力，可以使用代理模式
// 如此这般，符合开闭原则，（修改闭合，新增开发）
// 这就是代理模式，此外，上面代码的实现中，还应用了接口嵌入结构体的特性。
func TestProxy(t *testing.T) {
	db := &ProxyQuery{
		query: new(MySql),
	}
	res := db.Exec("select id,name,email from `user`")
	t.Log(res)
}

// 优点：在不修改子系统代码的前提下，灵活新增子系统功能。符合开闭原则，一定程度上解耦
