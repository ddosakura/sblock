package sbgo

import (
	"../common"
)

// Plugin is the lib for golang
type Plugin struct {
	c *sbi.Config
}

// New is the entry of the plugin
func New(c *sbi.Config) *Plugin {
	return &Plugin{
		c,
	}
}

// Encode is the func for generater
func (*Plugin) Encode() {

}

// Example help user to know how to use the library
func (*Plugin) Example() {

}

// Decode is the func for your project
func (*Plugin) Decode() {

}
