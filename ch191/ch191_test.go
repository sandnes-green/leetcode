package ch191

import "testing"

func TestHammingWeight(t *testing.T) {
	var num uint32 = 64
	var expect int = 1
	ret := HammingWeight(num)
	if ret != expect {
		t.Errorf("expect %d, return %d", expect, ret)
	}
}
