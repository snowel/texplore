package texploreVoodoo

/*--- 
		  This is a package for testing voodoo layouts.
		  -repeating single finger use 
		  -repeating same hand use
		  -etc
---*/


import (
		  "strings"
		  tfmt "texplore/dataformat"
)

var (

		  Keymap1 = map[string][]string{
					 "Right-Index": {"n", "N", "h", "H", "p", "P", "m", "M"},
					 "Right-Middle": {"e", "E", " ", ";", ":", ","},
					 "Right-Ring": {"i", "I", "l", "L", ".", "'", "\""},
					 "Right-Pinky": {"o", "O", "w", "W", "x", "X"},
					 "Left-Index": {"t", "T", "d", "D", "b", "B", "g", "G"},
					 "Left-Middle": {"r", "R", "y", "Y", "v", "V", "c", "C"},
					 "Left-Ring": {"s", "S", "f", "F", "k", "K", "j", "J"},
					 "Left-Pinky": {"a", "A", "u", "U", "q", "Q", "z", "Z"},

		  }
		  Keymap2 = map[string][]string{
					 "Right-Index": {"n", "N", "h", "H", "p", "P", "m", "M"},
					 "Right-Middle": {"e", "E", "l", "L", ";", ":", ","},
					 "Right-Ring": {"i", "I", " ", ".", "'", "\""},
					 "Right-Pinky": {"o", "O", "w", "W", "x", "X"},
					 "Left-Index": {"t", "T", "d", "D", "b", "B", "g", "G"},
					 "Left-Middle": {"r", "R", "y", "Y", "v", "V", "c", "C"},
					 "Left-Ring": {"s", "S", "f", "F", "k", "K", "j", "J"},
					 "Left-Pinky": {"a", "A", "u", "U", "q", "Q", "z", "Z"},

		  }
		  Keymap3 = map[string][]string{
					 "Right-Index": {"n", "N", "h", "H", "p", "P", "m", "M"},
					 "Right-Middle": {"e", "E", "l", "L", ";", ":", ","},
					 "Right-Ring": {"i", "I", " ", ".", "'", "\""},
					 "Right-Pinky": {"o", "O", "w", "W", "x", "X"},
					 "Left-Index": {"t", "T", "d", "D", "b", "B", "g", "G"},
					 "Left-Middle": {"r", "R", "y", "Y", "v", "V", "c", "C"},
					 "Left-Ring": {"s", "S", "f", "F", "k", "K", "j", "J"},
					 "Left-Pinky": {"a", "A", "u", "U", "q", "Q", "z", "Z"},

		  }
		  Keymap4 = map[string][]string{
					 "Right-Index": {"n", "N", "h", "H", "p", "P", "m", "M"},
					 "Right-Middle": {"e", "E", "y", ";", ":", ","},
					 "Right-Ring": {"i", "I", "l", "L", ".", "'", "\""},
					 "Right-Pinky": {"o", "O", "w", "W", "x", "X"},
					 "Left-Index": {"t", "T", "d", "D", "b", "B", "g", "G"},
					 "Left-Middle": {"r", "R", "y", "Y", "v", "V", "c", "C"},
					 "Left-Ring": {"s", "S", "f", "F", "k", "K", "j", "J"},
					 "Left-Pinky": {"a", "A", "u", "U", "q", "Q", "z", "Z"},
		  }
)

// Check if any of the strings(keys) the finger is reponsible for appear in the block
func fingerUse(block string, fingerMap []string) int {
		  keys := len(fingerMap)
		  counter := 0

		  for i := 0; i < keys; i++ {
					 if strings.Count(block, fingerMap[i]) > 0 {
								counter++
					 }
		  }

		  return counter
}


// Takes a slice pair and retruns a map of how frequently one finger is used is a given block where the blocks are bigrams
// if it's a bigram it same finger rpetition
// chracter slice is finger use frentchecy
func BigramEval(pair tfmt.Slicepair, keymap map[string][]string) map[string]int {
		  blocks := pair.Blocks
		  freq := pair.Occurences

		  eval := make(map[string]int)
		  length := len(blocks)

		  for i := 0; i < length; i++ {
					 for finger, mappedKeys := range keymap {
								if fingerUse(blocks[i], mappedKeys) == 2 {
										  eval[finger] += freq[i]
								}
					 } 
		  }

		  return eval
}
