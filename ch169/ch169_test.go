package ch169_test

import (
	"leetcode/ch169"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	expect := 2
	ret := ch169.MajorityElement(nums)
	if expect != ret {
		t.Errorf("expect %d,return %d", expect, ret)
	}
}
