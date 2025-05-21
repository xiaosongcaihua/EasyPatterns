package main

import "fmt"

type Banker struct {
}

func (b *Banker) Save() {
	fmt.Println("存款")
}
func (b *Banker) Transfer() {
	fmt.Println("转账")
}
func (b *Banker) Stock() {
	fmt.Println("股票")
}

/**
开闭原则：类的改动是通过增加代码进行的，而不是修改源代码
*/
func main() {
	// 在新增股票功能的时候是在原有的类中新增的，当项目比较庞大的时候不建议直接在类上修改，肯能有不必要的麻烦
	b := new(Banker)

	b.Save()
	b.Transfer()
	b.Stock()

	fmt.Println("-----------------")

	// 通过这种方式创建的对象，最大程度上满足了开闭原则，不修改原有的代码，只在原来有的代码上做新增
	s := new(SaveBusiness)
	t := new(TransferBusiness)
	st := new(StockBusiness)

	s.duBusiness()
	t.duBusiness()
	st.duBusiness()
}

type BankBusiness interface {
	duBusiness()
}

type SaveBusiness struct {
}
type TransferBusiness struct {
}
type StockBusiness struct {
}

func (s *SaveBusiness) duBusiness() {
	fmt.Println("存款")
}
func (t *TransferBusiness) duBusiness() {
	fmt.Println("转账")
}
func (s *StockBusiness) duBusiness() {
	fmt.Println("股票")
}
