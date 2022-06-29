# Texplore

Texplore is a simple CLI tool to explore the composition of a piece of text. The some of the text properties which can be pulled from the file with texplore including word frequency, character frequency, bigram frequency, etc

Texplore was originally created as a tool to help optimize the defualt layout of the voodoo keyboard *(link comming soon)*.

### Features

#### Character frequency

The number of occurences of a given character in a body of text.

#### Bigrams

This ranks the requency of bigrams, which is to say, how often two chracters show up side by side in a peice of text.

Th purpose of checking bigrams is to determin how offten a keymap will requiere the typist to use the same finger 2 characters in a row. To check this, a psuedo-keyma is made, which assicates every character with a finger. Currently psuedo-keymaps are hardcodded.
