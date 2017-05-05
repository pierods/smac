
# SMAC - Small autocompletion engine in Go.

SMAC is a tiny autocompletion engine written in Go. It supports UTF-8 alphabets. 

Performance on a 355k word dictionary on a modern computer is :

* initialization time ~1.5 seconds
* average completion time 0.05 ms (20k completions/sec)
* average nodes/words ratio is 2.7

SMAC has not been tested on logographic alphabets, for which it would not make much sense, unless 
an ortographic spelling is provided (hiragana to kanji for example). If there is such a need SMAC 
could be adapted to it.



Rune tables only store the offset of the rune ordinal, with respect to the lowest rune in the set of characters provided at initialization time and the word tree only stores the rune value, so to minimize the space used to store the input dictionary.

