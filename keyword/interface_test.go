package keyword

import (
	"fmt"
	"testing"
)

func TestMethodInvoke(t *testing.T) {
	chinaCat := ChinaCat{
		name: "chinaCat",
	}
	chinaCat.Walk()
	chinaCat.Yell()

	var cat Cat = &ChinaCat{
		name: "cat",
	}
	cat.Walk()
	cat.Yell()
}

func TestTypeAssertion(t *testing.T) {
	var cat Cat = &ChinaCat{
		name: "cat",
	}

	_, ok := cat.(Animal)
	if ok {
		fmt.Println("This cat is an animal")
	} else {
		fmt.Println("Assert failted")
	}

	_, ok = cat.(*ChinaCat)
	if ok {
		fmt.Println("This cat is a china Cat")
	} else {
		fmt.Println("Assert failted")
	}

	var animal Animal = &ChinaCat{
		name: "cat",
	}

	_, ok = animal.(Cat)
	if ok {
		fmt.Println("This animal is cat")
	} else {
		fmt.Println("Assert failted")
	}

	_, ok = cat.(*ChinaCat)
	if ok {
		fmt.Println("This animal is a china cat")
	} else {
		fmt.Println("Assert failted")
	}
}
