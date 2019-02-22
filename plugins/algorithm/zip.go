package algorithm

import (
	"archive/zip"
	"os"
	"path/filepath"
)

// Zip used by zip/origin
func Zip(deflate bool, relPath string, fi os.FileInfo, b []byte, w *zip.Writer) error {
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
	if deflate {
		fHeader.Method = zip.Deflate
	}
	f, err := w.CreateHeader(fHeader)
	if err != nil {
		return err
	}
	_, err = f.Write(b)
	return err
}
