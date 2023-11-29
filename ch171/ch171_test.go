package ch171

import (
	"testing"
)

func TestTitleToNumber(t *testing.T) {
	s := "ZY"
	expect := 701
	ret := TitleToNumber(s)
	if ret != expect {
		t.Errorf("expect %d,return %d", expect, ret)
	}
}
