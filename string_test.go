package vfp

import (
	"testing"
)

func TestAline(t *testing.T) {
	zstr := "we go to school"
	zn := Aline(zstr, " ")
	if len(zn) != 4 {
		t.Errorf("split count is not 4 for string [%v]", zstr)
	}
}
