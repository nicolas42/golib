package nick

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
	"math/rand"
	"time"
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

func MarkovChain( arg string, chainLen int, matchLen int) string {

	words := strings.Fields(arg)
		
	// create chain and fill it with two words from a random location
	rand.Seed(time.Now().Unix())
	chain := []string{}
	pos := rand.Intn(len(words)-matchLen)
	
	for w:=0; w<chainLen; w+=1 {
		chain = append(chain, words[pos+w])
	}
			
	for i:=1; i<=chainLen; i+=1 {
	
		query := chain [ len(chain)-matchLen: ] // last two words in chain
		matches := []string{}	
		
		for i,_ := range words {
			isMatch := true
			for j:=0; j<matchLen; j+=1 {
				if words[i+j] != query[j] {
					isMatch = false
					break
				}
			}
			if isMatch == true {
				matches = append(matches, words[i:i+matchLen]...)
			}
		}

		if len(matches) == 0 { 
			break 
		} else if len(matches) == 1 { 
			chain = append(chain, matches[0])
		} else if len(matches) > 1 {
			chain = append( chain, matches[rand.Intn(len(matches)-1)] )
		}
				
	}
	return strings.Join(chain, " ")
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

