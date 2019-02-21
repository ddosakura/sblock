package sbgo

import (
	"archive/zip"
	"io"
	"os"

	"github.com/ddosakura/gklang"
)

var (
	w *zip.Writer
)

func initAlgorithm(p *Plugin, writer io.Writer) {
	if p.c.Algorithm == "default" {
		p.c.Algorithm = "zip"
	}
	switch p.c.Algorithm {
	case "origin", "zip":
		w = zip.NewWriter(writer)
	default:
		gklang.Er(ErrUnknowAlgorithm)
	}
}

func finishAlgorithm(p *Plugin) {
	switch p.c.Algorithm {
	case "origin", "zip":
		if err := w.Close(); err != nil {
			gklang.Er(err)
		}
	default:
		gklang.Er(ErrUnknowAlgorithm)
	}

}

func encodeAlgorithm(p *Plugin, relPath string, fi os.FileInfo, b []byte) error {
	switch p.c.Algorithm {
	case "origin", "zip":
		return doZip(p, relPath, fi, b)
	}
	return ErrUnknowAlgorithm
}