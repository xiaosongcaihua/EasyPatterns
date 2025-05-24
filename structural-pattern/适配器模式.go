package main

import "fmt"

//适配的目标
type V5 interface {
	Use5V()
}

//业务类，依赖V5接口
type iPhone struct {
	v V5
}

func NewPhone(v V5) *iPhone {
	return &iPhone{v}
}

func (p *iPhone) Charge() {
	fmt.Println("Phone进行充电...")
	p.v.Use5V()
}

//被适配的角色，适配者
type V220 struct{}

func (v *V220) Use220V() {
	fmt.Println("使用220V的电压")
}

//电源适配器
type Adapter struct {
	v220 *V220
}

func (a *Adapter) Use5V() {
	fmt.Println("使用适配器进行充电")

	//调用适配者的方法
	a.v220.Use220V()
}

func NewAdapter(v220 *V220) *Adapter {
	return &Adapter{v220}
}

/**
优点：在使用适配器模式的时候想要做适配无需修改源代码，只需要新增适配类即可，满足开闭原则
*/
// ------- 业务逻辑层 -------
func main() {
	iphone := NewPhone(NewAdapter(new(V220)))

	iphone.Charge()
}
