package main

import "fmt"

type Cat struct {
}

func (c *Cat) eat() {
	fmt.Println("cat eat something")
}

type CatA struct {
	Cat
}

//
//func (a *CatA) eat() {
//	fmt.Println("eat aaa")
//}
func (a *CatA) Sleep() {
	fmt.Println("cat sleep")
}

type CatB struct {
	C Cat
}

func (b *CatB) eat() {
	b.C.eat()
}
func (a *CatB) Sleep() {
	fmt.Println("cat sleep")
}

/**
对于类的功能扩充优先采用组合而不是继承，组合的依赖更小，可以在方法入参或者内部变量组合
*/
func main() {

	ca := new(CatA)
	ca.eat()
	ca.Sleep()

	fmt.Println("-----------")

	cb := new(CatB)
	cb.C.eat()
	cb.Sleep()
}
