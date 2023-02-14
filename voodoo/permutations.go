package texploreVoodoo
import (
		  "reflect"
		  tfmt "texplore/dataformat"
		  "fmt"
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


//  -- Layer switching combinaitons


// All possible layer seperated maps are created by taking
// a pool of characters (an array of length n) and generatign sets of length n/m of layers of length m.
// All possible layer seperated maps are then created by combining every layer from an array of layers
// SEMANTICS:
// n is the numebr of characters we're placing
// m is the number of character tha fit on one layer
// ASSUMPTIONS:
// These layers are unordered sets
// n % m == 0
// all elements in the char-pool array are unique
/*
i.e.
for [1, 2, 3, 4]
and we want layers of size 2, we get:
[[1,2], [3,4]]
[[1,4], [3,2]]
[[1,3], [4,2]]

*/

// In the specific case wehre n/m == 2, which is to say we have 2 layers, we can clauclate them this way:
// Generate all possible length m layers, and each time generate a matching layer with all elements not included
// This can be repeated onto a sub array to break down into mutiple layer:
// if we have a 12 character pool and want to seperate them into every possible layer sep map with layers that are 4 character wide we do the follwoing:
// generate one possible combination of 4 character from the 12 char pool, save it,
// recursively generate a combinaiton of 4 characters form the remainins 8 characters, save it and save the remaining characters
// then return the three saved arrays as a layer seperated keymap
// The problem with this apporach is that I will end up with repreat keymaps... a lot of them.
// This should be avoidable by forcing the generations of possible cominations to use the first element of the pool in all it's outputed combinaitons
// The resonaing for this is that a layer witht the first element MUST occur,
// and if it occurs, all other layers must occur


// All possible unordered combinaitons of length m (layerSize) for a given slice of strings
//TODO acc could be [][m]string, a way to pre-copile confude these array funcitons, like with deinfes, would also apply to fingermaps for finger repeat
func AllLayers(charSet []string, layerSize int, acc *[][]string) int {

		  setLen := len(charSet)
		  if setLen % layerSize != 0 {return 1}
		  buff := make([]string, layerSize)

		  unorderedHelp(charSet, setLen, layerSize, acc, &buff, 0, 0)
		  return 0
}

func unorderedHelp(charSet []string, setLen int, layerSize int, acc *[][]string, buff *[]string, buffIndex int, setIndex int) {

		  // add current temp buff to acc
		  if buffIndex == layerSize {
					 layer := make([]string, layerSize)
					 copy(layer, *buff)
					 *acc = append(*acc, *buff)
					 return
		  }
		  
		  // end of the set
		  if setIndex >= setLen {return}

		  // Set current buff element to match set elem.
		  (*buff)[buffIndex] = charSet[setIndex]
		  unorderedHelp(charSet, setLen, layerSize, acc, buff, buffIndex + 1, setIndex + 1)

		  // Set current buff elem to match next set elem.
		  unorderedHelp(charSet, setLen, layerSize, acc, buff, buffIndex, setIndex + 1)
}

// All possible combinaitons of layers
func AllLayerMaps(charSet []string, layerSize int, acc *[][][]string) int {

		  setLen := len(charSet)
		  if setLen % layerSize != 0 {return 1}
		  buff := make([]string, layerSize)
		  var leftovers []string
		  layerCount := setLen / layerSize


		  allFirstLayers(charSet, &leftovers, layerCount, setLen, layerSize, acc, &buff, 1, 1)
		  return 0
}

// Helper creating the first of each 
func allFirstLayers(charSet []string, remainingChars *[]string, layerCount int, setLen int, layerSize int, acc *[][][]string, buff *[]string, buffIndex int, setIndex int) {

		  if buffIndex == layerSize {
					 fmt.Println("First layer, charSet: ", charSet)
					 layer := make([]string, layerSize)
					 copy(layer, *buff)
					 var layerAcc [][]string
					 layerAcc = append(layerAcc, *buff)
					 fmt.Println("First layers, layer accumulator: ", layerAcc)

					 charLeft := len(*remainingChars)
					 //complete the remaining chars
					 if charLeft != setLen - layerSize {*remainingChars = append(*remainingChars, charSet...)}
					 
					 newSet := make([]string, charLeft)
					 copy(newSet, *remainingChars)
					 fmt.Println("First layer, new set:", newSet)
					 fmt.Println("First layer, remaining chars: ", *remainingChars)
					 newLen := len(newSet)
					 depth := 1


					 *remainingChars = nil // reset remain buffer
					 allOtherLayers(newSet, remainingChars, depth, layerCount, newLen, layerSize, &layerAcc, buff, 1, 1)

					 *acc = append(*acc, layerAcc)
					 // Doesn't need to be copied here as this version of layerAcc will immediiately go out of scope and never be changed.

					 *remainingChars = nil // reset remain buffer
					 return
		  }
		  
		  // end of the set
		  if setIndex >= setLen {return}
		  

		  // First element is always the first element of the set.
		  // TODO this is redundant effort, but it limits the arguments required
		  (*buff)[0] = charSet[0]

		  // Set current buff element to match set elem.
		  (*buff)[buffIndex] = charSet[setIndex]
		  allFirstLayers(charSet, remainingChars, layerCount, setLen, layerSize, acc, buff, buffIndex + 1, setIndex + 1)

		  // Set current buff elem to match next set elem.
		  *remainingChars = append(*remainingChars, charSet[setIndex])
		  allFirstLayers(charSet, remainingChars, layerCount, setLen, layerSize, acc, buff, buffIndex, setIndex + 1)
}

// The helper for recusing. Effectively creates all the other layers for a given layer.
func allOtherLayers(charSet []string, remainingChars *[]string, depth int, layerCount int, setLen int, layerSize int, acc *[][]string, buff *[]string, buffIndex int, setIndex int) {

		  // Once our buffer is full we recurse into creting the next layer from the remaining characters.
		  if buffIndex == layerSize {
					 // The layer found and stored in the buffer is appended to the layer *accumulator.
					 
					 fmt.Println("All other layers, accumulator: ", *acc)
					 var layer []string
					 copy(layer, *buff)
					 
					 *acc = append(*acc, layer)

					 if len(*remainingChars) != setLen - (layerSize * depth) { *remainingChars = append(*remainingChars, charSet[setIndex:]...)} // add trailing elements

					 // Are we done?
					 if depth == layerCount { return }
					
					 // Still more layers to find

					 var newSet []string
					 copy (newSet, *remainingChars)
					 newLen := len(newSet)
					 newDepth := depth + 1

					 *remainingChars = nil // reset remain buffer
					 allOtherLayers(newSet, remainingChars, newDepth, layerCount, newLen, layerSize, acc, buff, 1, 1)
					 return
		  }
		  
		  // end of the set; avoid inifnite loop
		  if setIndex >= setLen {return}

		  (*buff)[0] = charSet[0]
		  
		  // Set current buff element to match set elem.
		  (*buff)[buffIndex] = charSet[setIndex]
		  allOtherLayers(charSet, remainingChars, depth, layerCount, setLen, layerSize, acc, buff, buffIndex + 1, setIndex + 1)

		  // Set current buff elem to match next set elem.
		  *remainingChars = append(*remainingChars, charSet[setIndex])
		  allOtherLayers(charSet, remainingChars, depth, layerCount, setLen, layerSize, acc, buff, buffIndex, setIndex + 1)
}
