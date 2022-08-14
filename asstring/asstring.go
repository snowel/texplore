package asstring

import (
		  "strings"
		 // tfmt "texplore/dataformat"
)

/*---Pure Text Options---*/

//TODO There are better approaches to cleaning the text than this func being called each time.
// Return a string with all occurences of the substring being removed.
func purgeSubStrings(target string, ex []string) string {
		  exLen := len(ex)
		  stringHolder := []string{target}

		  for i := 0; i < exLen; i++ {
					 if strings.Contains(stringHolder[0], ex[i]){
								rep := strings.NewReplacer(ex[i], "")
								stringHolder[0] = rep.Replace(stringHolder[0])
					 }
		  }

		  return stringHolder[0]
}

// Figure out if a rune is in a slice of searched.
func runeMatch(input rune, set []rune) bool {
		  length := len(set)

		  for i := 0; i < length; i++ {
					 if input == set[i] {
								return true
					 }
		  }

		  return false
}


func SimplifyString(fullstring string) string {
		  runeslice := []rune(strings.ToLower(fullstring))
		  var simplerunes []rune
		  for _, v := range runeslice {
					 num := int(v)
					 if num == 0x5F { // Underscore to hyphen
								simplerunes = append(simplerunes, '-')
					 } else if num == 0x3A {// Colon to semi
								simplerunes = append(simplerunes, ';')
					 } else if num == 0x3F {
								simplerunes = append(simplerunes, '/')
					 } else if num < 137 {
								simplerunes = append(simplerunes, v)
					 } else if num == 0x2019 || num == 0x2018{// single quote
								simplerunes = append(simplerunes, '\'')
					 } else if num == 0x201C || num == 0x201D{// double quote
								simplerunes = append(simplerunes, '\'') // CURRENTLY SET TO UNIFY QUOTES
					 }
		  }

		  return string(simplerunes)
}

// TODO factor out.
// Add an occurence counter of a block to a map.
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
		  n := letters
		  length := len(text) - (n - 1) // Currently ignoring the final character

		  for i := 0; i < length; i++ {
					 bigram := string(text[i:i + n])
					 frequencyMapAppend(collect, bigram)
		  }
		  return collect
}

// Counts sentences, a sentence is defined by "!", "?", "." also "... {Uppoer-case}" and "!?"
// TODO Doesn't quite work with quotes, also newlines are problemeatic and "..." will need to be addressed. I.E. Currently broken
func CountSentences(text string) map[string]int {
		  collect := make(map[string]int)
		  runestring := []rune(text)
		  length := len(runestring)
		  var (
					 sentence []rune
		  )
		  endChars := []rune("!?.")

		  for i := 0; i < length; i++ {
					 if runeMatch(runestring[i], endChars){
								sentence = append(sentence, runestring[i])
								frequencyMapAppend(collect, string(sentence))
								sentence = nil
					 } else {
								sentence = append(sentence, runestring[i])
					 }
		  }
		  return collect
}


func CountWords(text string) map[string]int {
		  collect := make(map[string]int)
		  runestring := []rune(text)
		  length := len(runestring)
		  var wordStartMark int
		  var readingWord bool
		  punctuation := []string{"!", ",", ".", ";", ":", "(", ")", " "}// hard coded TODO
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
