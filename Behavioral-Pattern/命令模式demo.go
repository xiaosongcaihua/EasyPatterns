package main

import "fmt"

//联想路边烧烤场景、有烤羊肉、烤鸡翅命令，有烤串师傅和服务员根据命令模式设计场景

type cookRoast interface {
	Roast()
}

type cook struct {
}

func (c *cook) roastLamb() {
	fmt.Println("厨师进行烤羊肉")
}
func (c *cook) roastCW() {
	fmt.Println("厨师进行烤鸡翅")
}

type Lamb struct {
	c *cook
}

func (l *Lamb) Roast() {
	l.c.roastLamb()
}

type ChickenW struct {
	c *cook
}

func (l *ChickenW) Roast() {
	l.c.roastCW()
}

type MM struct {
	cr []cookRoast
}

func (m *MM) Notify() {
	if m == nil {
		return
	}

	for _, roast := range m.cr {
		roast.Roast()
	}
}
func main() {
	ck := new(cook)
	l := &Lamb{c: ck}
	c := &ChickenW{c: ck}

	m := new(MM)
	m.cr = append(m.cr, l)
	m.cr = append(m.cr, c)
	m.Notify()
}
