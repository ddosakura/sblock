package main

import (
	"errors"
	"plugin"

	"../plugins/golang"
	"../sbi"
	"github.com/ddosakura/gklang"
)

var (
	errUnknowLang = errors.New("Unknow Language")
)

// set config (dev/pkg/comment/algorithm)
// choose lang / load plugin
func loadPlugin() {
	cfg := &sbi.Config{
		PKG:       pkg,
		Comment:   comment,
		Algorithm: algorithm,
		Params:    params,
	}

	if mod == "" {
		switch lang {
		case "golang":
			generater = sbgo.New(cfg)
		default:
			gklang.Er(errUnknowLang)
		}
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
