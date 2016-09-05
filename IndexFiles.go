package golib

import (
	"io/ioutil"
)

func indexFiles(dir string, out *[]string) {
	// mechanics of append
	// https://blog.golang.org/slices
	files, _ := ioutil.ReadDir(dir) // returns []os.FileInfo
	for _, f := range files {
		if f.IsDir() {
			*out = append(*out, f.Name()+"/")
			indexFiles( dir+"/"+f.Name(), out)
		} else {
			*out = append(*out, f.Name())
		}
	}
}

func IndexFiles(dir string) []string {
	out := []string{}
	indexFiles(dir, &out)
	return out
}
