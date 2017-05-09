
# SMAC - Small autocompletion engine in Go.

SMAC is a tiny autocompletion engine written in Go. It supports UTF-8 alphabets. Emphasis is on speed and simplicity.

Performance on a 355k word dictionary on a modern computer is :

* initialization time ~1.5 seconds
* average completion time 0.05 ms (20k completions/sec)
* average nodes/words ratio is 2.7

Rune tables only store the offset of the rune ordinal, with respect to the lowest rune in the set of characters provided at initialization time and the word tree only stores the rune value, so to minimize the space used to store the input dictionary.

SMAC is case-sensitive. If case-insensitivity is needed, clients should make sure to convert to either upper or lowercase before adding words and calling Complete().

Paging is not supported, since it is mostly responsibility of the client.

SMAC learns new words on the go, via the Learn() function, can UnLearn() them and can also UnLearn() words provided in the bootstrap dictionary, and keeps into account the frequency of acceptance of words (frequently used words) by giving them priority when generating completion lists.

If customizations of the autocompleter (learnt words etc.) are to survive re-instantiations, or they must be transferred somewhere else, SMAC can Save() what it has learnt to file and can subsequently Retrieve() them from file.

SMAC has not been tested on logographic alphabets, for which it would not make much sense, unless 
an ortographic spelling is provided (hiragana to kanji for example). If there is such a need SMAC 
could be adapted to it.
