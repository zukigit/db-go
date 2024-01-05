package main

import (
	"fmt"
	"zuki/main/dbutil"
)

func main() {
	fmt.Println("This is main")
	dbutil.DBconnect()
}
