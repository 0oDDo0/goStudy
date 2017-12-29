package main

import (
	"fmt"
	"unicode/utf8"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"reflect"
	"sync"
	"time"
)

/*
http://colobu.com/2015/09/07/gotchas-and-common-mistakes-in-go-golang/

1) 开大括号不能放在单独的一行
2) 未使用的包和变量
3) 方法的接收者就像常规的函数参数,如果声明为值, 那么函数/方法得到的是接收者参数的拷贝.
   这意味着对接收者所做的修改将不会影响原有的值, 除非接收者是一个map或者slice变量
4) 并不总是知道变量是分配到栈还是堆上. 在C++中, 使用new创建的变量总是在堆上. 在Go中, 即使是使用new()或者make()函数来分配,
   变量的位置还是由编译器决定. 编译器根据变量的大小和"泄露分析"的结果来决定其位置. 这也意味着在局部变量上返回引用是没问题的,
   而这在C或者C++这样的语言中是不行的
*/

// 不能在一个单独的声明中重复声明一个变量,但在多变量声明中这是允许的,其中至少要有一个新的声明变量.
// 重复变量需要在相同的代码块内, 否则你将得到一个隐藏变量
func test1()  {
	{
		a := 0
		fmt.Println(&a) // 0xc042008270
		a, b := 1,2
		fmt.Println(&a, &b) // 0xc042008270 0xc042008278
	}
	// 隐藏变量
	{
		x := 1
		fmt.Println(x)     // 1
		{
			fmt.Println(x) // 1
			x := 2
			fmt.Println(x) // 2
		}
		fmt.Println(x)     // 1
	}
}

// 不使用显式类型, 无法使用"nil"来初始化变量,  nil标志符用于表示interface,函数,maps,slices和channels的"零值",
// 如果不指定变量的类型, 编译器将无法编译代码, 因为猜不出具体的类型
func test2()  {
	// 在一个nil的slice中添加元素是没问题的,但对一个map做同样的事将会生成一个运行时的panic
	var s []int
	s = append(s,1) // right
	m1 := make(map[string]int)
	m1["one"] = 1

	var m2 map[string]int
	m2["one"] = 1 // error, panic: assignment to entry in nil map
}

func test3()  {
	m := make(map[string]int,99)
	// cap(m) // 你可以在map创建时指定它的容量，但你无法在map上使用cap()函数。
	fmt.Println(m)
}

// 字符串不会为nil
func test4()  {
	{
		var x string //defaults to "" (zero value)
		if x == "" {
			x = "default"
		}
		fmt.Println(x)
	}
	{
		//var x string = nil // error
		//if x == nil { // error
		//	x = "default"
		//}
		//fmt.Println(x)
	}
}

// Go中的数组是数值,因此当你向函数中传递数组时,函数会得到原始数组数据的一份复制
func test5()  {
	{
		x := [3]int{1,2,3}
		func(arr [3]int) {
			fmt.Println(&arr)
			arr[0] = 7
		}(x)
		fmt.Println(&x, x) //prints [1 2 3] (not ok if you need [7 2 3])
	}
	// 指针/slice(切片)解决这个问题
	{
		x := [3]int{1,2,3}
		func(arr *[3]int) {
			fmt.Println(arr)
			(*arr)[0] = 7
		}(&x)
		fmt.Println(&x, x) //prints [7 2 3]
	}
	{
		x := []int{1,2,3}
		func(arr []int) {
			fmt.Println(&arr)
			arr[0] = 7
		}(x)
		fmt.Println(&x, x) //prints [7 2 3]
	}
}

// 字符串不总是utf-8, 可以使用“unicode/utf8”包中的ValidString()函数
func test6()  {
	data1 := "ABC"
	fmt.Println(utf8.ValidString(data1)) // true
	data2 := "A\xfeC"
	fmt.Println(utf8.ValidString(data2)) // false
}

// 以小写字母开头的结构体将不会被(json,xml,gob等)编码
func test7()  {
	type MyData struct {
		One int
		two string
	}
	in := MyData{1,"two"}
	fmt.Printf("%#v\n",in) //prints main.MyData{One:1, two:"two"}
	encoded,_ := json.Marshal(in)
	fmt.Println(string(encoded)) //prints {"One":1}
	var out MyData
	json.Unmarshal(encoded,&out)
	fmt.Printf("%#v\n",out) //prints main.MyData{One:1, two:""}
}

