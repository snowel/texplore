package texploreASCII

import(
		  "fmt"
)

/*--- ASCII Enconded Text ---*/
/*---ASCII Text Options---*/
var whitespaceMap = map[byte]string {
		  9: "[tab]",
		  10: "[line-feed]",
		  32: "[space]",
}

// Takes a byte and returns a printalble represesntations of the ascii requivalent
func PrintASCII(char byte) string {
		  nonprint, ok := whitespaceMap[char]
		  if ok {
					 return nonprint
		  } else if char > 127 {
					 return "<UTF-8>"
		  } else {
					 return string(char)
		  }
}

func PrintASCIIStr(char []byte) string {
		  length := len(char)
		  str := ""

		  for i := 0; i < length; i++ {
					 nonprint, ok := whitespaceMap[char[i]]
					 if ok {
								str = str + nonprint
					 } else if char[i] > 127 {
								str = str + "<UTF-8>"
					 } else {
								str = str + string(char)
					 }
		  }

		  return str
}

func frequencyMapAppend(collect map[string]int, newKey string) {

			 _, ok := collect[newKey]
			 if ok == true {
						collect[newKey]++
			 } else {
						collect[newKey] = 1 
			 }

}


func CountChars(text []byte) map[string]int {
		  collect := make(map[string]int)
		  length := len(text)

		  for i := 0; i < length; i++ {
					 char := PrintASCII(text[i])
					 frequencyMapAppend(collect, char)
		  }
		  return collect
}

func CountBis(text []byte) map[string]int {
		  collect := make(map[string]int)
		  length := len(text) - 1 // Currently ignoring the final character

		  for i := 0; i < length; i++ {
					 bigram := PrintASCII(text[i]) + PrintASCII(text[i + 1])
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

func mirrorSort(elems []string, nums []int) ([]string, []int) {
//		  if len(elems) != len(nums) {
//					 return _, _, err
//		  }
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
		  return sortedElems, sortedNums
}

func SortMap(collect map[string]int) ([]string, []int) {
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
		  textKey, occurs = mirrorSort(textKey, occurs)

		  return textKey, occurs
}

func PrintSlicepair(text []string, occ []int) {
		  length := len(occ)
		  for i := 0; i < length; i++ {
					 fmt.Printf(":::  %s -> %d\n", text[i], occ[i])
		  }
}

