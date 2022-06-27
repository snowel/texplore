package main

import (
		  "fmt"
		  "os"
		  "log"

		  "texplore/ascii"
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
		  fmt.Println(charOcc)
		  bits, occs := texploreASCII.SortMap(charOcc)

		  
		  texploreASCII.PrintSlicepair(bits, occs)
}
