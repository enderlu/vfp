package vfp

import "os/exec"
import "os"
import "crypto/md5"
import "fmt"
import "github.com/axgle/service"

//Run
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
func Sys(ztype int, zoptArgs ...interface{}) (zret interface{}) {
	switch ztype {
	case 2: //Seconds since midnight.
		zret = Seconds()
	case 16: //Executing program file name.
		zret = os.Args[0]
	case 2003: //Current directory.
		zret = Curdir()
	case 2340: //enable ,disable service
		//zopt := "install"
		//if len(zoptArgs) > 0 {
		//	zopt = Lower(zoptArgs[0].(string))
		//}
		//zname := Juststem(Justfname(Program()))
		////if os is window
		//if zopt == "1" {
		//	Run("net start " + zname)
		//} else {
		//	Run("net stop" + zname)
		//}
	default:
		zret = nil
	}
	return
}

func XServiceAdd(zname, zdisplyName, zdesc string, onStart, onStop func() error) error {

	var ws, err = service.NewService(zname, zdisplyName, zdesc)

	err = ws.Install()
	if err != nil {
		return err
	}
	err = ws.Run(onStart, onStop)
	if err != nil {
		return err
	}
	return nil
}
func XServiceRemove(zname string) error {
	var ws, err = service.NewService(zname, "", "")
	err = ws.Remove()
	if err != nil {
		return err
	}
	return nil
}

func XServiceStart(zname string) error {
	return Run("net start "+zname, true)
}
func XServiceStop(zname string) error {
	return Run("net stop "+zname, true)
}

//Returns the name of the program at a specified program level,
//the name of the currently executing program, the current program level, or the name of the program executing when an error occurred.
func Program() string {
	return Sys(16).(string)
}

func Md5(zstrArg interface{}) string {
	var zstr []byte
	switch zstrArg.(type) {
	case string:
		zstr = []byte(zstrArg.(string))
	case []byte:
		zstr = zstrArg.([]byte)
	default:
		zstr = []byte(fmt.Sprintf("%v", zstrArg))
	}
	zs := md5.New()
	zs.Write([]byte(zstr))
	return fmt.Sprintf("%x", zs.Sum(nil))
}

//Returns the name and version number of the operating system
func OS() string {
	return ""
}
