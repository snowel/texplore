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
		  
		  TestLayerMap = [][]string{
					 []string{"h", "m", "l", "w", "d", "c", "f", "u"},
					 []string{"r", "l", "d", "u", "m", "c", "w", "g"}, 
					 []string{"f", "y", "p", ",", "b", ".", "k", "v"}, 
					 []string{"\"", "'", "_", "-", ";", "!", "j", "?"}, 
		  }
		  
		  MachineMix1 = [8][]string{
					 []string{"h"},
					 []string{"s"},
					 []string{"o"},
					 []string{"t"},
					 []string{"a"},
					 []string{"e"},
					 []string{"i"},
					 []string{"n"},
		  }

		  MachineMix2 = [8][]string{
					 []string{"h", "r"},
					 []string{"s", "c"},
					 []string{"o", "d"},
					 []string{"t", "m"},
					 []string{"a", "g"},
					 []string{"e", "w"},
					 []string{"i", "u"},
					 []string{"n", "l"},
		  }

		  MachineMix3 = [8][]string{
					 []string{"h", "r", "v", ";"},
					 []string{"s", "c", "f", "j"},
					 []string{"o", "d", "k", "\""},
					 []string{"t", "m", "b", "!"},
					 []string{"a", "g", ",", "'"},
					 []string{"e", "w", ".", "?"},
					 []string{"i", "u", "y", "-"},
					 []string{"n", "l", "p", "_"},
		  }
		  MachineRow2 = []string{"r", "l", "d", "u", "m", "c", "w", "g"} 
		  MachineRow3 = []string{"f", "y", "p", ",", "b", ".", "k", "v"} 
		  MachineRow4 = []string{"\"", "'", "_", "-", ";", "!", "j", "?"} 

		  ArrayMap1 = [8][]string{
					 []string{"h", "n"},
					 []string{"e", "m"},
					 []string{"i", "l"},
					 []string{"o", "w"},
					 []string{"t", "d"},
					 []string{"r", "c"},
					 []string{"s", "f"},
					 []string{"a", "u"},
		  }
		  ArrayMap2 = [8][]string{
					 []string{"a"},
					 []string{"s"},
					 []string{"r"},
					 []string{"t"},
					 []string{"n"},
					 []string{"e"},
					 []string{"i"},
					 []string{"o"},
		  }

		  SingleLayerMap = map[string][]string{
					 "ri": {"h"},
					 "rm": {"m"},
					 "rr": {"l"},
					 "rp": {"w"},
					 "li": {"d"},
					 "lm": {"c"},
					 "lr": {"f"},
					 "lp": {"u"},
		  }
		  SingleLayerMap2 = map[string][]string{
					 "ri": {"h", "n"},
					 "rm": {"e", "m"},
					 "rr": {"i", "l"},
					 "rp": {"o", "w"},
					 "li": {"t", "d"},
					 "lm": {"r", "c"},
					 "lr": {"s", "f"},
					 "lp": {"a", "u"},
		  }

		  Keymap1 = map[string][]string{
					 "Right-Index": {"n", "h", "p", "k"},
					 "Right-Middle": {"e", "m", ":", ","},
					 "Right-Ring": {"i", "l", ".", "'", "\""},
					 "Right-Pinky": {"o", "w", "x" },
					 "Left-Index": {"t", "d", "b", "g" },
					 "Left-Middle": {"r", "R", "y", "Y", "v", "V", "c", "C"},
					 "Left-Ring": {"s", "S", "f", "F", "k", "K", "j", "J"},
					 "Left-Pinky": {"a", "A", "u", "U", "q", "Q", "z", "Z"},

		  }
		  Keymap2 = map[string][]string{
					 "Right-Index": {"n", "N", "h", "H", "p", "P", "m", "M"},
					 "Right-Middle": {"e", "E", "l", "L", ";", ":", ","},
					 "Right-Ring": {"i", "I", " ", ".", "'", "\""},
					 "Right-Pinky": {"o", "O", "w", "W", "x", "X"},
					 "Left-Index": {"t", "T", "d", "D", "b", "B", "g", "G"},
					 "Left-Middle": {"r", "R", "y", "Y", "v", "V", "c", "C"},
					 "Left-Ring": {"s", "S", "f", "F", "k", "K", "j", "J"},
					 "Left-Pinky": {"a", "A", "u", "U", "q", "Q", "z", "Z"},

		  }
		  Keymap3 = map[string][]string{
					 "Right-Index": {"n", "N", "h", "H", "p", "P", "m", "M"},
					 "Right-Middle": {"e", "E", "l", "L", ";", ":", ","},
					 "Right-Ring": {"i", "I", " ", ".", "'", "\""},
					 "Right-Pinky": {"o", "O", "w", "W", "x", "X"},
					 "Left-Index": {"t", "T", "d", "D", "b", "B", "g", "G"},
					 "Left-Middle": {"r", "R", "y", "Y", "v", "V", "c", "C"},
					 "Left-Ring": {"s", "S", "f", "F", "k", "K", "j", "J"},
					 "Left-Pinky": {"a", "A", "u", "U", "q", "Q", "z", "Z"},

		  }
		  Keymap4 = map[string][]string{
					 "Right-Index": {"n", "N", "h", "H", "p", "P", "m", "M"},
					 "Right-Middle": {"e", "E", "y", ";", ":", ","},
					 "Right-Ring": {"i", "I", "l", "L", ".", "'", "\""},
					 "Right-Pinky": {"o", "O", "w", "W", "x", "X"},
					 "Left-Index": {"t", "T", "d", "D", "b", "B", "g", "G"},
					 "Left-Middle": {"r", "R", "y", "Y", "v", "V", "c", "C"},
					 "Left-Ring": {"s", "S", "f", "F", "k", "K", "j", "J"},
					 "Left-Pinky": {"a", "A", "u", "U", "q", "Q", "z", "Z"},
		  }
)






