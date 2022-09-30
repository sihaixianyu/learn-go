package keyword

import "fmt"

type Animal interface {
	Eat()
}

type Cat interface {
	Walk()
	Yell()
}

type ChinaCat struct {
	name string
}

func (c *ChinaCat) Eat() {
	fmt.Printf("%s is eating!\n", c.name)
}

// 这里如果初始化为值类型，那么该类型是没有实现Cat接口的，因为T类型不持有*T类型的方法
func (c ChinaCat) Walk() {
	fmt.Printf("%s is walking!\n", c.name)
}

// 这里如果初始化为值类型，那么该类型是实现了Cat接口的，因为*T类型持有T类型的方法
func (c *ChinaCat) Yell() {
	fmt.Printf("%s is quacking!\n", c.name)
}
