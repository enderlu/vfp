package vfp

import "os/exec"
import "os"

//export Run
func Run(zfile string, zwait_arg ...bool) (zerr error) {

	zwait := false
	if len(zwait_arg) > 0 {
		zwait = zwait_arg[0]
	}
	zcmd := exec.Command(zfile)

	if zwait {
		return zcmd.Run()
	}

	return zcmd.Start()
}

//Returns system information.
func Sys(ztype int) (zret interface{}) {
	switch ztype {
	case 2: //Seconds since midnight. 
		zret = Seconds()
	case 16: //Executing program file name.
		zret = os.Args[0]
	case 2003: //Current directory.
		zret = Curdir()
	default:
		zret = nil
	}
	return
}
