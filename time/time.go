package main

import (
	"fmt"
	"reflect"
	"time"
	"sync"
)

func study() {
	fmt.Println("1s = ", time.Second, int(time.Second))
	fmt.Println("1ms = ", time.Millisecond, int(time.Millisecond))
	// 1 ns = 1.0 × 10^(-9)s
	fmt.Println("1ns = ", time.Nanosecond, int(time.Nanosecond))
	// 返回当前本地时间
	fmt.Println("Now = ", time.Now())
	/*
		func (t Time) UnixNano() int64
		UnixNano将t表示为Unix时间, 即从时间点January 1, 1970 UTC到时间点t所经过的时间(单位纳秒)
		如果纳秒为单位的unix时间超出了int64能表示的范围, 结果是未定义的
	*/
	sunday := time.Sunday
	fmt.Println("sunday =", sunday, "type =", reflect.TypeOf(sunday), "tostring =", sunday.String())

	{
		start := time.Now()
		fmt.Println("year =", start.Year(), "month =", start.Month(),
			"day =", start.Day(), "hour =", start.Hour(),
			"minute =", start.Minute(), "s =", start.Second())
		time.Sleep(1 * time.Second) // 如果直接写int, 1代表1ns
		end := time.Now()
		/*
			func (t Time) Sub(u Time) Duration, 返回一个时间段t-u
			func (t Time) Before(u Time) bool, 如果t代表的时间点在u之前, 返回真; 否则返回假
			func (t Time) After(u Time) bool, 如果t代表的时间点在u之后, 返回真; 否则返回假
		*/
		fmt.Println("sub = end - start = ", end.Sub(start))
		fmt.Println("before = ", start.Before(end))
		fmt.Println("after = ", end.After(start))
	}
}

func timeAfter(waitTime int) {
	ch := make(chan int, 10)
	wg := sync.WaitGroup{}
	wg.Add(2)
	runFlag := true
	send := func() {
		for i := 0; i < 10 && runFlag; i++ {
			time.Sleep(time.Duration(i) * time.Second)
			ch <- i
		}
		wg.Done()
	}
	recv := func() {
		for runFlag {
			select {
			case i := <-ch:
				fmt.Println("recv", i)
			case <-time.After(time.Duration(waitTime) * time.Second):
				wg.Done()
				fmt.Println("break")
				runFlag = false
				break // 这个break对应的是select, 可以直接写return
			}
		}
	}
	go send()
	go recv()
	wg.Wait()
	fmt.Println("finish")
}

func getLocation() *time.Location {
	local, _ := time.LoadLocation("America/New_York")
	return local
}

func now()  {
	fmt.Println(time.Now().In(getLocation()).Format("20060102"))
	fmt.Println(fmt.Sprintf(time.Now().In(getLocation()).Format("20060102150405")))
}
func main() {
	//timeAfter(3)
	now()
}
