package main

import (
	"fmt"
	"math"
	"unsafe"
)

func BoolType() {
	// 定义变量
	ok := true
	fmt.Printf("ok: %v\n", ok)

	// 用于判断
	age := 17
	if age >= 18 {
		fmt.Println("已成年")
	} else {
		fmt.Println("未成年")
	}

	// 用于循环
	count := 10
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}

	// 用于逻辑表达式
	manage := 18
	gender := "男"
	if manage >= 18 && gender == "男" {
		fmt.Println("男性已成年")
	}
}

func NumberType() {
	// 整数类型
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	var ui8 uint8
	var ui16 uint16
	var ui32 uint32
	var ui64 uint64

	fmt.Printf("%T %dB %v~%v\n", i8, unsafe.Sizeof(i8), math.MinInt8, math.MaxInt8)
	fmt.Printf("%T %dB %v~%v\n", i16, unsafe.Sizeof(i16), math.MinInt16, math.MaxInt16)
	fmt.Printf("%T %dB %v~%v\n", i32, unsafe.Sizeof(i32), math.MinInt32, math.MaxInt32)
	fmt.Printf("%T %dB %v~%v\n", i64, unsafe.Sizeof(i64), math.MinInt64, math.MaxInt64)

	fmt.Printf("%T %dB %v~%v\n", ui8, unsafe.Sizeof(ui8), 0, math.MaxUint8)
	fmt.Printf("%T %dB %v~%v\n", ui16, unsafe.Sizeof(ui16), 0, math.MaxUint16)
	fmt.Printf("%T %dB %v~%v\n", ui32, unsafe.Sizeof(ui32), 0, math.MaxUint32)
	fmt.Printf("%T %dB %v~%v\n", ui64, unsafe.Sizeof(ui64), 0, uint64(math.MaxUint64))

	var imax int
	var imin int
	imax = int(math.MaxInt64) // 再+1会导致overflows错误
	imin = int(math.MinInt64) // 再-1会导致overflows错误
	fmt.Printf("%T %dB %v~%v\n", imax, unsafe.Sizeof(imax), imin, imax)

	var ui uint
	ui = uint(math.MaxUint64) // 再+1会导致overflows错误
	fmt.Printf("%T %dB %v~%v\n", ui, unsafe.Sizeof(ui), 0, ui)

	// 浮点数类型
	var f32 float32
	var f64 float64

	fmt.Printf("%T %dB %v~%v\n", f32, unsafe.Sizeof(f32), -math.MaxFloat32, math.MaxFloat32)
	fmt.Printf("%T %dB %v~%v\n", f64, unsafe.Sizeof(f64), -math.MaxFloat64, math.MaxFloat64)

	// 科学计数法表示
	var f1 float32 = 314e-2
	var f2 float64 = 314e+2
	fmt.Println("f1:", f1) // f1: 3.14
	fmt.Println("f2:", f2) // f2: 31400

	// 精度丢失的实例
	var f3 float32 = 256.000000098
	var f4 float64 = 256.000000098
	fmt.Println("f3:", f3) // f3: 256
	fmt.Println("f4:", f4) // f4: 256.000000098
}

func StringType() {

}

func main() {
	BoolType()
	NumberType()
	StringType()
}
