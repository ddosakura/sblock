package sbgo

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
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

func unZip(raw string) (afero.Fs, error) {
	zipReader, err := zip.NewReader(strings.NewReader(raw), int64(len(raw)))
	if err != nil {
		return nil, err
	}
	// files := make(map[string]afero.File, len(zipReader.File))
	fs := afero.NewMemMapFs()
	for _, zipFile := range zipReader.File {
		data, err := unzip(zipFile)
		if err != nil {
			return nil, fmt.Errorf("error unzipping file %q: %s", zipFile.Name, err)
		}

		f, err := fs.Create("/" + zipFile.Name)
		_, err = f.Write(data)
		err = f.Close()

		// files["/"+zipFile.Name] = f
	}
	/*
		for fn := range files {
			dn := path.Dir(fn)
			if _, ok := files[dn]; !ok {
				files[dn] = file{FileInfo: dirInfo{dn}, fs: fs}
			}
		}
	*/
	return fs, nil
}

func unzip(zf *zip.File) ([]byte, error) {
	rc, err := zf.Open()
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return ioutil.ReadAll(rc)
}
