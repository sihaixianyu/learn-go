package io

import (
	"bufio"
	"fmt"
	"os"
)

func ReadMainFile() {
	f, err := os.Open("../main.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
}
