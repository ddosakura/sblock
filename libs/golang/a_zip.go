package sb

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/spf13/afero"
)

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
