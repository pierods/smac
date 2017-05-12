package smac

import (
	"reflect"
	"sort"
	"testing"
)

func Test_LinoConstruction(t *testing.T) {

	t.Log("Given the need to test the prefix map creation")
	{
		words := []string{"aaaa", "aabc", "abc", "bbb", "v", "vvv", "vvvaaa"}
		prefix0 := make(map[string]string)
		prefix1 := make(map[string]string)
		prefix2 := make(map[string]string)
		prefix3 := make(map[string]string)

		prefix1["a"] = "aaaa"
		prefix1["b"] = "bbb"
		prefix1["v"] = "v"

		for k, v := range prefix1 {
			prefix2[k] = v
		}

		prefix2["aa"] = "aaaa"
		prefix2["ab"] = "abc"
		prefix2["bb"] = "bbb"
		prefix2["vv"] = "vvv"

		for k, v := range prefix2 {
			prefix3[k] = v
		}

		prefix3["aaa"] = "aaaa"
		prefix3["aab"] = "aabc"
		prefix3["abc"] = "abc"
		prefix3["bbb"] = "bbb"
		prefix3["vvv"] = "vvv"

		sort.Strings(words)

		prefixes := makePrefixMap(words, 0)
		if !reflect.DeepEqual(prefixes, prefix0) {
			t.Log("Should be able to create a 0 len prefix map", ballotX)
		}
		t.Log("Should be able to create a 0 len prefix map", checkMark)

		prefixes = makePrefixMap(words, 1)

		if !reflect.DeepEqual(prefixes, prefix1) {
			t.Fatal("Should be able to create a 1 len prefix map", ballotX)
		}
		t.Log("Should be able to create a 1 len prefix map", checkMark)

		prefixes = makePrefixMap(words, 2)

		if !reflect.DeepEqual(prefixes, prefix2) {
			t.Fatal("Should be able to create a 2 len prefix map", ballotX)
		}
		t.Log("Should be able to create a 2 len prefix map", checkMark)

		prefixes = makePrefixMap(words, 3)

		if !reflect.DeepEqual(prefixes, prefix3) {
			t.Fatal("Should be able to create a 3 len prefix map", ballotX)
		}
		t.Log("Should be able to create a 3 len prefix map", checkMark)

	}

	t.Log("Given the need to test the lino creation")
	{

		words := []string{"aaaa", "aabc", "abc", "bbb", "v", "vvv", "vvvaaa"}
	}
}
