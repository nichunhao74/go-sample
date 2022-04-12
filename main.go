package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("hello world")
	fmt.Println("今の時刻：" + t.String())
}
