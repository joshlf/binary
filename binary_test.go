// Copyright 2013 The Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package binary

import (
	"fmt"
	"math"
	"testing"
)

func stringRev(str string) string {
	n := 0
	rune := make([]rune, len(str))
	for _, r := range str {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	// Convert back to UTF-8.
	return string(rune)
}

func TestReverseBits8(t *testing.T) {
	for i := uint8(0); i < math.MaxUint8; i++ {
		rev := ReverseBits8(i)
		str := fmt.Sprintf("%08b", i)
		strRev := stringRev(fmt.Sprintf("%08b", rev))
		if str != strRev {
			t.Errorf("(x, reverse(x)) = (%08b, %08b)", i, rev)
		}
	}
	i := uint8(math.MaxUint8)
	rev := ReverseBits8(i)
	str := fmt.Sprintf("%08b", i)
	strRev := stringRev(fmt.Sprintf("%08b", rev))
	if str != strRev {
		t.Errorf("(x, reverse(x)) = (%08b, %08b)", i, rev)
	}
}

func TestReverseBits16(t *testing.T) {
	for i := uint16(0); i < math.MaxUint16; i++ {
		rev := ReverseBits16(i)
		str := fmt.Sprintf("%016b", i)
		strRev := stringRev(fmt.Sprintf("%016b", rev))
		if str != strRev {
			t.Errorf("(x, reverse(x)) = (%016b, %016b)", i, rev)
		}
	}
	i := uint16(math.MaxInt16)
	rev := ReverseBits16(i)
	str := fmt.Sprintf("%016b", i)
	strRev := stringRev(fmt.Sprintf("%016b", rev))
	if str != strRev {
		t.Errorf("(x, reverse(x)) = (%016b, %016b)", i, rev)
	}
}

func TestReverseBits32(t *testing.T) {
	diff := uint32(1)
	// Too expensive to exhaustively test the space; thus,
	// test it by incrementing the value by which to increment i
	//
	// To detect wrappingn around, see if the next index
	// will be less than the current index
	for i := uint32(0); i+diff > i; i, diff = i+diff, diff+1 {
		rev := ReverseBits32(i)
		str := fmt.Sprintf("%032b", i)
		strRev := stringRev(fmt.Sprintf("%032b", rev))
		if str != strRev {
			t.Errorf("(x, reverse(x)) = (%032b, %032b)", i, rev)
		}
	}
	i := uint32(math.MaxInt32)
	rev := ReverseBits32(i)
	str := fmt.Sprintf("%032b", i)
	strRev := stringRev(fmt.Sprintf("%032b", rev))
	if str != strRev {
		t.Errorf("(x, reverse(x)) = (%032b, %032b)", i, rev)
	}
}

func TestReverseBits64(t *testing.T) {
	diff := uint64(1)
	// Too expensive to use the approach taken in the 32-bit
	// test function. Thus, take a slightly exponential
	// approach (with a multiple close to 1)
	//
	// To detect wrappingn around, see if the next index
	// will be less than the current index
	//
	// Note that the multiple is set carefully - changing it
	// may vastly reduce or increase the time it takes to run
	// this test (currently it takes 1.5s on a quad-core i7)
	for i := uint64(0); i+diff > i; i, diff = i+diff, uint64(float64(diff)*1.0001)+1 {
		rev := ReverseBits64(i)
		str := fmt.Sprintf("%064b", i)
		strRev := stringRev(fmt.Sprintf("%064b", rev))
		if str != strRev {
			t.Errorf("(x, reverse(x)) = (%064b, %064b)", i, rev)
		}
	}
	i := uint64(math.MaxInt64)
	rev := ReverseBits64(i)
	str := fmt.Sprintf("%064b", i)
	strRev := stringRev(fmt.Sprintf("%064b", rev))
	if str != strRev {
		t.Errorf("(x, reverse(x)) = (%064b, %064b)", i, rev)
	}
}
