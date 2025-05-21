package main

import "fmt"

/**
		 依赖倒转原则
			实现层

			抽象层

			业务逻辑层
实现层和业务逻辑层只关心抽象层，降低代码的耦合性值。不依赖于具体的实现类而只依赖于接口，面向接口编程

里氏代换原则：鼓励用接口编程

*/

// 抽象层

type Driver interface {
	Driver(c Car)
}

type Car interface {
	Run()
}

// 实现层

type BenC struct {
}

func (b *BenC) Run() {
	fmt.Println("奔驰正在运行")
}

type BMW struct {
}

func (b *BMW) Run() {
	fmt.Println("宝马正在运行")
}

type ZhangSan struct {
}

func (z *ZhangSan) Driver(c Car) {
	fmt.Println("zhangsan 启动")
	c.Run()
}

type lisi struct {
}

func (z *lisi) Driver(c Car) {
	fmt.Println("lisi 启动")
	c.Run()
}

// 业务逻辑层
func main() {
	var zs Driver = new(ZhangSan)
	var bc Car = new(BenC)
	zs.Driver(bc)

	var ls Driver = new(lisi)
	var bm Car = new(BMW)
	ls.Driver(bm)
}
