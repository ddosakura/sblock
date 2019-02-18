package sbgo

import (
	"../common"
	"github.com/spf13/afero"
)

const (
	sourceFileName  = "raw.go"
	exampleFileName = "example.go"
)

// SourceFileName return the FileName of the source
func (*Plugin) SourceFileName() string {
	return sourceFileName
}

//ExampleFileName return the FileName of the example
func (*Plugin) ExampleFileName() string {
	return exampleFileName
}

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

// NewFs is the entry of library
func NewFs(raw string) (afero.Fs, error) {
	return New(&sbi.Config{}).Decode(raw)
}
