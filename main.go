package main

import (
	"fmt"
	"time"
)

func main() {
	//Parse YYYY-MM-DD
	x := "2020-01-29"
	timeT, _ := time.Parse("2006-01-02", x)
	//timeT, _ := time.Parse(time.RFC3339, x)
	
	fmt.Println(timeT)

}
