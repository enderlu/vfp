package vfp

import (
	
	"testing"
)

func TestRun(t *testing.T) {
	Run(`cmd.exe`)

}

func TestSys(t *testing.T) {
	Sys(2)
	Sys(16)
	Sys(2003)

}
