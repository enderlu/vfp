//https://bitbucket.org/mikespook/golib/src/27c65cdf8a77/funcmap/funcmap.go?at=default
package main

import (
	"fmt"
	"reflect"
)

func hello(z string, x string) string {
	return fmt.Sprintln("hello", z, x)
}
func main() {
	actions := make(map[string]reflect.Value)
	actions["hello"] = reflect.ValueOf(hello)
	zz := Call(actions["hello"], "lux", "tt")
	zv, ok := zz[0].Interface().(string)
	fmt.Println(zz, zv, ok)
}
func Call(f reflect.Value, params ...interface{}) []reflect.Value {
	in := GetParas("lux", "lex")
	return f.Call(in)
}
func GetParas(params ...interface{}) (in []reflect.Value) {
	in = make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	return
}
