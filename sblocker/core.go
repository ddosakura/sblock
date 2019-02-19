package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"../libs/common"
	"github.com/ddosakura/gklang"
)

// test src/target/example with force
func cleanTarget() {
	clean(target)
	clean(example)
}

func clean(dest string) {
	var err error
	if _, err = os.Stat(dest); !os.IsNotExist(err) {
		if force {
			if err = os.RemoveAll(dest); err != nil {
				gklang.Log(gklang.LCrash, err, fmt.Errorf("file %q could not be deleted", dest))
			}
		} else {
			gklang.Er(fmt.Errorf("file %q already exists; use -f to overwrite", dest))
		}
	}
}

var (
	generater sbi.Plugin
)

// generate code
func generate() {
	gklang.Log(gklang.LInfo, "generating")

	rename(generateSource().Name(), target, generater.SourceFileName())

	rename(generateExample().Name(), example, generater.ExampleFileName())
}

func generateSource() (f *os.File) {
	// ${pkg}${rand}
	f, err := ioutil.TempFile("", pkg)
	if err != nil {
		gklang.Er(err)
	}

	generater.Source(f, func() {
		if err = filepath.Walk(src, func(path string, fi os.FileInfo, err error) error {
			relPath, err := filepath.Rel(src, path)
			if err != nil {
				return err
			}
			return generater.Encode(path, relPath, fi, err)
		}); err != nil {
			gklang.Er(err)
		}
	})

	return
}

func generateExample() (f *os.File) {
	// ${pkg}.example${rand}
	f, err := ioutil.TempFile("", pkg+".example")
	if err != nil {
		gklang.Er(err)
	}

	generater.Example(f)

	return
}

func rename(src, dest, filename string) {
	err := os.MkdirAll(dest, 0755)
	if err != nil {
		gklang.Er(err)
	}

	targetFile := path.Join(dest, filename)

	// Try to rename generated source.
	if err := os.Rename(src, targetFile); err == nil {
		return
	}
	// If the rename failed (might do so due to temporary file residing on a
	// different device), try to copy byte by byte.
	rc, err := os.Open(src)
	if err != nil {
		gklang.Er(err)
	}
	defer func() {
		rc.Close()
		os.Remove(src) // ignore the error, source is in tmp.
	}()

	wc, err := os.Create(targetFile)
	if err != nil {
		gklang.Er(err)
	}
	defer wc.Close()

	if _, err = io.Copy(wc, rc); err != nil {
		// Delete remains of failed copy attempt.
		os.RemoveAll(dest)
	}
}
