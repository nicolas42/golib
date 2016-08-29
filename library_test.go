package golib

import (
	"testing"
	"fmt"
	"strings"
	"os"
)

func TestTree(t *testing.T){
	// Should go first as can print a lot of text
	wd,_ := os.Getwd()

	fmt.Println("Current dir")
	Tree(wd,"")
	
//	fmt.Println("Up one dir")
//	Tree("../", "")

//	if len(os.Args)>=2 {
//		if IsDir(os.Args[1]){
//			Tree(os.Args[1],"")
//		}
	println()
}


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

