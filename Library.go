package golib

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
	"path/filepath"
	"os"
)

func ReadUrl(url string) []byte {
	// To Do: error handling
	fmt.Printf("Reading %v...\n", url)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func AutoRename(file string ) string {

	// renames filename string if there is a collision
	nf := file  // newfile
	n := 1
	for	IsExist(nf) { // file exists
		n += 1
		ext := filepath.Ext(nf)
		name := strings.TrimSuffix(nf, ext)
		nf = fmt.Sprintf("%v %v%v", name, n, ext)
	}
	return nf
}

func IsExist(file string) bool {
	_, err := os.Stat(file) 
	return err == nil
}

func PrintCurly(v []string) {
	for _,v := range v {
		fmt.Printf("{%v} ", v)
	}

}

func SplitAny(str string, chars string) []string {

	// Splits string at any of the given chars but also appends the chars to the list
	// in the output string array
	// Example:
	// splitAny("234 * 3 *  234 (4234 )/ 242/*^", "*/+-^()")
	// == 234, *, 3, *, 234, (, 4234, ), /, 242, /, *, ^

	out := []string{}
	p1, p2 := 0, 0 // positions
	var num, op, lastNum string

	for {

		// find the next operation character
		p2 = strings.IndexAny(str[p1:], chars)
		if p2 == -1 {
			break
		}
		p2 = p1 + p2

		num = strings.TrimSpace(string(str[p1:p2]))
		op = strings.TrimSpace(string(str[p2]))

		// append if not empty
		if num != "" {
			out = append(out, num)
		}
		if op != "" {
			out = append(out, op)
		}

		p1 = p2 + 1
	}

	lastNum = strings.TrimSpace(string(str[p1:]))
	if lastNum != "" {
		out = append(out, lastNum)
	}

	return out
}

