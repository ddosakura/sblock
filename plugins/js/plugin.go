package sbjs

import (
	"../../sbi"
)

const (
	sourceFileName  = "raw.js"
	exampleFileName = "example.js"
)

var (
	// ErrUnknowAlgorithm in sbi
	ErrUnknowAlgorithm = sbi.ErrUnknowAlgorithm
)

// Plugin is the lib for js
type Plugin struct {
	c *sbi.Config
}

// SourceFileName return the FileName of the source
func (*Plugin) SourceFileName() string {
	return sourceFileName
}

// ExampleFileName return the FileName of the example
func (*Plugin) ExampleFileName() string {
	return exampleFileName
}

// DescriptionForParams to print help info
func (*Plugin) DescriptionForParams() map[string]string {
	return map[string]string{
		"no-dump":  "no dump in raw.js to use sblock-web-driver",
		"raw-name": "define the name of constant in raw.js to use sblock-web-driver",
	}
}

// New is the entry of the plugin
func New(c *sbi.Config) sbi.Plugin {
	return &Plugin{
		c,
	}
}
