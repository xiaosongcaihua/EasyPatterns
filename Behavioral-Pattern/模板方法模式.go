package main

import "fmt"

type Beverage interface {
	BoilWater()
	Brew()
	PourInCup()
	AddThings()
	WantAddThings() bool
}

type BTemplate struct {
	b Beverage
}

func (t *BTemplate) MakeBeverage() {
	if t == nil || t.b == nil {
		return
	}

	t.b.BoilWater()
	t.b.Brew()
	t.b.PourInCup()
	if t.b.WantAddThings() {
		t.b.AddThings()
	}
}

type CoffeeBeverage struct {
	BTemplate
}

func (b *CoffeeBeverage) BoilWater() {
	fmt.Println("将咖啡的水煮到100摄氏度")
}
func (b *CoffeeBeverage) Brew() {
	fmt.Println("将咖啡进行打磨")
}
func (b *CoffeeBeverage) PourInCup() {
	fmt.Println("将咖啡装进杯子中")
}
func (b *CoffeeBeverage) AddThings() {
	fmt.Println("添加咖啡小料")
}
func (b *CoffeeBeverage) WantAddThings() bool {
	return true
}
func NewCoffer() *CoffeeBeverage {
	cof := new(CoffeeBeverage)
	cof.b = cof
	return cof
}

type TeaBeverage struct {
	BTemplate
}

func (b *TeaBeverage) BoilWater() {
	fmt.Println("将茶的水煮到100摄氏度")
}
func (b *TeaBeverage) Brew() {
	fmt.Println("将茶进行打磨")
}
func (b *TeaBeverage) PourInCup() {
	fmt.Println("将茶装进杯子中")
}
func (b *TeaBeverage) AddThings() {
	fmt.Println("添加茶小料")
}
func (b *TeaBeverage) WantAddThings() bool {
	return false
}
func NewTea() *TeaBeverage {
	cof := new(TeaBeverage)
	cof.b = cof
	return cof
}

/**
模板方法模式的优点：
1、将代码做一个抽离，提高代码的复用性
缺点：
1、在定义模板时候考虑要全面，否则后面修改模板之后，基本所有的代码都要修改
*/
func main() {
	cof := NewCoffer()
	cof.MakeBeverage()
	fmt.Println("-----------")
	tea := NewTea()
	tea.MakeBeverage()
}