// 关闭http的响应
// 当使用标准http库发起请求时,得到一个http的响应变量. 如果不读取响应主体, 依旧需要关闭它.
// 注意对于空的响应也一定要这么做
// resp, err := http.Get(url)
// 当http响应失败时, resp变量将为nil, 而err变量将是non-nil. 然而, 当你得到一个重定向的错误时,
// 两个变量都将是non-nil. 这意味着你最后依然会内存泄露
// 正确的做法如下
func test8() {
	resp, err := http.Get("https://api.ipify.org?format=json")
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

// 如果结构体中的各个元素都可以使用等号来比较的话, 那就可以使用 ==来比较结构体变量
// 如果结构体中的元素无法比较, 那使用等号将导致编译错误. 注意数组仅在它们的数据元素可比较的情况下才可以比较
// 最常用的方法是使用reflect包中的DeepEqual()函数
func test9()  {
	{
		type data struct {
			num int
			fp float32
			complex complex64
			str string
			char rune
			yes bool
			events <-chan string
			handler interface{}
			ref *byte
			raw [10]byte
		}
		v1 := data{}
		v2 := data{}
		fmt.Println("v1 == v2:", v1 == v2) //prints: v1 == v2: true
	}
	{
		type data struct {
			num int                //ok
			checks [10]func() bool //not comparable
			doit func() bool       //not comparable
			m map[string] string   //not comparable
			bytes []byte           //not comparable
		}
		v1 := data{}
		v2 := data{}
		// fmt.Println(v1, v2)
		// compile error, invalid operation: v1 == v2 (struct containing [10]func() bool cannot be compared)
		// fmt.Println("v1 == v2:",v1 == v2)
		fmt.Println("v1 == v2:", reflect.DeepEqual(v1, v2))
	}
}

// recover()的调用仅当它在defer函数中被直接调用时才有效。
func test10()  {
	defer func() {
		func() {
			fmt.Println("recovered =>",recover()) // panic is not recovered
		}()
	}()
	panic("not good")
}

// 在Slice, Array, Map "range"语句中更新引用元素的值
// 在"range"语句中生成的数据的值是真实集合元素的拷贝, 它们不是原有元素的引用
// 这意味着更新这些值将不会修改原来的数据, 同时也意味着使用这些值的地址将不会得到原有数据的指针
// 如果集合保存的是指针, 则会改变原来的值
func test11()  {
	{
		data := []*struct{ num int }{ {1}, {2}, {3} }
		for _, v := range data {
			v.num *= 10
		}
		fmt.Println(data[0], data[1], data[2]) // &{10} &{20} &{30}
	}
}

// 当重新划分一个slice时, 新的slice将引用原有slice的数组. 如果忘了这个行为的话, 在应用分配大量临时的slice用于创建新的slice
// 来引用原有数据的一小部分时, 会导致难以预期的内存使用
func test12()  {
	func () []byte {
		raw := make([]byte,10000)
		fmt.Println(len(raw),cap(raw),&raw[0]) //prints: 10000 10000 <byte_addr_x>
		return raw[:3] // 返回的数据是对raw的引用, 所以会造成空间的浪费
	}()
	// 正确的做法如下

	func() []byte {
		raw := make([]byte,10000)
		fmt.Println(len(raw),cap(raw),&raw[0]) //prints: 10000 10000 <byte_addr_x>
		res := make([]byte,3)
		copy(res,raw[:3])
		return res
	}()
}

// 多个slice可以引用同一个数据. 比如, 当从一个已有的slice创建一个新的slice时, 这就会发生。
// 如果你的应用功能需要这种行为, 那么你将需要关注下"走味的"slice.
// 在某些情况下, 在一个slice中添加新的数据, 在原有数组无法保持更多新的数据时, 将导致分配一个新的数组.
// 而现在其他的slice还指向老的数组
func test13()  {
	s1 := []int{1,2,3}
	fmt.Println(len(s1), cap(s1), s1) //prints 3 3 [1 2 3]
	s2 := s1[1:]
	fmt.Println(len(s2), cap(s2), s2) //prints 2 2 [2 3]
	for i := range s2 { s2[i] += 20 }
	//still referencing the same array
	fmt.Println(s1) //prints [1 22 23]
	fmt.Println(s2) //prints [22 23]

	s2 = append(s2,4)
	for i := range s2 {
		s2[i] += 10
	}
	//s1 is now "stale"
	fmt.Println(s1) //prints [1 22 23]
	fmt.Println(s2) //prints [32 33 14]
}

// 当通过把一个现有(非interface)的类型定义为一个新的类型时, "新的类型不会继承现有类型的方法"(可以使用结构体- -)
func test14()  {
	type myMutex sync.Mutex
	var mtx myMutex
	// mtx.Lock()   // error
	// mtx.Unlock() // error
	fmt.Println(mtx)
}

// for语句中的迭代变量在每次迭代时被重新使用. 这就意味着在for循环中创建的闭包将会引用同一个变量,
// 而在那些goroutine开始执行时就会得到那个变量的值
func test15()  {
	data := []string{"one","two","three"}
	for _,v := range data {
		// error
		go func() {
			fmt.Println(1, v)
		}()
	}

	for _,v := range data {
		// right
		go func(in string) {
			fmt.Println(2, in)
		}(v)
	}
	time.Sleep(2 * time.Second)
}

// Defer函数调用参数的求值, 被defer的函数的参数会在defer声明时求值 (而不是在函数实际执行时)
func test16()  {
	var i int = 1
	defer fmt.Println("result =>",func() int { return i * 2 }()) // 2 not 4
	i++
}
func main()  {
	test16()
}