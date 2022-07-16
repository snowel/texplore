package texploreASCII


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
					 char := string(text[i])
					 frequencyMapAppend(collect, char)
		  }
		  return collect
}

func CountBis(text []byte) map[string]int {
		  collect := make(map[string]int)
		  length := len(text) - 1 // Currently ignoring the final character

		  for i := 0; i < length; i++ {
					 bigram := string(text[i]) + string(text[i + 1])
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


