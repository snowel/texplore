package texploreVoodoo

/*--- 
		  This is a package for testing voodoo layouts.
		  -repeating single finger use 
		  -repeating same hand use
		  -etc
---*/


import (
		  "C"
)

var (

		  EmptyMap = make(map[string][]string)
		  OffHomerow = []string{"b", "g", "h", "j", "k", "m", "l", "p", "q", "v", "x", "y", "z", "w", "d", "c", "f", "u"}
		  SingleLayer2 = []string{"h", "n", "e", "m", "i", "l", "o", "w", "t", "d", "r", "c", "s", "f", "a", "u"}
		  SingleLayer1 = []string{"h", "m", "l", "w", "d", "c", "f", "u"}
	
// Finger REP
		  NaiveHumanFingerMap = [8][]string{
					 []string{"y", "r", "c", "j"},
					 []string{"s", "x", "f", "k"},
					 []string{"o", "w", "{[enter]}", "{[backspace]}"},
					 []string{"t", "d", "b", "g"},
					 []string{"a", "u", "z", "q"},
					 []string{"e", "m", ",", ";"},
					 []string{"i", "u", "y", "-"},
					 []string{"n", "h", "p", "v"},
		  }
		 
// Based on a limited corpus

		  CorpusV1 = [8][]string{
					 []string{"h", "l", "m", "j"},
					 []string{"s", "c", "f", "z"},
					 []string{"o", "g", "y", "'"},
					 []string{"t", "d", "p", "q"},
					 []string{"a", "u", "k", "-"},
					 []string{"e", "w", ".", "\t"},
					 []string{"i", ",", "b", "\n"},
					 []string{"n", "r", "v", "x"},
		  }

		  Corp1Block1 = []string{"e", "t", "a", "i", "n", "o", "s", "h"}
		  Corp1Block2 = []string{"r", "l", "d", "u", "g", "c", "w", ","}
		  Corp1Block3 = []string{"f", "m", "p", "y", "b", ".", "k", "v"} 
		  Corp1Block4 = []string{"x", "\n", "q", "'", "j", "z", "-", "\t"} 


// Layer switching

		  // 32 moest comment chars from corpus 1
		  Corpus1Set = []string{"e", "t", "a", "i", "n", "o", "s", "h",
										  "r", "l", "d", "u", "g", "c", "w", ",",
										  "f", "m", "p", "y", "b", ".", "k", "v",
										  "x", "\n", "q", "'", "j", "z", "-", "\t"} 
)






