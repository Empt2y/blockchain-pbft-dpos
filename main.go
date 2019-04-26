package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2019, 4, 26, 15, 30, 0, 0, time.Local).Unix()
	fmt.Println(t, time.Now().Unix())
}
