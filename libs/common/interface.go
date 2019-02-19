package sbi // import "github.com/ddosakura/sblock/libs/common"

import (
	"errors"
	"os"

	"github.com/spf13/afero"
)

// Config define the raw data of plugin
type Config struct {
	Dev bool

	PKG     string
	Comment string

	Algorithm string
}

// Plugin is the interface of plugin
type Plugin interface {
	SourceFileName() string
	ExampleFileName() string

	Source(*os.File, func())
	Encode(path, relPath string, fi os.FileInfo, err error) error

	Example(*os.File)
	Decode(string) (afero.Fs, error)
}

var (
	// ErrUnknowAlgorithm throw when the algorithm is unknow
	ErrUnknowAlgorithm = errors.New("Unknow Algorithm")
)
