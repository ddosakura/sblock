package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ddosakura/gklang"
)

var (
	dev = false
)

func main() {
	l := log.New(os.Stdout, "[SBlocker]: ", log.LstdFlags)
	if os.Getenv("SBlockerDev") == "true" {
		dev = true
	}
	if dev {
		gklang.LoadLogger(l, gklang.LvDebug)
	} else {
		gklang.LoadLogger(l, gklang.LvInfo)
	}
	fmt.Println(logo)

	execute()
}
