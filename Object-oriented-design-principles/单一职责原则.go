package main

import "fmt"

type Clothes struct {
}

func (c *Clothes) onShop() {
	fmt.Println("逛街衣服")
}
func (c *Clothes) onWork() {
	fmt.Println("工作衣服")
}

/**
单一职责原则：类的职责单一，对外只提供一种功能，从而引起变化的原因只有一个
在后面开发过程中只能尽量满足，避免创建太多类，要求高内聚低耦合
*/
func main() {
	// 当前对象设计不满足单一职责原则，因为这个衣服相当于有两个职责一个是逛街衣服，另一个是工作衣服
	c := new(Clothes)

	fmt.Println("在工作")
	c.onWork()

	fmt.Println("在逛街")
	c.onShop()

	fmt.Println("---------------")
	// 下方写法满足单一职责原则，一个对象只有一个职责
	wc := new(WorkClothes)
	sl := new(ShopClothes)

	fmt.Println("在工作")
	wc.style()

	fmt.Println("在逛街")
	sl.style()
}

type WorkClothes struct {
}
type ShopClothes struct {
}

func (sl *ShopClothes) style() {
	fmt.Println("逛街衣服")
}
func (wl *WorkClothes) style() {
	fmt.Println("工作衣服")
}
