package main

import (
		  "os"
		  "log"  
		  "fmt"
		  

		  "texplore/ascii"
		  "texplore/voodoo"
)

func openFile(filename string) []byte{

		  f, ok := os.ReadFile(filename)
		  
		  if ok != nil {
					 log.Fatal(ok)
		  }

		  return f 
}


func main() {
		  argsLen := len(os.Args[1:])
		  if argsLen < 1 {
					 fmt.Println("Please specify a file.")
		  } else if argsLen > 1 {
					 fmt.Println("Please specify only one file.")
		  } else {
					 filename := os.Args[1]
					 text := openFile(filename)
					 charOcc := texploreASCII.CountBis(text)
					 bigramPair := texploreASCII.SortMap(charOcc)
					 biReps := texploreVoodoo.BigramEval(bigramPair, texploreVoodoo.Keymap2)
					 fingerRep:= texploreASCII.SortMap(biReps)
					 
					 //texploreASCII.PrintSlicepair(bits, occs)
					 texploreASCII.PrintSlicepair(fingerRep)
		  }
}
