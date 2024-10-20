package main

import (
	"flag"
	en "texplore/encod"
		  //voodoo "texplore/voodoo"
		  // tfmt "texplore/dataformat"
		  //tstring "texplore/asstring"
)

func main() {
//if len(os.Args[1:]) < 1 {
// Layer permutaitons

//var acc [][]string
//set := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L"}
//voodoo.AllLayerMaps(set, 2, &acc)
//voodoo.AllLayers(set, 3, &acc)
//fmt.Println(acc)
//voodoo.AllLayerMaps(voodoo.Corpus1Set, 16, acc)
//fmt.Println(acc)


//text := OpenFile("JL_Corpus.txt")
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
//return
//}

	input:= flag.String("f", "", "Name of the file to load..")
	output := flag.String("o", "wiseword-output", "Name of the file to output.")
//mode := flag.Int("m", 0, "Infomation you want out of the text. 0: Character freq\n 1: Ngrams\n 2: Word freq\n")
//lowerCase := flag.Int("case", 0, "Set all characters to lower case (as to not differentiate) 0: no\n 1: yes\n")
//simple := flag.Int("simp", 0, "Set all characters to lower case and only ascii compatible for input evals 0: no\n 1: yes\n")
//Ngram := flag.Int("ngram", 0, "Number of characters per group")

// Ecoding flags

	encoding := flag.String("e", "", "Path to an encoding yaml file use to encode or decode text.")

	flag.Parse()
	cmd := flag.Args()
	switch cmd[0] {
	case "ex": { // Explore mode
		}
	case "en": { // Encode mode
		en.EnUI(*input, *output, *encoding)	
		}
	case "de": {
		en.DeUI(*input, *output, *encoding)
	}
	case "km": { // Keymap mode
		
		}	
	}
		  
}
