package main

import (
	"fmt"
	"os"

	"../libs/common"
	"github.com/ddosakura/gklang"
)

// test src/target/example with force
func cleanTarget() {
	clean(target)
	clean(example)
}

func clean(dest string) {
	var err error
	if _, err = os.Stat(dest); !os.IsNotExist(err) {
		if force {
			if err = os.RemoveAll(dest); err != nil {
				gklang.Log(gklang.LCrash, err, fmt.Errorf("file %q could not be deleted", dest))
			}
		} else {
			gklang.Er(fmt.Errorf("file %q already exists; use -f to overwrite", dest))
		}
	}
}

var (
	generater sbi.Plugin
)

// generate code
func generate() {
	gklang.Log(gklang.LInfo, "generating")
	// TODO: generate
}
