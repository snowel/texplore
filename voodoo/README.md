# Keymap Optimization

NOTE: in all cases, each keystroke are represented by a `string`

## Finger Repetition Opt

Finger rep optimization minimizes the back-to-back use of a given finger.

i.e. For a standard qwerty keybaord, while typing with the "standard" touch-typing technique. The word "king" has a repetition of the right middle finger. 


### Type Sematics

A keymap for finger rep optimization is a `[8][]string`

Each slice of the array is the list of keystrokes that finger is reposnible for. The fingers are indetified by their index in the array.

In the future, I'll be moving away from the hardcoded array length to optimize a map for an arbitrary numebr of fingers.


## layer Switching Opt

Layer switching optimization is mostly relevant to super-minimalist keyboards like the Voodoo it was originally designed for, but will neverthelsee work for any format.

### Type semantics

A keymap for layer switch optimization is a `[][]string`

Each element of the slice is the list of keystrokes on a given keymap. NOTE: for `n` keys you want to place in a map `kMap`, `len(kMap) * len(kMap[0]) == n` In other words, the number of layers `m` should satisfy `n % m == 0`

