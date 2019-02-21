package main

import (
	"fmt"

	"github.com/ddosakura/gklang"
)

var (
	dev = false
)

func main() {
	gklang.Init("SBlocker")
	dev = gklang.DevMode

	fmt.Println(logo)

	execute()
}
