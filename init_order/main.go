package main

import "fmt"

func singleNumber(nums []int) int {

	for i := range nums {
	innerLoop:
		for j := range nums {
			if nums[i] == nums[j+1] {
				break innerLoop
			}
		}
		return nums[i]
		// fmt.Println(nums[i])
		// break
	}
	return 111
}

func main() {
	nums := []int{2, 2, 1}
	fmt.Println(singleNumber(nums))
	// ch1 := make(chan int, 10)
	// ch2 := make(chan int, 10)
	// ch3 := make(chan int, 10)
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		ch1 <- i
	// 		ch2 <- i
	// 		ch3 <- i
	// 	}
	// }()
	// <-ch1
	// <-ch3
	// <-ch2
	// for i := 0; i < 1; i++ {
	// 	select {
	// 	case ch1 <- 111:
	// 		fmt.Printf("receive %d from channel 1\n")
	// 	case ch2 <- 111:
	// 		fmt.Printf("receive %d from channel 2\n")
	// 	case ch3 <- 111:
	// 		fmt.Printf("receive %d from channel 3\n")
	// 	}
	// }
}

// import (
// 	"fmt"
// 	"time"
// )

// func say(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(100 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

// func main() {
// 	go func() {
// 		fmt.Println("run goroutine in closure")
// 	}()
// 	go func(string) {
// 	}("gorouine: closure params")
// 	go say("in goroutine: world")
// 	say("hello")
// }

//import (

//	_ "github.com/learn/init_order/pkg1"
//)

// const mainName string = "main"

// var mainVar string = getMainVar()

//	func init() {
//		fmt.Println("main init method invoked")
//	}
// func main() {
// // 方式1
// for i := 0; i < 10; i++ {
//     fmt.Println("方式1，第", i + 1,"次循环")
// }

// // 方式2
// b := 1
// for b < 10 {
//     fmt.Println("方式2，第", b,"次循环")
// }

// // 方式3，无限循环
// ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
// var started bool
// var stopped atomic.Bool
// for {
//     if !started {
//         started = true
//         go func() {
//             for {
//                 select {
//                 case <-ctx.Done():
//                     fmt.Println("ctx done")
//                     stopped.Store(true)
//                     return
//                 }
//             }
//         }()
//     }
//     fmt.Println("main")
//     if stopped.Load() {
//         break
//     }
// }

// // 遍历数组
// var a [10]string
// a[0] = "Hello"
// for i := range a {
//     fmt.Println("当前下标：", i)
// }
// for i, e := range a {
//     fmt.Println("a[", i, "] = ", e)
// }

// // 遍历切片
// s := make([]string, 10)
// s[0] = "Hello"
// for i := range s {
//     fmt.Println("当前下标：", i)
// }
// for i, e := range s {
//     fmt.Println("s[", i, "] = ", e)
// }

// m := make(map[string]string)
// m["b"] = "Hello, b"
// m["a"] = "Hello, a"
// m["c"] = "Hello, c"
// for i := range m {
// 	fmt.Println("当前key：", &i)
// 	fmt.Println(i)
// }
// for k, v := range m {
// 	fmt.Println("m[", k, "] = ", v)
// }
// for {
// 中断select
// select {
// case <-time.After(time.Second * 2):
// 	fmt.Println("过了2秒")
// case <-time.After(time.Second * 3):
// 	fmt.Println("进过了1秒")
// 	if true {
// 		break
// 	}
// 	fmt.Println("break 之后")
// }
// // }
// var nilSlice []int
// fmt.Println("nilSlice length:", len(nilSlice))
// fmt.Println("nilSlice capacity:", cap(nilSlice))

// s2 := []int{9, 8, 7, 6, 5}
// fmt.Println("s2 length: ", len(s2))
// fmt.Println("s2 capacity: ", cap(s2))
// str := "hello, 123, 你好"
// var bytes []byte = []byte(str)
// var runes []rune = []rune(str)
// fmt.Printf("bytes 的值为: %v \n", bytes)
// fmt.Printf("runes 的值为: %v \n", runes)

// str2 := string(bytes)
// str3 := string(runes)
// fmt.Printf("str2 的值为: %v \n", str2)
// fmt.Printf("str3 的值为: %v \n", str3)

// }

// func main() {
// 	// 十六进制
// 	// var a uint8 = 0xF
// 	// var b uint8 = 0xf

// 	// // 八进制
// 	// var c uint8 = 017
// 	// var d uint8 = 0o17
// 	// var e uint8 = 0o17

// 	// // 二进制
// 	// var f uint8 = 0b1111
// 	// var g uint8 = 0b1111

// 	// // 十进制
// 	// var h uint8 = 15
// 	// fmt.Println(a == b)
// 	// fmt.Println(c == b)
// 	// fmt.Println(c == d)
// 	// fmt.Println(d == e)
// 	// fmt.Println(e == f)
// 	// fmt.Println(f == g)
// 	// fmt.Println(h == g)
// 	// var float1 float32 = 10
// 	// float2 := 10.0
// 	// float1 = float32(float2)
// 	// fmt.Println(float1)
// 	// var c1 complex64
// 	// c1 = 1.10 + 0.1i
// 	// c2 := 1.10 + 0.12i
// 	// c3 := complex(1.10, 0.1)
// 	// fmt.Println(c1 == complex64(c2))
// 	// fmt.Println(c2 == complex128(c1))
// 	// fmt.Println(c3 == c2)
// 	// x, m := real(c2), 1
// 	// y := imag(c2)
// 	// fmt.Println(x)
// 	// fmt.Println(y)
// 	// fmt.Println(m)
// 	// var s string = "Hello, world!"
// 	// var bytes []byte = []byte(s)
// 	// fmt.Println("convert \"Hello, world!\" to bytes: ", bytes)
// 	// fmt.Println(string(bytes) == s)
// 	// var aa uint8 = 1
// 	// var bb byte = 1
// 	// fmt.Println(aa == bb)
// 	// var r1 rune = 'a'
// 	// var r2 rune = '世'
// 	// fmt.Println(r1)
// 	// fmt.Println(r2)
// 	// var s string = "abc，你好，世界！"
// 	// var runes []rune = []rune(s)
// 	// fmt.Println(s)
// 	// fmt.Println(runes)
// 	// fmt.Println(len(runes))
// 	// 	var s string = "Hello\nworld!\n"
// 	// 	var s1 string = `Hello
// 	// world!
// 	// `
// 	// 	fmt.Println(s == s1)
// 	// var s string = "Go语言"
// 	// var bytes []byte = []byte(s)
// 	// var runes []rune = []rune(s)

// 	// fmt.Println("string length: ", len(s))
// 	// fmt.Println("bytes length: ", len(bytes))
// 	// fmt.Println("runes length: ", len(runes))

// 	// fmt.Println("string sub: ", s[0:7])
// 	// fmt.Println("bytes sub: ", string(bytes[0:7]))
// 	// fmt.Println("runes sub: ", string(runes[0:3]))
// 	// var digit int
// 	// var s string
// 	// var b bool
// 	// fmt.Println(digit)
// 	// fmt.Println(s)
// 	// fmt.Println(b)
// 	// var s1 string = "Hello"
// 	// var zero int
// 	// var b1 = true

// 	// var (
// 	// 	i  int = 123
// 	// 	b2 bool
// 	// 	s2 = "test"
// 	// )

// 	// var (
// 	// 	group = 2
// 	// )

// }

// func getMainVar() string {
// 	fmt.Println("main.getMainVar method invoked!")
// 	return mainName
// }
