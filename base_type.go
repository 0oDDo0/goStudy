/* https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.2.md */
package main

import (
	"fmt"
	"errors"
	"github.com/orcaman/concurrent-map"
)
/*
默认的行为:
	1 大写字母开头的变量是可导出的, 也就是其它包可以读取的, 是公有变量; 小写字母开头的就是不可导出的, 是私有变量;
	2 大写字母开头的函数也是一样, 相当于class中的带public关键词的公有函数；小写字母开头的就是有private关键词的私有函数。
*/

/* learn 内置类型 */

func func_bool()  {
	// 布尔值的类型为bool, 值是true或false, 默认为false
	{
		var a bool
		var b bool = true
		fmt.Println(a, b)
	}
}

func func_int()  {
	/*
	整数类型有无符号和带符号两种. 默认为0, Go同时支持int和uint, 这两种类型的长度相同, 但具体长度取决于不同编译器的实现;
	Go里面也有直接定义好位数的类型: rune, int8, int16, int32, int64和byte, uint8, uint16, uint32, uint64;
	其中rune是int32的别称, byte是uint8的别称;
	这些类型的变量之间不允许互相赋值或操作, 不然在编译时编译器会报错;
	 */
	{
		var a int8 = 0
		var b uint16 = 1
		// c := a + b // error, invalid operation: a + b (mismatched types int8 and uint16)
		fmt.Println(a, b)
	}
}

func func_string()  {
	/*
	字符串, 都是采用utf-8编码, ""或``(不是单引号, 是反引号)括住
	Go中字符串是不可变的, 字符串用一个两字节的数据结构存储字符串的指针和长度
	 */
	{
		var a string = "a"
		// ``括起的字符串为Raw字符串,即字符串在代码中的形式就是打印时的形式,
		// 它没有字符转义, 换行符将原样输出
		var b = `b\n`
		fmt.Println(a, b)
		// a[0] = 'c' // error, 字符串不可变
		// 字符串可通过+连接, 可以利用字符数组来改变字符串中的值
		s1 := a + b
		c := []byte(s1)  // 将字符串 s 转换为 []byte 类型
		c[0] = '0'
		s2 := string(c)  // 再转换回 string 类型
		fmt.Println(len(s2), s2)
		fmt.Println(s2[0:]) // 支持切片操作
	}
}
func func_error()  {
	/*
	Go内置有一个error类型, 专门用来处理错误信息.
	*/
	{
		err := errors.New("errors")
		if err != nil {
			fmt.Println(err)
		}
	}
}

func func_array()  {
	/*
	array, var name [n]type, n表示数组的长度, type表示存储元素的类型
	slice, var name []int
	*/
	{
		/*
		由于长度也是数组类型的一部分, 因此[3]int与[4]int是不同的类型;
		数组不能改变长度;
		数组之间的赋值传递的是指针, 会相互影响;
		当把一个数组作为参数传入函数的时候, 传入的是它的指针, 会影响外部变量的值;
		*/
		var arr [10]int  // 声明了一个int类型的数组
		arr1 := [...]int{4, 5, 6} // 可以省略长度而采用...的方式, Go会自动根据元素个数来计算长度
		arr[0] = 0      // 数组下标是从0开始的
		arr[1] = 9      // 赋值操作
		fmt.Printf("The 1  element is %d\n", arr[0])  // 获取数据，返回42
		fmt.Printf("The -1 element is %d\n", arr[9])  // 返回未赋值的最后一个元素，默认返回0
		fmt.Println(len(arr1), arr1)
		// 声明了一个二维数组,该数组以两个数组作为元素,其中每个数组中又有4个int类型的元素
		twoArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
		fmt.Println(twoArray)
	}
}

