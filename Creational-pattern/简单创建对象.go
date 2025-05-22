package main

import "fmt"

type Fruit struct {
}

func NewFruit(name string) *Fruit {
	if name == "banana" {
		return new(Fruit)
	} else if name == "apple" {
		return new(Fruit)
	}
	return nil
}
func (f *Fruit) show(name string) {
	if name == "apple" {
		fmt.Println("apple show")
	} else if name == "banana" {
		fmt.Println("banana show")
	}
}

/**
简单创建对象的时候会违反两个设计原则
1、职责单一原则：当前水果有show和创建水果因此违反原则
2、开闭原则：当有新的水果产生的时候需要在修改原有代码，不规范
*/
func main() {
	f := NewFruit("apple")
	f.show("apple")
}
