package asstring

import (
		  "fmt"
)

type Slicepair struct {
		  Blocks []string
		  Occurences []int
}

/*---Pure Text Options---*/


// This should be once out of file
func frequencyMapAppend(collect map[string]int, newKey string) {

			 _, ok := collect[newKey]
			 if ok == true {
						collect[newKey]++
			 } else {
						collect[newKey] = 1 
			 }

}


func CountChars(text string) map[string]int {
		  collect := make(map[string]int)
		  length := len(text)

		  for i := 0; i < length; i++ {
					 char := string(text[i])
					 frequencyMapAppend(collect, char)
		  }
		  return collect
}

func CountNgrams(text string, letters int) map[string]int {
		  collect := make(map[string]int)
		  n := letters - 1
		  length := len(text) - n // Currently ignoring the final character

		  for i := 0; i < length; i++ {
					 bigram := string(text[i:i + n])
					 frequencyMapAppend(collect, bigram)
		  }
		  return collect
}

// Counts sentences, asentece asre defined by "!", "?" and "." also "... {Uppoer-case}"(technically not alway true as it could be a name)
func CountSentences(text []byte) map[string]int {
		  collect := make(map[string]int)
		  length := len(text) - 1 // Currently ignoring the final character

		  for i := 0; i < length; i++ {
					 bigram := string(text[i]) + string(text[i + 1])
					 frequencyMapAppend(collect, bigram)
		  }
		  return collect
}


func CountWords(text []byte) map[string]int {
		  collect := make(map[string]int)
		  length := len(text) - 1 // Currently ignoring the final character
		  var word []byte // each word will get strore in here
		  for i := 0; i < length; i++ {
					 if text[i] != 32 {
								word = append(word, text[i])
					 } else if word == nil {
								continue
					 }else {
								frequencyMapAppend(collect, PrintASCIIStr(word))
								word = nil
					 }
		  }
		  return collect
}

/*--- Print formating ---*/

func mirrorSort(unsortedPair Slicepair) Slicepair {

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
