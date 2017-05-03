package smac

import (
	"reflect"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func Test_fIFO(t *testing.T) {

	var fifo fIFO

	t.Log("Given the need to test a FIFO")
	{
		t.Log("When initializing a FIFO")
		{
			if fifo.size() != 0 {
				t.Fatal("Should be able to initialize an empty FIFO", ballotX)
			}
			t.Log("Shoulf be able to initialize an empty FIFO", checkMark)
		}

		t.Log("When growing/shrinking a FIFO")
		{
			a := branch{
				node:   nil,
				parent: []rune("a"),
			}

			fifo.add(a)

			if fifo.size() != 1 {
				t.Fatal("Should be able to grow a FIFO by 1", ballotX)
			}
			t.Log("Should be able to grow a FIFO by 1", checkMark)
			newA := fifo.remove()
			if !reflect.DeepEqual(newA, a) {
				t.Fatal("Should be able to retrieve an element from a FIFO", ballotX)
			}
			t.Log("Should be able to retrieve an element from a FIFO", checkMark)
			b := branch{
				node:   nil,
				parent: []rune("b"),
			}
			c := branch{
				node:   nil,
				parent: []rune("c"),
			}
			fifo.add(a)
			fifo.add(b)
			fifo.add(c)
			if fifo.size() != 3 {
				t.Fatal("Should be able to grow a fifo by 3", ballotX)
			}
			t.Log("Should be able to grow a fifo by 3", checkMark)
			el := fifo.remove()
			if !reflect.DeepEqual(el, a) {
				t.Fatal("Should be able to shrink a fifo in order", ballotX)
			}
		}
	}
}
