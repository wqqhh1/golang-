package main

import (
	"bytes"
	"fmt"
	"strings"
)
import "hello/user"

type WebSite struct { //结构体定义，对比java的类
	Name string
}

func main() {
	a, b := 1, 2
	fmt.Println(a + b)
	hello := user.Hello()
	fmt.Printf("hello: %v\n", hello)
	var name string
	var age int
	var phone float64
	var flag bool
	name = "123"
	age = 123
	phone = 18616112085
	flag = true
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(phone)
	fmt.Println(flag)
	if age >= 123 {
		fmt.Println("你真是个老不死的")
	} else {
		fmt.Println("你还不够老")
	}
	for i := 0; i <= 10; i++ {
		fmt.Printf("i=%v\n", i)
	}
	s1 := `
	line1
	line2
	line3`
	fmt.Printf("s1:%v\n", s1)

	s2 := "wqq"
	s3 := 23
	name1 := "wqq"
	age1 := "24"
	msg := fmt.Sprintf("name=%s,age=%v", s2, s3)
	msg2 := strings.Join([]string{name1, age1}, ",")
	fmt.Printf("msg:%v\n", msg)
	fmt.Printf("msg2:%v\n", msg2)

	//效率最高，直接写到缓存区里
	println()
	println("========直接写入缓存区（最快）========")
	var buffer bytes.Buffer
	buffer.WriteString("wqq")
	buffer.WriteString("，已经")
	buffer.WriteString("23岁了!")
	fmt.Printf("buffer.String(): %v\n", buffer.String())

	//字符串的切片操作
	println()
	println("========字符串的切片操作========")
	s := "hello world"
	n := 2
	m := 5
	fmt.Printf("s[n]: %v\n", s[n]) //表示对应索引位置上的字符的整数表示 ASCII🐎
	fmt.Printf("s[n:m]: %v\n", s[n:m])
	fmt.Printf("s[n:]: %v\n", s[n:])
	fmt.Printf("s[:m]: %v\n", s[:m])

	//字符串函数
	println()
	println("==========字符串函数=========")
	l := len(s)
	println(l)
	fmt.Printf("strings.Split(s,\" \"): %v\n", strings.Split(s, " ")) //分割
	println(strings.Contains(s, "world"))                             //是否包含
	println(strings.ToUpper(s))                                       //转换为大写
	println(strings.ToLower(strings.ToUpper(s)))                      //转换为小写
	println(strings.HasPrefix(s, "hello"))                            //是否以什么开头
	println(strings.HasSuffix(s, "world"))                            //是否以什么结尾
	println(strings.Index(s, "wo"))                                   //查找字串对应的索引，返回字串头的索引位置

	//格式化输出,结构体定义（定义在main函数外面）
	println()
	println("==========格式化输出=========")
	site := WebSite{Name: "wqq"}
	fmt.Printf("site: %v\n", site)  //%v 相应值的默认输出
	fmt.Printf("site: %#v\n", site) //%#v 相应值的Go语法表示，完整结构
	fmt.Printf("site: %T\n", site)  //%T 代表类型，相应值的类型
	flag1 := true
	fmt.Printf("falge1: %v\n", flag1) //%t 布尔类型的输出，用%v也可以的
	/*
		%b:二进制输出
		%c：对应的字符输出
		%d：十进制输出
		%o:八进制输出
		%q：单引号围绕的字符字面值，由Go语法安全的转义
		%x：十六进制输出，字母为小写字母表示a-f
		%X：十六进制输出，字母为大写字母表示A-F
		%U:Unicode格式。
		%p:指针类型输出，输出为地址
	*/

	//++和--在g里面不能运用到表达式里面，注意一下Java之类的可以
	num := 100
	num1 := 101
	r := num ^ num1
	fmt.Printf("r: %v\n", r) //异或运算
	num++
	println(num)
	num--
	println(num)

	println()
	println("=======for range语法=========") //和foreach差不多
	x := [...]int{1, 2, 3, 4, 5, 6, 7}
	for a, v := range x { //a是索引，v是值，不想使用变量的话可以用 ‘_’ 替代，这样不会报未使用变量的错
		println(a)
		println(v)
	}
	for _, v := range x {
		println(v)
	}
	println()
	println("=======if语法=========")
	if a < b { //其他语言的非0可以表示真，0可以表示假。  go语言不行，注意！！！ 并且必须加大括号
		println("wqq")
	}

	if height := 20; height > 19 { //go的if语句里面可以定义变量，用分号 ‘;' 分隔，但height的作用域只在if语句里面，在外面为未定义的存在
		println("h1wqq")
	}

}
