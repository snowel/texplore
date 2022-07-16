# Texplore

Texplore is a simple CLI tool to explore the composition of a piece of text. The some of the text properties which can be pulled from the file with texplore including word frequency, character frequency, bigram frequency, etc

Texplore was originally created as a tool to help optimize the defualt layout of the voodoo keyboard *(link comming soon)*.

### Features

#### Character frequency

The number of occurences of a given character in a body of text.

#### Ngrams

This ranks the requency of bigrams, trigrams, etc, which is to say, how often n chracters show up side by side in a piece of text.

Ngrams allow to determin how offten a keymap will requiere the typist to use the same finger n characters in a row.


#### Word Frequency

Counts the occurence of each word in the text. Case sensitive.
The following characters are purged from the text.
*;
*:
*!
*,
*(
*)
