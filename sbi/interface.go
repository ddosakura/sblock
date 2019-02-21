package sbi // import "github.com/ddosakura/sblock/sbi"

import (
	"errors"
	"os"
)

// Config define the raw data of plugin
type Config struct {
	PKG     string
	Comment string

	Algorithm string

	Params map[string]string
}

// Plugin is the interface of plugin
type Plugin interface {
	SourceFileName() string
	ExampleFileName() string

	Source(*os.File, func())
	Encode(path, relPath string, fi os.FileInfo, err error) error

	Example(*os.File)
}

var (
	// ErrUnknowAlgorithm throw when the algorithm is unknow
	ErrUnknowAlgorithm = errors.New("Unknow Algorithm")
)
