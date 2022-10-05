package main

import (
		  "os"
		  "log"  
		  "fmt"
		  "flag"
		  "strings"

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

		  if len(os.Args[1:]) < 1 {
					 // Layer permutaitons

					 var acc [][][]string


					 set := []string{"A", "B", "C", "D"}
					 voodoo.AllLayerMaps(set, 2, acc)
					 //voodoo.AllLayerMaps(voodoo.Corpus1Set, 16, acc)
					 fmt.Println(acc)
					 

					 //text := openFile("JL_Corpus.txt")
					 //text = tstring.SimplifyString(text)
					 

					 //ngramOcc := tstring.CountNgrams(text, 2)
					 //grams := tfmt.SortMap(ngramOcc)
					 

					 //singleMap := voodoo.Corp1Block4
					 //smaps := voodoo.HeapSMap(singleMap)
					 //col := voodoo.EvalArrMaps(smaps, voodoo.CorpusV1, grams)
					 //voodoo.NSmallestTotalRep(col, 10)
					 //voodoo.NSmallestRep(col, 10)

					 //eval := voodoo.ArrFingerEval(text, voodoo.CorpusV1)
					 //fmt.Println(eval)
					 //fmt.Println(voodoo.ArrmapEvalSum(eval))
					 //fmt.Println(voodoo.CorpusV1)
					 return
		  }

		  filename := flag.String("f", "", "Name of the file to explore.")
		  mode := flag.Int("m", 0, "Infomation you want out of the text. 0: Character freq\n 1: Ngrams\n 2: Word freq\n")
		  lowerCase := flag.Int("case", 0, "Set all characters to lower case (as to not differentiate) 0: no\n 1: yes\n")
		  simple := flag.Int("simp", 0, "Set all characters to lower case and only ascii compatible for input evals 0: no\n 1: yes\n")
		  Ngram := flag.Int("ngram", 0, "Number of characters per group")

		  flag.Parse()

		  if *filename == "" {
					 fmt.Println("A file must be specified with the file name flag: -f=alice.txt")
					 return
		  }
		  text := openFile(*filename)
		  switch *mode {
		  case 0: {// This mode will probably be depreciated, as it can be the default of ngam=1.
					 if *lowerCase == 1 {
								text = strings.ToLower(text)
					 }
					 if *simple == 1 {
								text = tstring.SimplifyString(text)
					 }
					 charOcc := tstring.CountChars(text)
					 chars := tfmt.SortMap(charOcc)
					 totalChars := tfmt.SlicepairOccSum(&chars)
					 fmt.Println("This text has a total of -> ", totalChars, " <- characters.")
					 tfmt.PrintSlicepair(chars)
					 
		  }
		  case 1: {
					 if *Ngram == 0 {
								fmt.Println("Please specify the number of character you want to group together by using the ngram flag: -ngram=2")
								break
					 }
					 if *lowerCase == 1 {
								text = strings.ToLower(text)
					 }
					 if *simple == 1 {
								text = tstring.SimplifyString(text)
					 }
					 ngramOcc := tstring.CountNgrams(text, *Ngram)
					 grams := tfmt.SortMap(ngramOcc)
					 tfmt.PrintSlicepair(grams)
		  }
		  case 2: {
					 if *lowerCase == 1 {
								text = strings.ToLower(text)
					 }
					 if *simple == 1 {
								text = tstring.SimplifyString(text)
					 }
					 wordOcc := tstring.CountWords(text)
					 words := tfmt.SortMap(wordOcc)
					 totalWords := tfmt.SlicepairOccSum(&words)
					 fmt.Println("This text has a total of -> ", totalWords, " <- words.")
					 tfmt.PrintSlicepair(words)
		  }
		  case 3: {
					 if *lowerCase == 1 {
								text = strings.ToLower(text)
					 }
					 if *simple == 1 {
								text = tstring.SimplifyString(text)
					 }
					 sentOcc := tstring.CountSentences(text)
					 sentences := tfmt.SortMap(sentOcc)
					 totalSentences := tfmt.SlicepairOccSum(&sentences)
					 fmt.Println("This text has a total of -> ", totalSentences, " <- sentences.")
					 tfmt.PrintSlicepair(sentences)
		  }
		  case 4: {
					 if *Ngram == 0 {
								fmt.Println("Please specify the number of character you want to group together by using the ngram flag: -ngram=2")
								break
					 }
					 if *lowerCase == 1 {
								text = strings.ToLower(text)
					 }
					 if *simple == 1 {
								text = tstring.SimplifyString(text)
					 }
					 // Voodoo options here
		  }
		  }
		  
}
