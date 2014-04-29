package main

import (
	"reflect"
	"testing"
)

func Test_bopomofoToKey(t *testing.T) {
	var keySequence []uint8

	keySequence = bopomofoToKey("ㄘㄜˋㄕˋ")
	expected := []uint8{'h', 'k', '4', 'g', '4'}
	if !reflect.DeepEqual(keySequence, expected) {
		t.Fatal("Cannot convert ㄘㄜˋㄕˋ to keySequence hk4g4")
	}
}
