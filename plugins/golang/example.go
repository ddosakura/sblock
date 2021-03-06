package sbgo

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ddosakura/gklang"
)

// Example help user to know how to use the library
func (p *Plugin) Example(f *os.File) {
	path2sblock := p.c.Params["path2sblock"]
	if path2sblock == "" {
		path2sblock = "../" + p.c.PKG
	}

	raw := fmt.Sprintf(`// Example generated by sblock. CAN BE DELETED.

package main

import (
	"log"
	"net/http"

	"github.com/ddosakura/sblock/libs/golang"
	sblock "%s"
	"github.com/spf13/afero"
)

func main() {
	fs, err := sb.New(sblock.Raw, "default")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(afero.NewHttpFs(fs))))
	http.ListenAndServe(":8080", nil)
}
`, path2sblock)

	var qb bytes.Buffer
	fmt.Fprintf(&qb, raw)

	if err := ioutil.WriteFile(f.Name(), qb.Bytes(), 0644); err != nil {
		gklang.Er(err)
	}
}
