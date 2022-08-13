# Texplore

Texplore is a simple CLI tool to explore the composition of a piece of text. The some of the text properties which can be pulled from the file with texplore including word frequency, character frequency, bigram frequency, etc

# Big Note!

Texplore was originally created as a tool to help optimize the defualt layout of the voodoo keyboard *(link comming soon)*.

These featuers are actively being played with and have caused much of the resonable features to be organized alongsize a wild mess of code and test UI. If you wish to use texplore please ignore all keymap optimization features.

### Features

#### Character frequency

The number of occurences of a given character in a body of text.

#### Ngrams

This ranks the requency of bigrams, trigrams, etc, which is to say, how often n chracters show up side by side in a piece of text.

Ngrams allow to determin how offten a keymap will requiere the typist to use the same finger n characters in a row.


#### Word Frequency

Counts the occurence of each word in the text. Case sensitive.
The following characters are purged from the text.

#### Sentences

**Currently quite broken**

Counts the occurences of sentences. Generally not very useful for occurences, but allows to break the text up into sentences.
