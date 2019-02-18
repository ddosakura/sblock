package main

import (
	p ".."
	"../../common"
)

// New is the entry of the plugin
func New(c *sbi.Config) sbi.Plugin {
	return p.New(c)
}

func main() {}
