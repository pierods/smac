
# SMAC - Small autocompletion engine in Go.

SMAC is a tiny autocompletion engine written in Go. It supports UTF-8 alphabets. Emphasis is on speed and simplicity.

Performance on a 355k word dictionary on a modern computer is (see benchmark files):

* initialization time ~0.7 seconds
* average completion time 9000 ns (> 100k completions/sec)
* average memory occupation/dictionary size ratio varies between 13 and 16 (prefixMapDepth = 3 and 4 respectively)

SMAC is case-sensitive. If case-insensitivity is needed, clients should make sure to convert to either upper or lowercase before adding words and calling Complete().

Paging is not supported, since it is mostly responsibility of the client.

SMAC learns new words on the go, via the Learn() function, can UnLearn() them and can also UnLearn() words provided in the bootstrap dictionary, and keeps into account the frequency of acceptance of words (frequently used words) by giving them priority when generating completion lists.

If customizations of the autocompleter (learnt words etc.) are to survive re-instantiations, or they must be transferred somewhere else, SMAC can Save() what it has learnt to file and can subsequently Retrieve() them from file.

SMAC has not been tested on logographic alphabets, for which it would not make much sense, unless 
an ortographic spelling is provided (hiragana to kanji for example). If there is such a need SMAC 
could be adapted to it.

## Using SMAC

You would usually find a dictionary file and bootstrap SMAC from it:
```Go
ac, err := NewAutoCompleteLinoF(wordFile, 4, 10, 90)
	if err != nil {
		os.Exit(-1)
	}
```
The parameters to the constructor make sense for an average use (more detail below).

After getting an instance of SMAC, you can go straight to using it:
```Go
ac, _ := autoComplete.Complete("chair")
fmt.Println(ac)
```
To make SMAC smarter, make sure to Accept() every word that is selected after autocompletion:
```Go
err := autoComplete.Learn("chairman")
if err != nil {
  // do something
 }
```
Accept() will bounce an error if a word is not in the dictionary.

To make SMAC learn a new word, use Learn():
```Go
err := autoComplete.Learn("Pneumonoultramicroscopicsilicovolcanoconiosis")
if err != nil {
  // do something
 }
```
Learn() will bounce an error if the word to learn is already in the dictionary

To make SMAC forget a word, use UnLearn():
```Go
err := autoComplete.UnLearn("Pneumonoultramicroscopicsilicovolcanoconiosis")
if err != nil {
  // do something
 }
 ```
 UnLearn() will bounce an error if the word to forget was not learnt in the first place.
 
 To save to file whatever SMAC has learnt, use Save():
 ```Go
 err := autoComplete.Save("/home/....")
 if err != nil {
  // do something
 }
 ```
 Save() will only save a diff from a bootstrap dictionary. Save() will bounce an error if it cannot write to file.
 To retrieve a saved SMAC (possibly after using a bootstrap dictionary) use Retrieve():
  ```Go
 err := autoComplete.Retrieve("/home/....")
 if err != nil {
  // do something
 }
 ```
 Retrieve() will only retrieve a diff from a bootstrap dictionary. Retrieve() will bounce an error if it cannot read from file.
 ### Other constructors, finetuning
 You can also bootstrap from an array of strings:
```Go
ac, err := NewAutoCompleteLinoS([]string{"mary", "had", "a", "little", "lamb"}, 4, 10, 90)
	if err != nil {
		os.Exit(-1)
	}
```
 or just start from scratch and learn as you go:
```Go
ac, err := NewAutoCompleteLinoE(4, 10, 90)
	if err != nil {
		os.Exit(-1)
	}
```
 **Meaning of the constructor parameters**
 
 The first parameter is the prefixMapDepth. It speeds up autocompletion, at the expense of memory usage. Practical values are 1, 2, 3 and 4.
 Speed and memory usage are (for a list of 355k words):
 
 prefixMapDepth|speed| memory usage
 --------------|-----|-------------
 1|630k ns/completion (1500 completions/sec)| 46 MB
 2|160k ns/completion (6000 completions/sec)| 50 MB
 3|30k ns/completion (32k completions/sec)| 51 MB
 4|9k ns/completion (60k completions/sec)| 60 MB
 
 The second parameter is the result size. It means how many words you get for an autocompletion. So for example, a result size of 10 would yield these results for "chair":
 
 chair
 chairborne
 chaired
 chairer
 chairing

while a result size of 10 would yield:

chair
chairborne
chaired
chairer
chairing
chairladies
chairlady
chairless
chairlift
chairmaker

The third parameter is the radius. It indicates how deep SMAC will "fish" for frequently used words (marked with Accept() ). Lets say that i frequently use the word "chairmaker". If my result size is 5, and my radius is also 5, I will never see "chairmaker" when I type "chair". With a radius of 20, SMAC will go beyond the 5th result, find out that "chairmaker" is frequently used and put it in front of the list.
### Implementation details
TODO
