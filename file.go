package vfp

import "os"

import "os/exec"
import (
	"io/ioutil"
	"path/filepath"
)

//Places information about files into an array and then returns the number of files.
func Adir(zdir string) ([]os.FileInfo, error) {

	zpath := Justpath(zdir)
	zfile := Strtran(Justfname(zdir), "*", "", -1)
	zfile = Lower(Strtran(zfile, "?", "", -1))

	list, err := ioutil.ReadDir(zpath)

	zl := make([]os.FileInfo, 0)

	for _, zf := range list {
		if At(zfile, Lower(zf.Name())) > 0 {
			zl = append(zl, zf)
		}
	}

	return zl, err
}

func Strtofile(zstr, zfile string) error {
	//file, err := os.Create(zfile)
	//if err != nil {
	//	return err
	//}
	//defer file.Close()
	//file.WriteString(zstr)
	//return nil

	return ioutil.WriteFile(zfile, []byte(zstr), os.ModeAppend)
}

//Returns the contents of a file as a []byte.
func Filetostr(zfile string) (zstr []byte, err error) {
	return ioutil.ReadFile(zfile)
	//file, err := os.Open(zfile)
	//if err != nil {
	//	return
	//}

	//defer file.Close()
	//zbr := bufio.NewReader(file)
	//zstr, err = ioutil.ReadAll(zbr)

	//return

}

//Returns the characters of a file extension from a complete path.
func Justext(zpath string) string {
	return filepath.Ext(zpath)
}

//Returns the path portion of a complete path and file name.
//Example:
//	vfp.Justpath("c:/sek/ww.dat")
func Justpath(zpath string) string {
	return filepath.Dir(zpath)
}

//Returns the file name portion of a complete path and file name.
/*Example:
print ww.dat

fmt.Println("justfname:", vfp.Justfname("c:\\sek\\ww.dat"))	
*/
func Justfname(zpath string) string {
	_, zv := filepath.Split(zpath)
	return zv
}

//Returns the stem name (the file name before the extension) from a complete path and file name.
/*Example:
print xx

	vfp.Juststem("c:\\ccc\\xx.dat")

*/
func Juststem(zpath string) string {
	zv := Justfname(zpath)
	for zi := 0; zi < len(zv); zi++ {
		if zv[zi] == '.' {
			return zv[:zi]
		}
	}
	return zv
}

//Adds a backslash (if needed) to a path expression.
/*Example:
Both print C:\Windows\

vfp.Addbs( "C:\\Windows" )

vfp.Addbs( "C:\\Windows\\" )

*/
func Addbs(zpath string) (zret string) {

	if !os.IsPathSeparator(zpath[len(zpath)-1]) {
		return zpath + string(os.PathSeparator)
	}
	return zpath
}

func sourceExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

//Locates the specified file.
func File(zfile string) bool {
	return sourceExist(zfile)
}

//Locates the specified directory. 
func Directory(zdir string) bool {
	return sourceExist(zdir)
}

//Returns the current directory.
func Curdir() string {
	zpwd, _ := os.Getwd()
	return zpwd
}

//Returns the path to a specified file or the path relative to another file.
func Fullpath(zfile string) string {
	zp, zerr := exec.LookPath(zfile)
	if zerr != nil {
		return ""
	}
	//os.Args[0] 当前执行程序的文件路径
	return zp
}
