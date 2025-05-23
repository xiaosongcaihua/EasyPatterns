package main

import "fmt"

// ======抽象层=====

type Phone interface {
	ShowPhone()
}
type phoneDecorate struct {
	phone Phone
}

func (p *phoneDecorate) ShowPhone() {

}

// ======实现层======

//抽象组件的concrete

type Iphone struct {
}

func (i *Iphone) ShowPhone() {
	fmt.Println("This is iPhone show")
}

type HuaWei struct {
}

func (h *HuaWei) ShowPhone() {
	fmt.Println("This is HuaWei show")
}

// concreteDecorate
type ScreenDecorate struct {
	phoneDecorate
}

func (s *ScreenDecorate) ShowPhone() {
	s.phone.ShowPhone()
	fmt.Println("新增屏幕")
}

type CaseDecorate struct {
	phoneDecorate
}

func (s *CaseDecorate) ShowPhone() {
	s.phone.ShowPhone()
	fmt.Println("新增手机壳")
}
func NewScreenDecorate(phone Phone) Phone {
	return &ScreenDecorate{phoneDecorate{phone: phone}}
}
func NewCaseDecorate(phone Phone) Phone {
	return &CaseDecorate{phoneDecorate{phone: phone}}
}

/**
装饰器模式优点：
1、装饰器模式比继承更加灵活
2、可以读类进行多次装饰
3、符合开闭原则
缺点
1、在构建对象的时候会产生很多的小对象
2、排查复杂
*/
func main() {
	var phone Phone
	phone = new(Iphone)

	var sd Phone
	sd = NewScreenDecorate(phone)
	sd.ShowPhone()

	var cd Phone
	cd = NewCaseDecorate(sd)
	cd.ShowPhone()

}
