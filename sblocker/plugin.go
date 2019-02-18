package main

import (
	"plugin"

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

	if mod == "" {
		// TODO: choose lang
		generater = sbgo.New(cfg)
		return
	}

	gklang.Log(gklang.LInfo, "loading")

	pdll, err := plugin.Open(mod)
	if err != nil {
		gklang.Er(err)
	}

	newPlugin, err := pdll.Lookup("New")
	if err != nil {
		gklang.Er(err)
	}

	generater = newPlugin.(func(c *sbi.Config) sbi.Plugin)(cfg)
}
