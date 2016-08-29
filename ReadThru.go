package nick

import (
	"strings"
	"io/ioutil"
	"fmt"
	"os"
)

/*
func main(){

	a := ReadThru("http://www.gutenberg.org/cache/epub/236/pg236.txt")
	fmt.Println(string(a[:200]))
	// TestReadThru()
}
*/

func TestReadThru(){

//	a := ioutil.ReadFile("ProjectEuler.go")
//	words := strings.Split(string(a), "\n")

	wd,_ := os.Getwd()
	fmt.Println("Current Directory:", wd)
	
	url := "http://www.gutenberg.org/cache/epub/1661/pg1661.txt"
	filename := url

	filename = strings.Replace(filename, "http://", "", 1)
	filename = strings.Replace(filename, "/", ":", -1)

	fmt.Println("Url:", url)
	fmt.Println("Corresponding file:",filename)
	fmt.Println("Does file exist?:", IsExist(filename))
	
	a := ReadThru("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")
	fmt.Println("Head of file:", string(a[:200]))
}

func ReadThru(url string) []byte {

	bytes := []byte{}
	filename := ConvertToLocal(url)	
	
	if IsExist(filename) {
		bytes,_ = ioutil.ReadFile(filename)
	} else {
		bytes = ReadUrl(url)
		ioutil.WriteFile(filename, bytes, 0644)
	}
	return bytes


	// past code
	// _, filename := filepath.Split(url)	
	// filename := AutoRename(filename) // doesn't work with autorename

}

func ConvertToLocal(url string) string {
	filename := url
	filename = strings.Replace(filename, "http://", "", 1)
	filename = strings.Replace(filename, "/", ":", -1)
//	fmt.Println(url,filename)
	return filename
}
