package texploreVoodoo

/*--- This is a package for testing voodoo layouts.
		  repeating single finger use 
		  repeating same hand use
		  etc ---*/


/* Note that eventually, hand/finger count agnostic and keybaord agnostic packages are planned*/

import (
		  "strings"
		  "texplore/ascii"
		 // "reflect"
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
)

// procedurally genreated keymaps
/*
var L1charPool = []string{"u", "f", "c", "d", "h", " ", "l", "w" }

// generates all the maps of a voodoo map, one row is fixed, the other is permuted
func fixedRowPermRow(fixedMap map[string][]string, permMap map[string][]string) *[]map[string][]string {
		  var permutations [][]string
		  var allMyMaps []map[string][]string

		  Heaps(8, permMap, &permutations) // poppulated the permutaitons
//TODO aliase row_length = 8
		  nPerms := len(permutations) // determine how many varaints we have

		  for i := 0; i < nPerms; i++ {

					 thisMap := fixedMap

					 thisMap["Right-Index"] = append(thisMap["Right-Index"], permutations[i])
					 thisMap["Right-Middle"] = append(thisMap["Right-Middle"], permutations[i])
					 thisMap["Right-Ring"] = append(thisMap["Right-Ring"], permutations[i])
					 thisMap["Right-Pinky"] = append(thisMap["Right-Pinky"], permutations[i])
					 thisMap["Left-Index"] = append(thisMap["Left-Index"], permutations[i])
					 thisMap["Left-Middle"] = append(thisMap["Left-Middle"], permutations[i])
					 thisMap["Left-Ring"] = append(thisMap["Left-Ring"], permutations[i])
					 thisMap["Left-Pinky"] = append(thisMap["Left-Pinky"], permutations[i])

					 *allMyMaps = append(*allMyMaps, thisMap)
		  }

		  return &allMyMaps

}

func Heaps(k int, arr []string, permutations *[][]string) {
		  if k == 1 {
					 fmt.Println(arr)
					 *permutations = append(*permutations, arr)
		  } else {
					 swap := reflect.Swapper(arr)
					 Heaps(k - 1, arr, permutations)
					 
					 for i := 0; i < k - 1; i++ {
								if k % 2 == 0 {
										  swap(i, k - 1)
								} else {
										  swap(0, k - 1)
								}

								Heaps(k - 1, arr, permutations)
					 }
		  }
}
*/

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


// Takes a slice pair and retruns a map of how frequently one finger is used is a given block
// if it's a bigram it same finger rpetition
// chracter slice is finger use frentchecy
func BigramEval(pair texploreASCII.Slicepair, keymap map[string][]string) map[string]int {
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
