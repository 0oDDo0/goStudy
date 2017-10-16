package main

import (
	"fmt"
	"encoding/json"
	"reflect"
)

/*
定义结构体
type struct_name struct {
   member1 definition;
   member2 definition;
   ......
   member3 definition;
}
定义变量
variable_name := struct_name {value1, value2, ..., valuen}
*/
type Books struct {
	title   string
	author  string
	subject string
	book_id int
}
func testBooks()  {
	/* 声明 Book1 为 Books 类型 */
	var Book1 Books
	/* 声明 Book2 为 Books 类型, 按照顺序提供初始化值 */
	var Book2 = Books{"Python", "www.python.com", "Python 教程", 2}
	/* 通过field:value的方式初始化, 这样可以任意顺序 */
	Book3 := Books{title:"c++", book_id:3}
	/* 通过new函数分配一个指针, 类型为*Books */
	var Book4 *Books = new(Books)
	/* book 1 描述 */
	Book1.title = "Go"
	// Book1.author = "www.runoob.com" // 未赋值的话使用默认值
	Book1.subject = "Go 教程"
	Book1.book_id = 1
	/* 打印 Book1 信息, 实例和指针都是用.来获取属性值 */
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 2 title : %s\n", Book2.title)
	var ptr *Books = &Book1
	printBook(ptr)
	printBook(&Book3)
	printBook(Book4)
}
func printBook(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

/*
标签(Tag), 在转换成其它数据格式的时候, 会使用其中特定的字段作为键值
json tag, 转换成json格式时, 使用指定的json名称
bson tag，转换成bson格式时, 使用指定的bson名称
*/
func testStructTag()  {
	type User struct {
		UserId   int    `json:"json_id" bson:"bson_id"`
		UserName string `json:"json_name" bson:"bson_name"`
	}
	// 输出json格式
	u := &User{UserId: 1, UserName: "tony"}
	j, _ := json.Marshal(u)
	fmt.Println(string(j))
	// 输出内容：{"user_id":1,"user_name":"tony"}
	// 获取tag中的内容
	t := reflect.TypeOf(u)
	field := t.Elem().Field(0)
	fmt.Println(field.Tag.Get("json")) // json_id
	fmt.Println(field.Tag.Get("bson")) // bson_id
}
/*
匿名字段, 所有的内置类型和自定义类型都是可以作为匿名字段的
当匿名字段是struct时,这个struct所拥有的全部字段都被隐式地引入了当前定义的struct
*/
type Human struct {
	name string
	age int
	weight int
}

type Student struct {
	Human  // 匿名字段, Student包含了Human的所有字段
	int    // 基础类型的匿名字段
	// int // 不能同时出现多个相同的匿名字段
	grade string
	name string // 与Human中的name重名, 优先访问最外面的属性
}

func testStudents()  {
	// 结构体中初始化结构体
	student := Student{Human{"Mark", 25, 120}, 0, "6-1", "liujunsheng"}
	// 修改匿名字段
	student.int = 1
	student.Human = Human{"ljs", 19, 60}
	fmt.Println(student)
	// 匿名字段的访问
	fmt.Println(student.Human, student.Human.name, student.int)
}
func main() {
	// testBooks()
	// testStudents()
	testStructTag()
}

