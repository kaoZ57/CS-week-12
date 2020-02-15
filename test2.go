package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	fileSize := stat.Size()
	p := make([]byte, fileSize)
	file.Read(p)

	str := string(p)
	fmt.Println(str)
}
