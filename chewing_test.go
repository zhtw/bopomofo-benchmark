package main

import (
	"reflect"
	"testing"
)

func Test_bopomofoToKey(t *testing.T) {
	var keySequence []uint8
	var expected []uint8

	keySequence = bopomofoToKey("ㄘㄜˋㄕˋ")
	expected = []uint8{'h', 'k', '4', 'g', '4'}
	if !reflect.DeepEqual(keySequence, expected) {
		t.Fatalf("`%s' is converted to `%s', shall be `%s'", "ㄘㄜˋㄕˋ", keySequence, expected)
	}

	keySequence = bopomofoToKey("ㄏㄚㄏㄚ")
	expected = []uint8{'c', '8', ' ', 'c', '8', ' '}
	if !reflect.DeepEqual(keySequence, expected) {
		t.Fatalf("`%s' is converted to `%s', shall be `%s'", "ㄏㄚㄏㄚ", keySequence, expected)
	}
}
