package reflect

import (
	"fmt"
	"reflect"
)

type Author struct {
	Name         string `json:"name"`
	Publication string `json:"publication"`
}

func Demo() {
	t := reflect.TypeOf(Author{})
	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Name
		s, _ := t.FieldByName(name)
		fmt.Println(name + ":", s.Tag)
	}
}
