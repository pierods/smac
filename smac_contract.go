// Copyright Piero de Salvia.
// All Rights Reserved

// SMAC is a small autocomplete engine with an emphasis on simplicity and performance.
package smac

// The default result size (number of hits for a given stem) and radius (max lenght of words the autocompleter will
// descend to while searching)
const (
	DEF_RESULTS_SIZE = 10
	DEF_RADIUS       = 15
)

type AutoComplete interface {
	// Accept will consider a word as accepted, i.e. as a completion that has been picked for completion. Every time Accept is called
	// for a given word, its accept count will increase, increasing the priority it will be given when the completion list is created by
	// complete (it will appear before other possible completions). Accept should be called every time a word is accepted for completion,
	// so to make its autocompleter smarter.
	Accept(acceptedWord string) error
	// Learn will add a word to an autocompleter. If the word is also accepted for completion, Accept should also be called
	// on the same word
	Learn(word string) error
	// UnLearn will remove a word from an autocompleter.
	UnLearn(word string) error
	// Complete returns a slice of words from a stem word. The max number of words returned depends on the value of the resultSize parameter used when
	// constructing autoComplete, and the max length of matches depends on the value of the radius parameter used.
	// Matches are returned by default in order of length first and alpabetical second. The exceptions are words that were previously accepted as completions
	// (frequently used words) which bubble up to the top of the list, in order of frequency first and alphabetical second.
	Complete(word string) ([]string, error)
	// Save will save to file everything an autocompleter has learnt, which is, new words, removed words and word accepts.
	// It is up to the client to decide when to call Save (possibly just before shutdown).
	Save(fileName string) error
	// Retrieve will re-teach an autocompleter that has just been created all the learnt words, deleted words and accepted words.
	// It is up to the client to decide when to call Retrieve (possibly just after initialization)
	Retrieve(fileName string) error
}
