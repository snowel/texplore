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
		  bits, occs := texploreASCII.SortMap(charOcc)
		  biReps := texploreVoodoo.BigramEval(bits, occs, texploreVoodoo.Keymap1)
		  finger, repets := texploreASCII.SortMap(biReps)
		  
		  texploreASCII.PrintSlicepair(finger, repets)
}
