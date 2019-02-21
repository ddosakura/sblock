package sbgo

import (
	"archive/zip"
	"os"
	"path/filepath"
)

func doZip(p *Plugin, relPath string, fi os.FileInfo, b []byte) error {
	fHeader, err := zip.FileInfoHeader(fi)
	if err != nil {
		return err
	}
	/*
		if *flagNoMtime {
			// Always use the same modification time so that
			// the output is deterministic with respect to the file contents.
			// Do NOT use fHeader.Modified as it only works on go >= 1.10
			fHeader.SetModTime(mtimeDate)
		}
	*/
	fHeader.Name = filepath.ToSlash(relPath)
	if p.c.Algorithm == "zip" {
		fHeader.Method = zip.Deflate
	}
	f, err := w.CreateHeader(fHeader)
	if err != nil {
		return err
	}
	_, err = f.Write(b)
	return err
}
