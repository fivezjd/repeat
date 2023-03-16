/**
 * @Author: realpeanut
 * @Date: 2023/3/16 23:10
 */
package demo

import (
	"errors"
	"fmt"
	"sync"
	"testing"
)

func TestMap(t *testing.T) {
	var m sync.Map
	for i := 0; i < 1000; i++ {
		index := i
		go func() {
			m.Store("a", index)
			v, _ := m.Load("a")
			fmt.Println(v)
		}()
	}
}

func BenchmarkMap(b *testing.B) {
	var m sync.Map
	for i := 0; i < b.N; i++ {
		index := i
		go func() {
			m.Store("a", index)
			v, _ := m.Load("a")
			fmt.Println(v)
		}()
	}
}

func BenchmarkMapLock(b *testing.B) {
	var mu sync.RWMutex
	mapData := make(map[string]int)
	for i := 0; i < b.N; i++ {
		index := i
		go func() {
			mu.Lock()
			mapData["string"] = index
			fmt.Println(mapData["string"])
			mu.Unlock()
		}()
	}
}

type MapI interface {
	Get(key string) (string, error)
}

type WithOutMap struct {
	store map[string]string
}

func (w *WithOutMap) Get(key string) (string, error) {
	if v, ok := w.store[key]; ok {
		return v, nil
	}
	return "", errors.New("元素不存在")
}

type DecoratorMap struct {
	MapI
	mu sync.Mutex
}

func (d *DecoratorMap) Get(key string) (string, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.MapI.Get(key)
}

// 装饰器模式，在不修改WithOutMap 的前提下，为WithOutMap新增了线程安全
func TestName(t *testing.T) {
	ml := &DecoratorMap{
		MapI: &WithOutMap{
			store: make(map[string]string),
		},
	}
	ml.Get("a")
}
