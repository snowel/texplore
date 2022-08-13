package texploreVoodoo

/*--- 
		  This is a package for testing voodoo layouts.
		  -repeating single finger use 
		  -repeating same hand use
		  -etc
---*/


import (
		  "C"
		  "fmt"
		  "strings"
		  "reflect"
		  tfmt "texplore/dataformat"
)

var (

		  EmptyMap = make(map[string][]string)
		  SingleLayer2 = []string{"h", "n", "e", "m", "i", "l", "o", "w", "t", "d", "r", "c", "s", "f", "a", "u"}
		  SingleLayer1 = []string{"h", "m", "l", "w", "d", "c", "f", "u"}


		  ArrayMap1 = [8][]string{
					 []string{"h", "n"},
					 []string{"e", "m"},
					 []string{"i", "l"},
					 []string{"o", "w"},
					 []string{"t", "d"},
					 []string{"r", "c"},
					 []string{"s", "f"},
					 []string{"a", "u"},
		  }

		  SingleLayerMap = map[string][]string{
					 "ri": {"h"},
					 "rm": {"m"},
					 "rr": {"l"},
					 "rp": {"w"},
					 "li": {"d"},
					 "lm": {"c"},
					 "lr": {"f"},
					 "lp": {"u"},
		  }
		  SingleLayerMap2 = map[string][]string{
					 "ri": {"h", "n"},
					 "rm": {"e", "m"},
					 "rr": {"i", "l"},
					 "rp": {"o", "w"},
					 "li": {"t", "d"},
					 "lm": {"r", "c"},
					 "lr": {"s", "f"},
					 "lp": {"a", "u"},
		  }

		  Keymap1 = map[string][]string{
					 "Right-Index": {"n", "h", "p", "k"},
					 "Right-Middle": {"e", "m", ":", ","},
					 "Right-Ring": {"i", "l", ".", "'", "\""},
					 "Right-Pinky": {"o", "w", "x" },
					 "Left-Index": {"t", "d", "b", "g" },
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


// --- Procedural

// Takes a slicemap collection and evalueates each into a freq map

func EvalMaps(slicemaps [][]string, basemap map[string][]string, ref tfmt.Slicepair) ([]map[string][]string, []map[string]int) {
		  length := len(slicemaps)
		  evaluations := make([]map[string]int, length)
		  maps := make([]map[string][]string, length)
		  for i := 0; i < length; i++ {
					 keymap := SliceToKeymap(slicemaps[i], basemap)
					 eval := BigramEval(ref, keymap)
					 evaluations[i] = eval
					 maps[i] = keymap
					 fmt.Println(i, " of ", length)
		  }
		  return maps, evaluations
}

//Heaps takes in a slice of chars and produces a slice of all possible orders
func HeapSliceMap(chars []string) [][]string {
		  var collect [][]string
		  Heaps (len(chars), chars, &collect)
		  return collect
}

// Heaps algo, recursive
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

// If there is no forced map, eval mats is an empty map = make(map[string][]string)
// Takes a slice of strings and retruns a keymap, appending each block of n characters to a certain finger
func SliceToKeymap(slicemap []string, keymap map[string][]string) map[string][]string {
		  length :=  len (slicemap)
		  block := (length / 8) - length % 8// Need to hadle overflow
		  
		  for i := 0; i < length; i++ {
					 switch {
					 case i < block: tfmt.KeymapAppend("ri", slicemap[i], keymap)
					 case i < 2*block: tfmt.KeymapAppend("rm", slicemap[i], keymap)
					 case i < 3*block: tfmt.KeymapAppend("rr", slicemap[i], keymap)
					 case i < 4*block: tfmt.KeymapAppend("rp", slicemap[i], keymap)
					 case i < 5*block: tfmt.KeymapAppend("li", slicemap[i], keymap)
					 case i < 6*block: tfmt.KeymapAppend("lm", slicemap[i], keymap)
					 case i < 7*block: tfmt.KeymapAppend("lr", slicemap[i], keymap)
					 case i < 8*block: tfmt.KeymapAppend("lp", slicemap[i], keymap)
					 }
		  }
		  return keymap
}

// Optimization :: Slice keymaps - Is it worth it?
// a keyman will be a slice of length n*8 where keys 0 through n-1 will have all chars the left pinky uses, n-1 through 2n-1 all keys for left ring, etc

func arrFingerUse(block string, fingerMap []string) int {
		  keys := len(fingerMap)
		  counter := 0

		  for i := 0; i < keys; i++ {
					 if strings.Count(block, fingerMap[i]) > 0 {
								counter++
					 }
		  }

		  return counter
}

func indexToFinger(i int, mult int) int {
		  position := i / 8
		  overflow := i % 8
		  return i - position * mult - overflow
}
func fingerToIndex(i int, mult int) (int,int) {
		  return i*mult, i*mult + (mult-1)
}


func SMapBigramEval(pair tfmt.Slicepair, keymap []string, mult int) [8]int {
		  mapLen := len(keymap)
		  if mapLen % 8 != 0 {
					 fmt.Println("SMapBigramEval :: Keymaps isn't multiple of 8. Use \"{[empty]}\" for none keys on ifinger if not all are used")
		  }
		  blocks := pair.Blocks
		  freq := pair.Occurences

		  var eval [8]int
		  for i, v := range blocks {
					 
					 for j := 0; j < mapLen; j += mult {
								finger := indexToFinger(j, mult)
								start, finish := fingerToIndex(finger, mult) 
								var fingerMap []string
								if j == mapLen -1 {
										  fingerMap = keymap[start:]
								} else {
										  fingerMap = keymap[start:finish]
								}
								if arrFingerUse(v, fingerMap) > 1 {
										  eval[finger] += freq[i]
								}
					 } 
		  }

		  return eval
}

func MergeSMaps(sMap []string, sMap2 []string) []string {
		  mult := len(sMap) / 8
		  mult2 := len(sMap2) / 8
		  var catmap []string

		  for i := 0; i < 8; i++ {
					 start, finish := fingerToIndex(i, mult) 
					 catmap = append(catmap, sMap[start:finish + 1]...)
					 start, finish = fingerToIndex(i, mult2) 
					 catmap = append(catmap, sMap2[start:finish + 1]...)
		  }
		  
		  return catmap
}


func PrintEval(keymap []string, mult int, eval [8]int) {

		  for i:= 0; i < 8; i++ {
					 start, finish := fingerToIndex(i, mult) 
					 fmt.Println("Finger :: ", i, " is repeaded :: ", eval[i])
					 fmt.Println(keymap[start:finish + 1])
		  }
}

// Optimization :: Array keymaps - Is it worth it?
//A keymap will be an array of string slices, positions mathcing the left to right position of the fingers on the keys

func ArrBigramEval(pair tfmt.Slicepair, keymap [8][]string) [8]int {
		  mapLen := len(keymap)
		  if mapLen % 8 != 0 {
					 fmt.Println("SMapBigramEval :: Keymaps isn't multiple of 8. Use \"{[empty]}\" for none keys on ifinger if not all are used")
		  }
		  blocks := pair.Blocks
		  freq := pair.Occurences

		  var eval [8]int
		  for i, v := range blocks {
					 
					 for j := 0; j < 8; j++ {
								if arrFingerUse(v, keymap[j]) > 1 {
										  eval[j] += freq[i]
								}
					 } 
		  }

		  return eval
}

func MergeArrMaps(Map [8][]string, Map2 [8][]string) [8][]string {
		  var catmap [8][]string

		  for i := 0; i < 8; i++ {
					 catmap[i] = append(catmap[i], Map[i]...)
					 catmap[i] = append(catmap[i], Map2[i]...)
		  }
		  
		  return catmap
}


func ArrPrintEval(keymap [8][]string, eval [8]int) {

		  for i := 0; i < 8; i++ {
					 fmt.Println("Finger :: ", i, " is repeaded :: ", eval[i])
					 fmt.Println(keymap[i])
		  }
}


