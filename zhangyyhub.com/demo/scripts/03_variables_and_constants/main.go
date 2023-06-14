package main

import "fmt"

// VariableDeclaration 变量声明
func VariableDeclaration() {
	// 单个变量声明
	var name string
	var age int
	var ok bool

	// 多个变量声明
	var (
		class string
		sex   string
	)

	fmt.Printf("name: %v\n", name)   // name:
	fmt.Printf("age: %v\n", age)     // age: 0
	fmt.Printf("ok: %v\n", ok)       // ok: false
	fmt.Printf("class: %v\n", class) // class:
	fmt.Printf("sex: %v\n", sex)     // sex:
}

// VariableInitialization 变量初始化
func VariableInitialization() {
	// 初始化赋值
	var flag bool = true
	var path string = "/etc/"

	// 初始化赋值 - 类型推导
	var database = "mysql"
	var lang = "go"

	fmt.Printf("flag: %v\n", flag)         // flag: true
	fmt.Printf("path: %v\n", path)         // path: /etc/
	fmt.Printf("database: %v\n", database) // database: mysql
	fmt.Printf("lang: %v\n", lang)         // lang: go
}

// ShortVariableDeclaration 短变量说明
func ShortVariableDeclaration() {
	// 在函数内部，可以使用:=运算符对变量进行声明和初始化，该方法只使用于函数体内部
	mobile := "huawei"
	number := 1234567

	fmt.Printf("mobile: %v\n", mobile) // mobile: huawei
	fmt.Printf("number: %v\n", number) // number: 1234567
}

// AnonymousVariable 匿名变量
func AnonymousVariable() {
	// 使用下划线表示变量名称即匿名变量
	ok, _ := func() (bool, string) {
		return true, ""
	}()

	fmt.Printf("ok: %v\n", ok) // ok: true
}

// DefineConstant 定义常量
func DefineConstant() {
	// 声明单个常量
	const PI1 float64 = 3.14
	const PI2 = 3.1415926

	// 声明多个常量
	const (
		NUM1 = 12
		NUM2
		NUM3
	)

	// iota
	const (
		a0 = iota
		a1
		a2
		a3 = 100
		a4 = iota
	)

	fmt.Printf("PI1: %v\n", PI1)   // PI1: 3.14
	fmt.Printf("PI2: %v\n", PI2)   // PI2: 3.1415926
	fmt.Printf("NUM1: %v\n", NUM1) // NUM1: 12
	fmt.Printf("NUM2: %v\n", NUM2) // NUM2: 12
	fmt.Printf("NUM3: %v\n", NUM3) // NUM3: 12
	fmt.Printf("a0: %v\n", a0)     // a0: 0
	fmt.Printf("a1: %v\n", a1)     // a1: 1
	fmt.Printf("a2: %v\n", a2)     // a2: 2
	fmt.Printf("a3: %v\n", a3)     // a3: 100
	fmt.Printf("a4: %v\n", a4)     // a4: 4
}

func main() {
	// 变量声明
	VariableDeclaration()

	// 变量初始化
	VariableInitialization()

	// 短变量说明
	ShortVariableDeclaration()

	// 匿名变量
	AnonymousVariable()

	// 定义常量
	DefineConstant()
}