func func_slice()  {
	/*
	array, var name [n]type, n表示数组的长度, type表示存储元素的类型
	slice, var name []int, 没有长度的限制, slice的声明和array相似, 只是不需要长度
	slice, 在初始定义数组时,并不知道需要多大的数组,因此就需要"动态数组"
	slice并不是真正意义上的动态数组, 而是一个引用类型, slice总是指向array
	*/
	{
		var slice1 = []int{1,2} // 不添加长度的array
		fmt.Println(len(slice1), slice1)
		// slice可以从一个数组或一个已经存在的slice中再次声明
		// slice通过array[i:j]来获取, 其中i是数组的开始位置, j是结束位置, 但不包含array[j], 它的长度是j-i
		//  类似于Python中的切片
		arr := [...]int{1,2,3}
		slice2 := arr[0:3]
		/*
		slice是引用类型, 所以当引用改变其中元素的值时, 其它的所有引用都会改变该值
		slice2存储了指针, 指向了arr的地址, 并未复制arr中的值
		从概念上面来说slice像一个结构体, 这个结构体包含了三个元素:
			1) 一个指针, 指向数组中slice指定的开始位置
			2) 长度, 即slice的长度
			3) 最大长度, 也就是slice开始位置到数组的最后位置的长度, 相当于容量
		对于slice有几个有用的内置函数:
			1) len 获取slice的长度
			2) cap 获取slice的最大容量
			3) append 向slice里面追加一个或者多个元素, 然后返回一个和slice一样类型的slice
			4) copy 函数copy从源slice的src中复制元素到目标dst, 并且返回复制的元素的个数
			注:append函数会改变slice所引用的数组的内容, 从而影响到引用同一数组的其它slice.
			   但当slice中没有剩余空间(即(cap-len) == 0)时,此时将动态分配新的数组空间.
			   返回的slice数组指针将指向这个空间, 而原数组的内容将保持不变; 其它引用此数组的slice则不受影响.
		*/
		arr[0] = 4 // slice2的值也会改变
		slice2[1] = 5 // arr中的值也会改变
		fmt.Println(len(slice2), cap(slice2), slice2, arr)
		// 超过了最大长度, 重新申请内存, arr和slice2不会相互影响
		slice2 = append(slice2, 4)
		arr[0] = 6
		slice2[0] = 7
		fmt.Println(len(slice2), cap(slice2), slice2, arr)
	}
}

func func_map() {
	/*
	map是Python中的字典, 它的格式为map[keyType]valueType, 在使用前需要使用make初始化
	使用map过程中需要注意的几点:
	  1) map是无序的,每次打印出来的map都会不一样,它不能通过index获取, 必须通过key获取;
	  2) 和slice一样, map也是一种引用类型
	  3) 内置的len函数同样适用于map,返回map拥有的key的数量
	  4) map和其他基本类型不同,它不是thread-safe,在多个go-routine存取时,必须使用mutex lock机制
	*/
	{
		var amap = make(map[string]int)
		bmap := map[string]float32{"C":5, "Go":4.5, "Python":4.5, "C++":2 }
		var cmap = map[string]int{"a":1, "b":2}
		var dmap = map[string]int{}
		var emap map[string] int // 声明
		// emap["a"] = 1 // error, 如果要使用, 需要使用 make申请内存
		amap["a"] = 1 // map[key] = value
		fmt.Println(amap, bmap, cmap, dmap, emap)
		// map有两个返回值,(value, exist_flag)
		// 如果存在该key, exist_flag = true, value = map[key]
		// 如果不存在, exist_flag = false, value = 0(对应value类型的默认值, int默认为0)
		value, ok := bmap["C#"]
		if ok {
			fmt.Println(ok, value)
		} else {
			fmt.Println(ok, value)
		}
		delete(bmap, "C")  // 删除key为C的元素
		delete(bmap, "1C") // 删除不存在的键也不会报错
		fmt.Println(bmap)
		// map也是一种引用类型,如果两个map同时指向一个底层,那么一个改变,另一个也相应的改变
		fmap1 := make(map[string]string)
		fmap1["hello"] = "world1"
		fmap2 := fmap1
		fmap1["hello"] = "world2"
		fmt.Println(fmap1, fmap2)
	}
}

/*
线程安全map(加了锁的map)
将存储区域分片, 每片使用一把锁, 提升效率,
*/
func study()  {
	var dict cmap.ConcurrentMap = cmap.New()
	tmp := make(map[string]int)
	tmp["key1"] = 1
	tmp["key2"] = 2
	dict.Set("key1", "value1")
	dict.Set("key2", 2)
	dict.Set("key3", tmp)
	fmt.Println(dict.Get("key3"))
	fmt.Println(dict.GetShard("key3"))
}


func func_default()  {
	/*
	每种类型的默认值
	int     0
	int8    0
	int32   0
	int64   0
	uint    0x0
	rune    0 //rune的实际类型是 int32
	byte    0x0 // byte的实际类型是 uint8
	float32 0 //长度为 4 byte
	float64 0 //长度为 8 byte
	bool    false
	string  ""
	*/
}

func main() {
	func_string()
}