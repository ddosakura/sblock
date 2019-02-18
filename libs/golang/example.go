package sbgo

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ddosakura/gklang"
	"github.com/spf13/afero"
)

// Example help user to know how to use the library
func (p *Plugin) Example(f *os.File) {
	raw := `package main

import (
	"log"
	"net/http"
	"os"

	"../../libs/golang"
	"../sblock"
	"github.com/ddosakura/gklang"
	"github.com/spf13/afero"
)

func main() {
	l := log.New(os.Stdout, "[Example]: ", log.LstdFlags)
	gklang.LoadLogger(l, gklang.LvDebug)
	fs, err := sbgo.NewFs(sblock.Raw)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(afero.NewHttpFs(fs))))
	http.ListenAndServe(":8080", nil)
}
`
	var qb bytes.Buffer
	fmt.Fprintf(&qb, raw)

	if err := ioutil.WriteFile(f.Name(), qb.Bytes(), 0644); err != nil {
		gklang.Er(err)
	}
}

// Decode is the func for your project
func (p *Plugin) Decode(raw string) (afero.Fs, error) {
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
