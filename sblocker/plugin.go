package main

import (
	"../libs/common"
	"../libs/golang"
	"github.com/ddosakura/gklang"
)

// set config (dev/pkg/comment/algorithm)
// choose lang / load plugin
func loadPlugin() {
	cfg := &sbi.Config{
		Dev:       dev,
		PKG:       pkg,
		Comment:   comment,
		Algorithm: algorithm,
	}

	gklang.Log(gklang.LInfo, "loading")

	// TODO: load
	generater = sbgo.New(cfg)
}
