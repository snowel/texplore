package asstring

import (
		  "strings"
)

type Slicepair struct {
		  Blocks []string
		  Occurences []int
}

/*---Pure Text Options---*/

// Return a strin with all occurences of the substring being removed
func purgeSubStrings(target string, ex []string) string {
		  exLen := len(ex)
		  stringHolder := []string{target}

		  for i := 0; i < exLen; i++ {
					 if strings.Contains(stringHolder[0], ex[i]){
								rep := strings.NewReplacer(ex[i], "")// I think this whole funciton can be just one replacer...TODO
								stringHolder[0] = rep.Replace(stringHolder[0])
					 }
		  }

		  return stringHolder[0]
}

// is this rune in my slice of searched runes?
func runeMatch(input rune, set []rune) bool {
		  length := len(set)

		  for i := 0; i < length; i++ {
					 if input == set[i] {
								return true
					 }
		  }

		  return false
}


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


func CountWords(text string) map[string]int {
		  collect := make(map[string]int)
		  // if I can't change the string []string will work
		  runestring := []rune(text)
		  length := len(runestring)
		  var wordStartMark int
		  var readingWord bool
		  punctuation := []string{"!", ",", ".", ";", ":", "(", ")", " "}// currently hard coded TODO
		  wordDelimit := []rune(" ")

		  for i := 0; i < length; i++ {
					 if !runeMatch(runestring[i], wordDelimit) && wordStartMark != i && !readingWord {
								wordStartMark = i
								readingWord = true
					 } else if runeMatch(runestring[i], wordDelimit) && wordStartMark != i && readingWord {
								word := string(runestring[wordStartMark:i])
								frequencyMapAppend(collect, purgeSubStrings(word, punctuation))
								readingWord = false
					 }
		  }
		  return collect
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
