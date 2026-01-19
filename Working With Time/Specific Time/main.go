package main

import (
	"fmt"
	"time"
)

func main() {
	birthday := time.Date(1990, time.March, 15, 14, 30, 0, 0, time.UTC)
	fmt.Println(birthday)
}
