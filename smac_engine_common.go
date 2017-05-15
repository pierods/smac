// Copyright Piero de Salvia.
// All Rights Reserved

package smac

type wordAccepts struct {
	Word    string
	Accepts int
}
type wordHit struct {
	word    string
	accepts int
	next    *wordHit
}

type sOLILI struct {
	start *wordHit
	end   *wordHit
}

func (list *sOLILI) insert(word string, accepts int) {
	hit := &wordHit{
		word:    word,
		accepts: accepts,
	}
	if list.start == nil {
		list.start = hit
		list.end = hit
		return
	}

	if accepts == 0 {
		list.end.next = hit
		list.end = hit
		return
	}

	if hit.accepts > list.start.accepts {
		hit.next = list.start
		list.start = hit
		return
	}
	cursor := list.start
	for cursor.next != nil {
		if hit.accepts > cursor.next.accepts {
			break
		}
		cursor = cursor.next
	}
	hit.next = cursor.next
	cursor.next = hit
}

func (list *sOLILI) flush() []string {

	slice := []string{}

	if list.start != nil {
		cursor := list.start
		for cursor != nil {
			slice = append(slice, cursor.word)
			cursor = cursor.next
		}
	}
	return slice
}

func (list *sOLILI) flushL(limit int) []string {

	slice := []string{}

	if list.start != nil {
		cursor := list.start
		for i := 0; cursor != nil && i < limit; i++ {
			slice = append(slice, cursor.word)
			cursor = cursor.next
		}
	}
	return slice
}
