package main

import (
	"sync"
	"fmt"
	"runtime"
)
/*

sync.Pool(临时对象池): 定位不是做类似连接池的东西, 它的用途仅仅是增加对象重用的几率, 减少gc负担, 开销也是很大的
可以把sync.Pool类型值看作是存放可被重复使用的值(临时对象)的容器, 这个容器是自动伸缩的, 高效的, 同时也是并发安全的


1. sync.Pool的缓存对象数量是没有限制的(只受限于内存), 因此使用sync.pool是没办法做到控制缓存对象数量个数

2. Pool中对象存活时间
	func init() {
		runtime_registerPoolCleanup(poolCleanup)
	}
	根据init()可知pool包在init的时候注册了一个poolCleanup函数,poolCleanup会清除所有pool里面的所有缓存的对象,
	该函数会在每次gc之前调用, 因此sync.Pool缓存的期限只是两次gc之间这段时间
	因为gc的时候会清掉缓存对象, 所以不用担心pool会无限增大的问题, 正因为这样,不可以使用sync.Pool去实现一个socket连接池或者
	类似的对象

3. 如何在多个goroutine之间使用同一个pool做到高效呢?
	官方的做法就是尽量减少竞争,因为sync.pool为每个CPU(对应cpu)都分配了一个子池,当执行一个pool的get或者put操作的时候都会先把
	当前的goroutine固定到某个CPU的子池上面, 然后对该子池进行操作. 每个子池里面有一个私有对象和共享列表对象, 私有对象只有对应
	的CPU能够访问, "因为一个CPU同一时间只能执行一个goroutine", 因此对私有对象存取操作是不需要加锁的.
	共享列表是和其他CPU分享的, 因此操作共享列表是需要加锁的

	获取对象过程是:
	1) 固定到某个CPU, 尝试从私有对象获取, 如果私有对象非空则返回该对象, 并把私有对象置空;
	2) 如果私有对象是空的时候, 就去当前子池的共享列表获取(需要加锁);
	3) 如果当前子池的共享列表也是空的, 那么就尝试去其他CPU的子池的共享列表偷取一个(需要加锁);
	4) 如果其他子池都是空的, 用指定的New函数产生一个新的对象返回
	一次get操作最少0次加锁, 最大N（N等于CPU数量)次加锁

	归还对象的过程：
	1) 固定到某个P, 如果私有对象为空则放到私有对象;
	2）否则加入到该P子池的共享列表中(需要加锁);
	可以看到一次put操作最少0次加锁, 最多1次加锁
*/

type Person struct {
	name string
}

func syncPool()  {
	// 获取对象的时候如果在池里面找不到缓存对象时, 将会使用指定的New函数创建一个返回, 如果没有new函数则返回nil
	// New相当于工厂函数
	pool := sync.Pool{
		New:func() interface{}{
			return Person{}
		},
	}
	// 从池里获取对象, 如果没有会用New函数创建一个对象
	tmp1 := pool.Get()
	tmp2 := pool.Get()
	if person, ok := tmp1.(Person); ok {
		person.name = "1"
		pool.Put(person) // 将对象放回池中
	}
	if person, ok := tmp2.(Person); ok {
		person.name = "2"
		pool.Put(person)
	}
	// 放回池中的对象不会清空里面的值, 放进去时是什么, 取出来还是什么
	fmt.Println(pool.Get(), pool.Get(), pool.Get())  // 1 2 {}, 1和2是上面赋值的对象
}

func gcPool() {
	pool := sync.Pool{
		New:func() interface{}{
			return Person{}
		},
	}
	p := pool.Get().(Person)
	p.name = "1"
	pool.Put(p)
	p = pool.Get().(Person)
	fmt.Println(p) // 2
	runtime.GC()
	fmt.Println(pool.Get()) // pool清空, 值为空
}

func main()  {
	// syncPool()
	gcPool()
}