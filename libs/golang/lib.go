package sb

import (
	"../../sbi"
	"github.com/spf13/afero"
)

// New is the entry of library
func New(raw string, algorithm string) (afero.Fs, error) {
	fn := decodeAlgorithm(algorithm)
	if fn != nil {
		return fn(raw)
	}
	return nil, sbi.ErrUnknowAlgorithm
}

func decodeAlgorithm(algorithm string) func(string) (afero.Fs, error) {
	if algorithm == "default" {
		algorithm = "zip"
	}
	switch algorithm {
	case "origin", "zip":
		return unZip
	}
	return nil
}
