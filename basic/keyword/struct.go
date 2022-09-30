package keyword

import "fmt"

type Student struct {
	name  string
	class int
	age   int
}

func (p Student) String() string {
	return fmt.Sprintf("%s: {class: %d, age: %d }", p.name, p.class, p.age)
}

func (p Student) Upgrade() {
	p.class += 1
}

func (p *Student) Grow() {
	p.age += 1
}
