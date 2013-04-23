package vfp

import "reflect"

type Collection map[interface{}]interface{}

func (b Collection) Add(zitem interface{}, zkey string) error {
	b[zkey] = zitem
	return nil
}

func (b Collection) Remove(zkey string) error {
	delete(b, zkey)
	return nil
}

func (b Collection) Clear() error {
	for zi, _ := range b {
		delete(b, zi)
	}

	return nil
}

//Places the names of properties, procedures, and member objects for an object into a variable array.
//func Amembers(zobj interface{}) []string {
//	return
//}

//Check if it is Method
func Ismethod(zobj interface{}, zname string) bool {
	_, ok := reflect.ValueOf(zobj).Type().MethodByName(zname)
	return ok
}

//Check if it is Property
func Isproperty(zobj interface{}, zname string) bool {
	_, ok := reflect.ValueOf(zobj).Type().FieldByName(zname)
	return ok
}
