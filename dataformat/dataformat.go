package dataformat

import (
		  "fmt"
)

type Slicepair struct {
		  Blocks []string
		  Occurences []int
}

func SlicepairOccSum(pair *Slicepair) int {
		  length := len(pair.Occurences)
		  counter := 0

		  for i := 0; i < length; i++ {
					 counter += pair.Occurences[i]
		  }

		  return counter
}


func KeymapAppend(finger string, newkey string, keymap map[string][]string) {
		 
          var empty []string
			 _, ok := keymap[finger]
			 if ok == true {
						keymap[finger] = append(keymap[finger], newkey)
			 } else {
						keymap[finger] = empty
						keymap[finger] = append(keymap[finger], newkey)
			 }
}


/*--- Print formating ---*/

// Sorts a slicepair.
func MirrorSort(unsortedPair Slicepair) Slicepair {

		  nums := unsortedPair.Occurences
		  elems := unsortedPair.Blocks

		  length := len(nums)
		  sortedElems := make([]string, length)
		  sortedNums := make([]int, length)

		  var index int
		  var biggest int
		  for i := 0; i < length; i++ {
					 index = i
					 biggest = nums[index]
					 for j := 0; j < length; j++ {
								if nums[j] > biggest {
										  index = j
										  biggest = nums[j]
								}
					 }

					 sortedNums[i] = nums[index]
					 sortedElems[i] = elems[index]

					 nums[index] = -1
		  }
		  
		  pair := Slicepair{Blocks: sortedElems, Occurences: sortedNums}

		  return pair
}

// Converts a map into a slicepair and sorts it.
func SortMap(collect map[string]int) Slicepair {
		  // Split the map into matching slices.
		  length := len(collect)
		  textKey := make([]string, length)
		  occurs := make([]int, length)
		  i:= 0
		  for key, val := range collect {
					 textKey[i] = key
					 occurs[i] = val
					 i++

		  }

		  // Mirror sort the slices.
		  sortedPair := Slicepair{Blocks: textKey, Occurences: occurs}
		  sortedPair = MirrorSort(sortedPair)

		  return sortedPair
}

func PrintSlicepair(pair Slicepair) {
		  text := pair.Blocks
		  occ := pair.Occurences

		  length := len(occ)
		  for i := 0; i < length; i++ {
					 fmt.Printf(":::  %s -> %d\n", text[i], occ[i])
		  }
}
