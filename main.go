package main

import (
		  "fmt"
		  "os"
		  "log"

		  "test/texplore/ascii/ascii.go"
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
		  charOcc := CountBis(text)
		  fmt.Println(charOcc)
		  bits, occs := SortMap(charOcc)

		  
		  printMap(bits, occs)
}
