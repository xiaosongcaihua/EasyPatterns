package main

import (
	"fmt"
)

type sickNote interface {
	Treat()
}

type doctor struct {
}

func (d *doctor) treatEyes() {
	fmt.Println("医生治疗眼睛")
}
func (d *doctor) treatLegs() {
	fmt.Println("医生治疗腿")
}

type TreatEyesSick struct {
	d *doctor
}

func (t *TreatEyesSick) Treat() {
	t.d.treatEyes()
}

type TreatLegsSick struct {
	d *doctor
}

func (t *TreatLegsSick) Treat() {
	t.d.treatLegs()
}

type Nurse struct {
	cmd []sickNote
}

func (n *Nurse) order() {
	if n.cmd == nil {
		return
	}
	for _, note := range n.cmd {
		note.Treat()
	}
}

/**
命令模式的优点
1、系统解耦合，请求者与接受请求者不存在引用关系，因此完全解耦，两者具备很好的独立性
2、满足开闭原则，新命令可以直接加入系统中，符合开闭原则
命令模式缺点：
1、可能会创建非常多的具体命令类

使用场景：
1、
*/
func main() {
	dc := new(doctor)
	ls := &TreatLegsSick{d: dc}
	es := &TreatEyesSick{d: dc}
	n := new(Nurse)
	n.cmd = append(n.cmd, ls)
	n.cmd = append(n.cmd, es)
	n.order()
}
