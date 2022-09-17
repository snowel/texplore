package texploreVoodoo
import (
		  "reflect"
		  tfmt "texplore/dataformat"
)

// --- Procedural permutaitons

// Takes a slicemap collection and evalueates each into a freq map

func EvalMaps(slicemaps [][]string, basemap map[string][]string, ref tfmt.Slicepair) ([]map[string][]string, []map[string]int) {
		  length := len(slicemaps)
		  evaluations := make([]map[string]int, length)
		  maps := make([]map[string][]string, length)
		  for i := 0; i < length; i++ {
					 keymap := SliceToKeymap(slicemaps[i], basemap)
					 eval := BigramEval(ref, keymap)
					 evaluations[i] = eval
					 maps[i] = keymap
					 //fmt.Println(i, " of ", length)
		  }
		  return maps, evaluations
}

//Heaps takes in a slice of chars and produces a slice of all possible orders
func HeapSMap(chars []string) [][]string {
		  var collect [][]string
		  mut := chars
		  Heaps(len(chars), &mut, &collect)
		  return collect
}

// Heaps algo, recursive
func Heaps(k int, arr *[]string, permutations *[][]string) {
		  if k == 1 {
					 //fmt.Println(*arr)
					 arrcp := make([]string, len(*arr))
					 copy(arrcp, *arr)
					 *permutations = append(*permutations, arrcp)
		  } else {
					 swap := reflect.Swapper(*arr)
					 Heaps(k - 1, arr, permutations)
					 
					 for i := 0; i < k - 1; i++ {
								if k % 2 == 0 {
										  swap(i, k - 1)
								} else {
										  swap(0, k - 1)
								}

								Heaps(k - 1, arr, permutations)
					 }
		  }
}

// If there is no forced map, eval mats is an empty map = make(map[string][]string)
// Takes a slice of strings and retruns a keymap, appending each block of n characters to a certain finger
func SliceToKeymap(slicemap []string, keymap map[string][]string) map[string][]string {
		  length :=  len (slicemap)
		  block := (length / 8) - length % 8// Need to hadle overflow
		  
		  for i := 0; i < length; i++ {
					 switch {
					 case i < block: tfmt.KeymapAppend("ri", slicemap[i], keymap)
					 case i < 2*block: tfmt.KeymapAppend("rm", slicemap[i], keymap)
					 case i < 3*block: tfmt.KeymapAppend("rr", slicemap[i], keymap)
					 case i < 4*block: tfmt.KeymapAppend("rp", slicemap[i], keymap)
					 case i < 5*block: tfmt.KeymapAppend("li", slicemap[i], keymap)
					 case i < 6*block: tfmt.KeymapAppend("lm", slicemap[i], keymap)
					 case i < 7*block: tfmt.KeymapAppend("lr", slicemap[i], keymap)
					 case i < 8*block: tfmt.KeymapAppend("lp", slicemap[i], keymap)
					 }
		  }
		  return keymap
}
