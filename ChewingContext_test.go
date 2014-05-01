/*
 * Copyright (c) 2014 ChangZhuo Chen <czchen@gmail.com>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 */
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
