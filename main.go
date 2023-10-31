package main

import (
	"bookingsys/api/core/router"
	"fmt"
)

func main() {

	fmt.Println("hello world")

	r := router.InitializeRouter()
	_ = r.Run(":8080")
}
