package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	for i := 1; i < 100001; i++ {
		c := &i
		//fmt.Println(i)
		fileName := "f" + strconv.Itoa(i) + ".txt"
		cFile, _ := os.Create(fileName)

		fmt.Println(i, "file was created", c, *c)
		defer cFile.Close()

		// прерывание счётчика на 1 милисекунду * N
		//time.Sleep(time.Millisecond * 10)
	}
	fmt.Println("count was done")
}
