package demo

import (
	"fmt"
	"log"
	"math"
	"sync"
	"time"
)

func getStingByInt(n uint) string {
	if n <= math.MaxUint8 {
		return string(byte(n))
	}
	return string(rune(n))
}

func getIntByChar(c rune) uint {
	return uint(c)
}

func caseSwitch() {
	switch i := 1; i {
	case 1:
		fmt.Println(1)
		fallthrough
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	}
}

func mutipara(option ...int) {
	for _, op := range option {
		fmt.Println(op)
	}
}

// 方法作为表达式赋值给一个变量？可以 不光函数是一等公民，方法也是一等公民

type server struct {
}

func (receiver server) returnFunc() int {
	return 1
}

func tServer() func() int {
	f := server{}.returnFunc
	return f
}

// ---------------------
// cond 的使用   最好不用，容易出现死锁
func showCond() {
	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(10)
	//var ready bool

	for i := 0; i < 10; i++ {
		index := i
		go func() {
			defer wg.Done()
			fmt.Printf("协议%d 就绪\n", index)
			cond.L.Lock()
			defer cond.L.Unlock()

			// 等待所有的goroutine都已经准备就绪
			//for !ready {
			cond.Wait()
			//}

			// 收到信号
			fmt.Printf("协议%d 收到信号\n", index)
		}()
	}

	// 通知所有goroutine开始等待
	fmt.Println("通知协程开始执行")
	time.Sleep(time.Second * 2)
	cond.L.Lock()
	//ready = true
	cond.Broadcast()
	cond.L.Unlock()

	wg.Wait()
}

func CondT() {
	c := sync.NewCond(&sync.Mutex{})
	var ready int

	for i := 0; i < 10; i++ {
		go func(i int) {
			c.L.Lock()
			ready++
			c.L.Unlock()
			log.Printf("运动员%d 已准备好", i)
			c.Broadcast()
		}(i)
	}
	c.L.Lock()
	for ready != 10 {
		c.Wait()
		log.Println("裁判员被唤醒1次")
	}
	c.L.Unlock()
	log.Println("所以运动员准备就绪")
}
