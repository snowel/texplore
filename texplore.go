package main

import (
		  "os"
		  "log"  
		  "fmt"
		  "flag"

		  voodoo "texplore/voodoo"
		  tfmt "texplore/dataformat"
		  tstring "texplore/asstring"
)

func openFile(filename string) string{

		  f, ok := os.ReadFile(filename)
		  
		  if ok != nil {
					 log.Fatal(ok)
		  }

		  return string(f) 
}

func main() {
		  filename := flag.String("f", "", "Name of the file to explore.")
		  mode := flag.Int("m", 0, "Infomation you want out of the text. 0: Character freq\n 1: Ngrams\n 2: Word freq\n")
		  Ngram := flag.Int("ngram", 0, "Number of characters per group")
//TODO - Better UI
// is there a way to hybridize the flags + args notation???
//argsLen := len(os.Args[1:])
//if argsLen < 1 {
//		 fmt.Println("Please specify a file.")
// } else if argsLen > 1 {
//			 fmt.Println("Please specify only one file.")
//} else {}
//filename := os.Args[1]

		  flag.Parse()// TODO error check?

		  if *filename == "" {
					 fmt.Println("A file must be specified with the file name flag: -f=alice.txt")
					 return
		  }
		  text := openFile(*filename)
		  switch *mode {
		  case 0: {// this mode will probably be depreciated, as it can be the defualt of ngam=1
					 charOcc := tstring.CountChars(text)
					 chars := tfmt.SortMap(charOcc)
					 tfmt.PrintSlicepair(chars)
					 
		  }
		  case 1: {
					 if *Ngram == 0 {
								fmt.Println("Please specify the number of character you want to group together by using the ngram flag: -ngram=2")
								break
					 }
					 ngramOcc := tstring.CountNgrams(text, *Ngram)
					 grams := tfmt.SortMap(ngramOcc)
					 tfmt.PrintSlicepair(grams)
		  }
		  case 2: {
					 wordOcc := tstring.CountWords(text)
					 words := tfmt.SortMap(wordOcc)
					 tfmt.PrintSlicepair(words)
		  }
		  case 3: {
					 if *Ngram == 0 {
								fmt.Println("Please specify the number of character you want to group together by using the ngram flag: -ngram=2")
								break
					 }
					 ngramOcc := tstring.CountNgrams(text, *Ngram)
					 grams := tfmt.SortMap(ngramOcc)
					 biReps := voodoo.BigramEval(grams, voodoo.Keymap2)
					 fingerRep:= tfmt.SortMap(biReps)
					 tfmt.PrintSlicepair(fingerRep)
		  }
		  }
		  
}
