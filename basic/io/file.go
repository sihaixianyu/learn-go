package io

import (
	"fmt"
	"os"
)

func FileRead() {
	f, err := os.OpenFile("../temp.txt", os.O_RDONLY, 0);
	if err != nil {
		panic(err)
	}
	
	buf := make([]byte, 1024)
	_, err = f.Read(buf)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(buf))
}