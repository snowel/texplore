package texploreVoodoo

import (
		  tfmt "texplore/dataformat"
)

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
