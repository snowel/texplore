package main

import (
		  "os"
		  "log"  
		  "fmt"
		  

		  //ascii "texplore/ascii"
		  //voodoo "texplore/voodoo"
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
		  argsLen := len(os.Args[1:])
		  if argsLen < 1 {
					 fmt.Println("Please specify a file.")
		  } else if argsLen > 1 {
					 fmt.Println("Please specify only one file.")
		  } else {
					 filename := os.Args[1]
					 text := string(openFile(filename))
					 wordOcc := tstring.CountChars(text)
					 //charOcc := ascii.CountBis(text)
					 words := tfmt.SortMap(wordOcc)
					 //biReps := voodoo.BigramEval(bigramPair, voodoo.Keymap2)
					 //fingerRep:= tfmt.SortMap(biReps)
					 
					 //texploreASCII.PrintSlicepair(bits, occs)
					 tfmt.PrintSlicepair(words)
		  }
}
