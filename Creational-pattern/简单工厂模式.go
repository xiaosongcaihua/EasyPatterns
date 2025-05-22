package main

import "fmt"

// <-----抽象层----->
type fruit interface {
	Show()
}

// <-----实现层----->
type Apple struct {
}

func (a *Apple) Show() {
	fmt.Println("this is apple show")
}

type Banana struct {
}

func (b *Banana) Show() {
	fmt.Println("this is banana show")
}

type EasyFactory struct {
}

func (f *EasyFactory) CreateF(name string) fruit {
	if name == "apple" {
		return new(Apple)
	} else if name == "banana" {
		return new(Banana)
	}
	return nil
}

/**
当前实现的是简单工厂模式
优点：相比直接创建对象而言，新增了抽象层，解决了单一职责问题
缺点：依然没有满足开闭原则，创建对象依然需需要修改代码
*/
// <-----业务逻辑层------>
func main() {
	f := new(EasyFactory)
	a := f.CreateF("apple")
	a.Show()

	b := f.CreateF("banana")
	b.Show()
}
