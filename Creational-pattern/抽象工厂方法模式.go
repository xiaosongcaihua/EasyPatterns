package main

import "fmt"

// <--------抽象层-------->

type AbstractApple interface {
	ShowApple()
}
type AbstractBanana interface {
	ShowBanana()
}

type AbstractCountryFactory interface {
	CreateApple() AbstractApple
	CreateBanana() AbstractBanana
}

// <-------实现层-------->

type ChineseFactory struct {
}

func (c *ChineseFactory) CreateApple() AbstractApple {
	return new(CApple)
}
func (c ChineseFactory) CreateBanana() AbstractBanana {
	return new(CBanana)
}

type JapaneseFactory struct {
}

func (j *JapaneseFactory) CreateApple() AbstractApple {
	return new(JApple)
}
func (j JapaneseFactory) CreateBanana() AbstractBanana {
	return new(JBanana)
}

type CApple struct {
}

func (a *CApple) ShowApple() {
	fmt.Println("chinese apple")
}

type CBanana struct {
}

func (a *CBanana) ShowBanana() {
	fmt.Println("chinese banana")
}

type JApple struct {
}

func (a *JApple) ShowApple() {
	fmt.Println("japanese apple")
}

type JBanana struct {
}

func (a *JBanana) ShowBanana() {
	fmt.Println("japanese banana")
}

/**
相比于工厂方法模式，多了一个产品结构的观念 ，例如cpu、内存、显卡相当于是产品结构，amd的cpu、inter的cpu、英伟达的cpu相当于
不同的族，对于抽象工厂方法适用于结构稳定的场景，可以创建一个工厂实现当前工厂内所有内容的创建，如果是用工厂方法模式，那么如果要获取
amd的cpu、内存分别要创建两个工厂，相对于当前场景更为适合
*/
func main() {
	var cf AbstractCountryFactory
	cf = new(ChineseFactory)
	var cbf AbstractBanana
	cbf = cf.CreateBanana()
	cbf.ShowBanana()
	var caf AbstractApple
	caf = cf.CreateApple()
	caf.ShowApple()

	var jf AbstractCountryFactory
	jf = new(JapaneseFactory)
	var jbf AbstractBanana
	jbf = jf.CreateBanana()
	jbf.ShowBanana()
	var jaf AbstractApple
	jaf = jf.CreateApple()
	jaf.ShowApple()
}
