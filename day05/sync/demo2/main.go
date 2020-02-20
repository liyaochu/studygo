package main

import (
	"fmt"
	"sync"
	"time"
)

//读比写多时,需要使用读写锁

var x int64
var wg sync.WaitGroup
var lock sync.Mutex
var lock2 sync.RWMutex  //多个goroutine同时读加的是读锁,写时加写锁
func read() {
	defer wg.Done()
	lock2.RLock()
	fmt.Println(x)

	time.Sleep(time.Millisecond*10)
	lock2.RUnlock()
}

func write() {
	defer wg.Done()
	lock2.Lock()
	x = x + 1
	time.Sleep(time.Millisecond*50)
	lock2.Unlock()

}

func main() {
	now := time.Now()

	for i:=0;i<10 ;i++  {
		wg.Add(1)
		go  write()
	}
	for i:=0;i<10000 ;i++  {
		wg.Add(1)
		go  read()
	}
	wg.Wait()
	end := time.Now()
	fmt.Printf("耗费了多少时间%v",end.Sub(now))
}
