package texploreVoodoo

import (
		  "fmt"
		  "strings"
		  tfmt "texplore/dataformat"
)

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
