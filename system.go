package vfp

import "os/exec"

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
