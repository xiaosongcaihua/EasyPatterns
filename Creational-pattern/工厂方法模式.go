package main

import "fmt"

//<------抽象层------>

type AbstractFactory interface {
	CreateFruit() Fruit
}
type Fruit interface {
	Show()
}

//<-------实现层----->
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

type AppleFactory struct {
}

func (a *AppleFactory) CreateFruit() Fruit {
	return new(Apple)
}

type BananaFactory struct {
}

func (b *BananaFactory) CreateFruit() Fruit {
	return new(Banana)
}

/**
工厂方法模式诠释了开闭原则和单一职责原则
缺点：每次创建对象的时候都需要创建一个工厂
*/
func main() {
	af := new(AppleFactory)
	apple := af.CreateFruit()
	apple.Show()

	bf := new(BananaFactory)
	banana := bf.CreateFruit()
	banana.Show()
}
