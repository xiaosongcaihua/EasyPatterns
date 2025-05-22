package main

import (
	"fmt"
	"sync"
)

type instance struct {
}

var ins *instance = new(instance)

func GetInstance() *instance {
	return ins
}
func (i *instance) doSomething() {
	fmt.Println("instance method do something")
}

var once sync.Once

func GetInstanceLazy() *instance {
	once.Do(func() {
		if ins == nil {
			ins = new(instance)
		}
	})
	return ins
}
func main() {
	ch := make(chan *instance, 100)
	for i := 1; i < 100; i++ {
		go func() {
			insz := GetInstanceLazy()
			ch <- insz
		}()
	}
	insz := <-ch
	for i := 0; i < 98; i++ {
		if insz != <-ch {
			fmt.Println("false")
			return
		}
	}
	fmt.Println("true")
}
