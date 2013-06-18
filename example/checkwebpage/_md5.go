package main

import "fmt"
import "bufio"
import "crypto/md5"
import "io/ioutil"
import "strings"
import "os"

func main() {
	fmt.Println(Md5("lu30952128"))
	fmt.Print("input password:")
	zx := Wait()
	Strtofile(strings.ToUpper(Md5(zx)), "pass.txt")
}

func Wait() string {
	reader := bufio.NewReader(os.Stdin)
	zr, _, _ := reader.ReadLine()
	zline := string(zr)
	return zline
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

func Strtofile(zstrArg interface{}, zfile string) error {
	//file, err := os.Create(zfile)
	//if err != nil {
	//	return err
	//}
	//defer file.Close()
	//file.WriteString(zstr)
	//return nil
	var zstr []byte
	switch zstrArg.(type) {
	case string:
		zstr = []byte(zstrArg.(string))
	case []byte:
		zstr = zstrArg.([]byte)
	default:
		zstr = []byte(fmt.Sprintf("%v", zstrArg))
	}
	return ioutil.WriteFile(zfile, zstr, os.ModeAppend)
}
