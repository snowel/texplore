package dataformat

import (
		  "fmt"
)

type Slicepair struct {
		  Blocks []string
		  Occurences []int
}


/*--- Print formating ---*/

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

func SortMap(collect map[string]int) Slicepair {
		  // split the map into matching slices
		  length := len(collect)
		  textKey := make([]string, length)
		  occurs := make([]int, length)
		  i:= 0
		  for key, val := range collect {
					 textKey[i] = key
					 occurs[i] = val
					 i++

		  }

		  // mirror sort the slices
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
