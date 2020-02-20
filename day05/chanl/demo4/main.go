package main

import (
	"fmt"
	"math/rand"
)

//生产者与消费者
//利用goroutine 和channel实现一个生产这消费者模型

//生产者;产生一个随机数
//消费者:计算每个随机数每个数字的和 123  6
type ch1 chan *item
 type resultChan chan *result
type ch2 chan *result

type item struct {
	id  int64
	num int64
}

type result struct {
	*item
	num int64
}

//生产者
func produce(ch chan *item) {
	//1.生成随机数
	var id int64
	for {
		id++
		num := rand.Int63()
		tmp := &item{
			id:  id,
			num: num,
		}
		//把随机数放到通道
		ch <- tmp
	}

	//2.把随机数发到通道中

}

//计算数字的和
func cacl(num int64) int64{
var  sum int64
for num > 0 {
		sum = sum + num%10
		num=num/10
	}
	return sum
}

//消费者
func consumer(ch chan *item,resultChan chan *result) {

	for tmp:= range ch{
		sum := cacl(tmp.num)
		retObj := &result{
			item: tmp,
			num:  sum,
		}
		resultChan<-retObj
	}

	}


func printResult(resultChan chan *result) {
 for rest:= range resultChan{
 	fmt.Printf("id:%v,  sum:%v  ,num:%v",rest.item.id,rest.item.num,rest.num)
 }
}

func startwork(n int,ch chan *item,resultChan chan *result){
	for i:=0; i<n;i++  {
      go consumer(ch,resultChan)
	}
}
func main() {
	itemChan := make(chan *item, 100)
	resChan := make(chan *result, 100)
	go produce(itemChan)
	//go consumer(itemChan,resChan)
	startwork(20,itemChan,resChan)
	printResult(resChan)
	////每一次执行都会是一个新的值
	//rand.Seed(time.Now().UnixNano())
	//
	//intn := rand.Intn(100) //返回0-100的随机数
	//fmt.Println(intn)

}
