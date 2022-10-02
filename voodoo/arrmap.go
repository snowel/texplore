package texploreVoodoo

import (
		  "fmt"
		  "strings"
		  tfmt "texplore/dataformat"
)

// Optimization :: Array keymaps - Is it worth it?
// A keymap will be an array of string slices: [n][]string, where n is thenumebr of fingers
// In the array positions indicate the key ID.
// In practice, they match the postions the 

// Evalcollect represents a slice of maps and some per-figner evaluation
type Evalcollect struct {
		  maps [][8][]string
		  evals [][8]int
}

type MapMeta struct {
		  Keymap [8][]string
		  Bireps [8]int
		  FingerUse [8]int
		  Handuse [8]int
}

// Combining Arr-maps.
// A single map represetns, for each finger, a slice of strings, which themselve are the keystrokes.
// The order of the strings in the slice is representative of which row the key is on.
// In this combinaiton, the rows of map2 are appended after the rows of map.
func MergeArrMaps(Map [8][]string, Map2 [8][]string) [8][]string {
		  var catmap [8][]string

		  for i := 0; i < 8; i++ {
					 catmap[i] = append(catmap[i], Map[i]...)
					 catmap[i] = append(catmap[i], Map2[i]...)
		  }
		  
		  return catmap
}


func ArrFingerUse(block string, fingerMap []string) int {
		  counter := 0
		  for _, v := range fingerMap {
					 counter += strings.Count(block, v)
		  }
		  return counter
}
// Evaluate how much each finger is used in total.
func ArrFingerEval(text string, keymap [8][]string) [8]int {
		  var eval [8]int
		  for j := 0; j < 8; j++ {
					 eval[j] += ArrFingerUse(text, keymap[j])
		  } 

		  return eval
}

// Evaluate how much each hand is used.
func ArrHandEval(fingerEval [8]int) [2]int {
		  var hands [2]int
		  for i, v := range fingerEval {
					 if i <=3 {
								hands[0] += v
					 } else {
								hands[1] += v
					 }
		  }
		  return hands
}

// Evaluate how much each finger is used twice in a row.
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

// Conver a slice of keystrokes to an array of slices, organizing into finger groups.
func SmapToArrmap(smap []string) [8][]string {
		  var arrmap [8][]string
		  for i, v := range smap {
					 arrmap[i % 8] = append(arrmap[i % 8], v)
		  }
		  return arrmap
}


func ArrPrintEval(keymap [8][]string, eval [8]int) {

		  for i := 0; i < 8; i++ {
					 fmt.Println("Finger :: ", i, " is repeaded :: ", eval[i])
					 fmt.Println(keymap[i])
		  }
}

// Takes a list of heap randomized slicemaps and gives the values of how often a finger is repeated
func EvalArrMaps(smaps [][]string, basemap [8][]string, ref tfmt.Slicepair) Evalcollect {
		  length := len(smaps)
		  evaluations := make([][8]int, length)
		  maps := make([][8][]string, length)
		  var (
					 arrmap [8][]string
					 keymap [8][]string
		  )
		  for i, v := range smaps {
					 arrmap = SmapToArrmap(v)
					 keymap = MergeArrMaps(arrmap, basemap)
					 eval := ArrBigramEval(ref, keymap)
					 evaluations[i] = eval
					 maps[i] = keymap
					 //fmt.Println(i, " of ", length)
					 //fmt.Println(keymap)
					 //fmt.Println(eval)
					 //fmt.Println(ArrmapEvalSum(eval))
		  }
		  newEval := Evalcollect{maps: maps, evals: evaluations}
		  return newEval
}

//TODO Poorly named...
// For a given eval, what is the maximum quantity?
func ArrMax(array [8]int) int {
		  max := array[0]
		  for _, v := range array {
					 if v > max {
								max = v
					 }
		  }

		  return max
}

// For a given eval, what is the minimum quantity?
func ArrMin(array [8]int) int {
		  min := array[0]
		  for _, v := range array {
					 if v < min {
								min = v
					 }
		  }

		  return min
}

// Gives the sum of all quantities in an eval.
func ArrmapEvalSum (array [8]int) int {
		  sum := 0
		  for _, v := range array {
					 sum += v
		  }
		  return sum
}

