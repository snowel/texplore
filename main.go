package main

import (
		  "fmt"
		  "os"
		  "log"
)

func openFile(filename string) []byte{

		  f, ok := os.ReadFile(filename)
		  
		  if ok != nil {
					 log.Fatal(ok)
		  }

		  return f 
}


/*---ASCII Text Options---*/
var whitespaceMap = map[byte]string {
		  9: "[tab]",
		  10: "[line-feed]",
		  32: "[space]",
}

// Takes a byte and returns a printalble represesntations of the ascii requivalent
func printASCII(char byte) string {
		  nonprint, ok := whitespaceMap[char]
		  if ok {
					 return nonprint
		  } else if char > 127 {
					 return "<UTF-8>"
		  } else {
					 return string(char)
		  }
}


func CountChars(text []byte) map[string]int {
		  collect := make(map[string]int)
		  length := len(text)

		  for i := 0; i < length; i++ {
					 char := printASCII(text[i])
					 _, ok := collect[char]
					 if ok == true {
								collect[char]++
					 } else {
								collect[char] = 1 
					 }
		  }
		  return collect
}

func CountBis(text []byte) map[string]int {
		  collect := make(map[string]int)
		  length := len(text) - 1 // Currently ignoring the final character

		  for i := 0; i < length; i++ {
					 bigram := printASCII(text[i]) + printASCII(text[i + 1])
					 _, ok := collect[bigram]
					 if ok == true {
								collect[bigram]++
					 } else {
								collect[bigram] = 1 
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

func sortMap(collect map[string]int) ([]string, []int) {
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

func printMap(text []string, occ []int) {
		  length := len(occ)
		  for i := 0; i < length; i++ {
					 fmt.Printf("%s -> %d\n", text[i], occ[i])
		  }
}

func main() {
		  text := openFile("alice.txt")
		  charOcc := CountBis(text)
		  fmt.Println(charOcc)
		  bits, occs := sortMap(charOcc)

		  
		  printMap(bits, occs)
}
