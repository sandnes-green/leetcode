package ch190

import (
	"testing"
)

func TestReverseBits(t *testing.T) {
	var expect uint32 = 964176192
	var num uint32 = 43261596
	ret := reverseBits(uint32(num))
	if expect != ret {
		t.Errorf("expect %d, return %d", expect, ret)
	}
}
