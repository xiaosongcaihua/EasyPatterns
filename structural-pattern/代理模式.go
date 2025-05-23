package main

import (
	"fmt"
)

// =====抽象层======

type Shop interface {
	Buy(goods *Goods)
}

// ======实现层======
type Goods struct {
	fact  bool
	kinds string
}
type ChineseShop struct {
}

func (s *ChineseShop) Buy(goods *Goods) {
	fmt.Println("购买中国商品", goods.kinds)
}

type JapaneseShop struct {
}

func (j *JapaneseShop) Buy(goods *Goods) {
	fmt.Println("购买日本商品", goods.kinds)
}

type OverseaProxy struct {
	shop Shop
}

func (o *OverseaProxy) InFact(goods *Goods) bool {
	if goods.fact {
		fmt.Println("当前商品是真货", goods.kinds)
		return true
	}
	fmt.Println("当前商品是假货", goods.kinds)
	return false
}
func (o *OverseaProxy) shipping(goods *Goods) {
	fmt.Println("当前商品已经发货", goods.kinds)
}
func (o *OverseaProxy) Buy(goods *Goods) {
	if o.InFact(goods) {
		o.shop.Buy(goods)
		o.shipping(goods)
	}
}
func CreateProxy(shop Shop) Shop {
	return &OverseaProxy{shop: shop}
}

/**
代理模式的优点：
1、满足开闭原则，当于代码做增强的时候不需要在源代码基础上修改
2、解耦合，将代码增强和基础逻辑实现分开
缺点
1、代码实现负载
*/
func main() {
	g1 := &Goods{fact: true, kinds: "笔记本"}
	g2 := &Goods{fact: false, kinds: "nike"}
	var cs Shop
	cs = new(ChineseShop)
	cs.Buy(g1)
	cs.Buy(g2)
	var js Shop
	js = new(JapaneseShop)
	js.Buy(g2)
	js.Buy(g1)

	var op Shop
	op = CreateProxy(cs)
	op.Buy(g1)

}
