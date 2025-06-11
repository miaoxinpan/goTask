package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("基础一任务二；")
	//指针
	num := 10
	fmt.Println("Before entering the function: Pointer1", num)
	Pointer1(&num)
	fmt.Println("After entering the function: Pointer1", num)

	nums := []int{2, 4, 6, 8, 10, 12, 14, 16, 18}
	fmt.Println("Before entering the function: Pointer2 ", nums)
	Pointer2(&nums)
	fmt.Println("After entering the function: Pointer2", nums)

	//	Goroutine
	go Goroutine1()
	go Goroutine2()
	time.Sleep(time.Second)
	TaskScheduler([]func(){
		func() {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("任务1完成")
		},
		func() {
			time.Sleep(300 * time.Millisecond)
			fmt.Println("任务2完成")
		},
		func() {
			time.Sleep(800 * time.Millisecond)
			fmt.Println("任务3完成")
		},
	})

	//面向对象
	var rect Rectangle
	rect.Area()
	rect.Perimeter()
	var cir Circle
	cir.Area()
	cir.Perimeter()

	var empl Employee = Employee{Person: Person{Name: "Tom", Age: 18}, EmployeeID: 1}
	empl.PrintInfo()

	//	Channel
	go Channel1Input()
	go Channel1Output()
	time.Sleep(3 * time.Second)
	go Channel2Input()
	go Channel2Output()
	time.Sleep(10 * time.Second)

	for i := 1; i < 11; i++ {
		go Sync1()
	}
	time.Sleep(10 * time.Second)
	fmt.Println("counter value:", counter)
	for i := 1; i < 11; i++ {
		go Sync2()
	}
	time.Sleep(10 * time.Second)
	fmt.Println("counter value:", counter64)
}

/* 指针
题目1 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
	考察点 ：指针的使用、值传递与引用传递的区别。
题目2 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
	考察点 ：指针运算、切片操作。
*/

/*--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------*/
/*Pointer*/
func Pointer1(addr *int) {
	*addr += 10
}
func Pointer2(addr *[]int) {
	for i := 0; i < len(*addr); i++ {
		(*addr)[i] *= 2
	}
}

/*--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------*/
/* Goroutine
题目1 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
	考察点 ： go 关键字的使用、协程的并发执行。
*/
/* Goroutine*/
func Goroutine1() {
	fmt.Println("Print odd numbers from 1 to 10")
	for i := 1; i < 11; i += 2 {
		fmt.Println("odd:", i)
	}
	fmt.Println("Goroutine is over--odd ")
}
func Goroutine2() {
	fmt.Println("Print even numbers from 1 to 10")
	for i := 0; i < 11; i += 2 {
		fmt.Println("even:", i)
	}
	fmt.Println("Goroutine is over--even ")
}

/*
题目2 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
	考察点 ：协程原理、并发任务调度。

	没有思路 暂时不做
	AI生成的
*/
//任务调度器
func TaskScheduler(tasks []func()) {
	//创建一个等待协程数
	var wg sync.WaitGroup
	//循环func切片
	for i, task := range tasks {
		//每有一个func  往wg 加1  代表有一个协程任务需要等待
		wg.Add(1)
		//这两个参数就是下面的(i, task)
		go func(idx int, t func()) {
			//defer 等执行完再执行defer后的代码
			//wg.Done() 关闭一个协程任务
			defer wg.Done()
			start := time.Now()
			//执行业务逻辑
			t()
			elapsed := time.Since(start)
			fmt.Printf("任务%d执行时间: %v\n", idx+1, elapsed)
		}(i, task)
	}
	//等待所有的并发任务执行完  再往下执行其他的任务
	wg.Wait()
}

/*--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------*/
/* 面向对象
题目1 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
	在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
	考察点 ：接口的定义与实现、面向对象编程风格。
题目2 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，
	再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
	为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
	考察点 ：组合的使用、方法接收者
*/
/* 面向对象*/
type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
}
type Circle struct {
}

// 结构体实现接口方法
func (r Rectangle) Area() {
	fmt.Println("Area function....Rectangle")
}
func (r Rectangle) Perimeter() {
	fmt.Println("Perimeter function....Rectangle")
}
func (c Circle) Area() {
	fmt.Println("Area function....Circle")
}
func (c Circle) Perimeter() {
	fmt.Println("Perimeter function....Circle")
}

type Person struct {
	Name string
	Age  int
}
type Employee struct {
	Person
	EmployeeID int
}

func (e Employee) PrintInfo() {
	fmt.Println("Employee Name:", e.Person.Name)
	fmt.Println("Employee Age:", e.Person.Age)
	fmt.Println("Employee EmployeeID:", e.EmployeeID)
}

/*--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------*/
/* Channel
题目1 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，
	另一个协程从通道中接收这些整数并打印出来。
	考察点 ：通道的基本使用、协程间通信。
题目2 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
	考察点 ：通道的缓冲机制。
*/
var channel1 chan int = make(chan int)

func Channel1Input() {
	for i := 1; i < 11; i++ {
		channel1 <- i
		fmt.Println("放入channel1的值:", i)
	}
}
func Channel1Output() {
	for {
		i := <-channel1
		fmt.Println("取出channel1的值:", i)
	}
}

var channel2 chan int = make(chan int, 10)

func Channel2Input() {
	for i := 1; i < 101; i++ {
		channel2 <- i
		fmt.Println("放入channel2的值:", i)
	}
}
func Channel2Output() {
	for {
		i := <-channel2
		fmt.Println("取出channel2的值:", i)
	}
}

/*--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------*/
/* 锁机制
题目1 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	考察点 ： sync.Mutex 的使用、并发数据安全。
题目2 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	考察点 ：原子操作、并发数据安全。
*/

var counter int
var mu sync.Mutex

func Sync1() {
	for i := 1; i < 1001; i++ {
		mu.Lock()
		counter++
		mu.Unlock()
	}
}

var counter64 int64

func Sync2() {
	for i := 1; i < 1001; i++ {
		atomic.AddInt64(&counter64, 1)
	}
}
