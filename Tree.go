// changed tabs to 4 spaces
// How to check if a file exists in Go?  http://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go

package nick

import (
  "io/ioutil"
  "os"
)

func main() {

	if len(os.Args)>=2 {
		if isDir(os.Args[1]){
			Rtree(os.Args[1],"")
		}
	} else {
		wd,_ := os.Getwd()
		Rtree(wd,"")
	}
}

func isDir(arg string) bool {
	f, err := os.Stat(arg); 
	if err != nil { 
		return false
	} 
	return f.IsDir()
}

func Rtree(dir, tabs string) {
	
	files, _ := ioutil.ReadDir(dir) // returns []os.FileInfo
	for _, f := range files {
		if f.IsDir() {
			println(tabs+f.Name()+"/")
			Rtree( dir+"/"+f.Name(), tabs+"    " )
		} else {
			println(tabs+f.Name())
		}
	}
}


/* 
// ioutil
func ReadDir(dirname string) ([]os.FileInfo, error)


// os
type FileInfo interface {
        Name() string       // base name of the file
        Size() int64        // length in bytes for regular files; system-dependent for others
        Mode() FileMode     // file mode bits
        ModTime() time.Time // modification time
        IsDir() bool        // abbreviation for Mode().IsDir()
        Sys() interface{}   // underlying data source (can return nil)
}

*/