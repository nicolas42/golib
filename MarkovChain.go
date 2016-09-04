package golib

import (
	"strings"
	"math/rand"
	"time"
)

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