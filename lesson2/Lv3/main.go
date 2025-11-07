package main

import (
	"fmt"
	"sync"
)

var (
	wg   sync.WaitGroup //定义wg为WaitGroup
	lock sync.Mutex     //定义并发锁
)

type Counter struct { //定义Counter结构体
	count int
}

func (a *Counter) Increment() { //实现Increment方法
	lock.Lock() //上锁防止多访问
	a.count += 1
	defer lock.Unlock() //解锁
}

func (a *Counter) Value() int { //实现Value方法
	lock.Lock()         //上锁
	defer lock.Unlock() //延迟解锁先返回值
	return a.count
}

func main() {
	wg.Add(100) //实现100个goroutine
	var counter Counter
	for i := 0; i < 100; i++ { //遍历
		go func() { //启动goroutine
			defer wg.Done()           //结束后减一
			for j := 0; j < 10; j++ { //循环10次
				counter.Increment()
			}
		}()
	}
	wg.Wait() //等待做完
	fmt.Println("最终计数:", counter.Value())
}
