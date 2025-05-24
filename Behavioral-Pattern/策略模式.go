package main

import "fmt"

//武器策略(抽象的策略)
type WeaponStrategy interface {
	UseWeapon() //使用武器
}

//具体的策略
type Ak47 struct{}

func (ak *Ak47) UseWeapon() {
	fmt.Println("使用Ak47 去战斗")
}

//具体的策略
type Knife struct{}

func (k *Knife) UseWeapon() {
	fmt.Println("使用匕首 去战斗")
}

//环境类
type Hero struct {
	strategy WeaponStrategy //拥有一个抽象的策略
}

//设置一个策略
func (h *Hero) SetWeaponStrategy(s WeaponStrategy) {
	h.strategy = s
}

func (h *Hero) Fight() {
	h.strategy.UseWeapon() //调用策略
}

/**
策略模式的优点
1、满足开闭原则，当新增一个策略的时候不用在源代码上进行修改，只需要新增代码
2、避免多重语句if-else,通过设置不同策略直接传递不同的策略即可
3、当前创建的策略可能会被复用
策略模式的缺点
1、维护复杂
*/
func main() {
	hero := Hero{}
	//更换策略1
	hero.SetWeaponStrategy(new(Ak47))
	hero.Fight()

	hero.SetWeaponStrategy(new(Knife))
	hero.Fight()
}
