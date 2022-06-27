package main

import (
		  "os"
		  "log"

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
		  text := openFile("alice.txt")
		  charOcc := texploreASCII.CountBis(text)
		  bigramPair := texploreASCII.SortMap(charOcc)
		  biReps := texploreVoodoo.BigramEval(bigramPair, texploreVoodoo.Keymap2)
		  fingerRep:= texploreASCII.SortMap(biReps)
		  
		  //texploreASCII.PrintSlicepair(bits, occs)
		  texploreASCII.PrintSlicepair(fingerRep)
}