// Finds the map with the smallest max repetition of any one finger
func SmallestRep(collect Evalcollect) ([8][]string, [8]int, int) {
		  alleval := collect.evals
		  minRep := ArrMax(collect.evals[0])
		  minIndex := 0
		  var singleMin int

		  for i, v := range alleval {
					 singleMin = ArrMax(v)
					 if singleMin < minRep {
								minIndex = i
								minRep = singleMin
					 }
		  }

		  return collect.maps[minIndex], alleval[minIndex], minIndex
}


//
func NSmallestRep(collect Evalcollect, n int) Evalcollect {
		  mutCol := collect
		  var topCol Evalcollect

		  for i := 0; i < n; i++ {
					 topmap, topeval, index := SmallestRep(mutCol)
					 topCol.maps = append(topCol.maps, topmap)
					 topCol.evals = append(topCol.evals, topeval)
					 
					 mutCol.maps = append(mutCol.maps[:index], mutCol.maps[index+1:]...)
					 mutCol.evals = append(mutCol.evals[:index], mutCol.evals[index+1:]...)

					 //finalElem := len(mutCol.maps) - 1
					 //mutCol.maps[index] = mutCol.maps[finalElem]
					 //mutCol.evals[index] = mutCol.evals[finalElem]
					 //mutCol.maps = mutCol.maps[:finalElem]
					 //mutCol.evals = mutCol.evals[:finalElem]
					 
					 fmt.Println()
					 fmt.Println("This is the ", i, " lowest biggest reps.")
					 fmt.Println(topmap)
					 fmt.Println(topeval)
					 fmt.Println(ArrmapEvalSum(topeval))
		  }
		  return topCol
}

// For a collectoin of evaluations,find the one with the minimum total repetition of all fingers 
func SmallestTotalReps(collect Evalcollect) ([8][]string, [8]int, int, int) {
		  alleval := collect.evals
		  minTotal := ArrmapEvalSum(alleval[0])
		  minIndex := 0
		  var singleMin int

		  for i, v := range alleval {
					 singleMin = ArrmapEvalSum(v)
					 if singleMin < minTotal {
								minIndex = i
								minTotal = singleMin
					 }
		  }

		  return collect.maps[minIndex], alleval[minIndex], minIndex, minTotal
}

func NSmallestTotalRep(collect Evalcollect, n int) Evalcollect {
		  mutCol := collect
		  var topCol Evalcollect

		  for i := 0; i < n; i++ {
					 topmap, topeval, index, total := SmallestTotalReps(mutCol)
					 topCol.maps = append(topCol.maps, topmap)
					 topCol.evals = append(topCol.evals, topeval)
					 
					 mutCol.maps = append(mutCol.maps[:index], mutCol.maps[index+1:]...)
					 mutCol.evals = append(mutCol.evals[:index], mutCol.evals[index+1:]...)

					 //finalElem := len(mutCol.maps) - 1
					 //mutCol.maps[index] = mutCol.maps[finalElem]
					 //mutCol.evals[index] = mutCol.evals[finalElem]
					 //mutCol.maps = mutCol.maps[:finalElem]
					 //mutCol.evals = mutCol.evals[:finalElem]
					
					 fmt.Println()
					 fmt.Println("This is the ", i, " lowest total reps.")
					 fmt.Println(topmap)
					 fmt.Println(topeval)
					 fmt.Println(total)
		  }
		  return topCol
}
/* -- Layer Swtich Optimization -- */

// Semantics:
// In the case of a layer optimizing for layer switching, the arrays of a [num][]string maps don't represent the fingers, but the layers.
// i.e. for a  16 key voodoo, this would be a map of the form [2][]string but for an 8 key it would be [4][]string 

// TODO CUrrently, the layercount of the map is hardcoded, as for finger speerated maps
// For a given layer seperated map, count the number of times a switch happens (including the switch back to the base layer)
func CountLayerSwitch(keymap [2][]string, refText string) int{
		  counter := 0
		  layers := len(keymap)
		  currentLayer := 0;
		  for _, v := range refText {
					 if !sliceContains(keymap[currentLayer], v) {

								// find the layer with the current character
								loops := 0 // if the layer map is incomplete the loops will be infinite
								for !sliceContains(keymap[currentLayer], v) {
										  // Cycle to the next layer
										  currentLayer++
										  // Wrap around
										  if currentLayer >= layers {currentLayer %= layers}

										  if loops > layers {return -1}// the layer map is missing the character
										  loops++
										  // This could also be planned to skip that character, but this might be more responsible.
								}

								counter++
					 }
		  }
		  return counter

}
