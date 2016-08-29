package nick

import (
	"testing"
	"fmt"
	"strings"
)

func TestSplitAny(t *testing.T) {

	fmt.Println(`splitAny("234 * 3 *  234 (4234 )/ 242/*^", "*/+-^()")`)

	split := SplitAny("234 * 3 *  234 (4234 )/ 242/*^", "*/+-^()")	

	/* // print
	out := []string{}
	for _,v := range split {
		out = append(out, fmt.Sprintf("[%v] ", v))
	}
	*/
	fmt.Println(strings.Join(split,", "))
	println()
}

func TestMarkovChain(t *testing.T) {

	// Read Project Gutenberg's The Adventures of Sherlock Holmes, by Arthur Conan Doyle
	bytes := ReadThru("http://www.gutenberg.org/cache/epub/1661/pg1661.txt") 
	
	fmt.Println("MarkovChain(string(bytes), 50, 2)")
	fmt.Println( MarkovChain(string(bytes), 50, 2) )
	println()
		
}
