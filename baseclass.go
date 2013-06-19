package vfp

import "reflect"

//func CreateObject()  {

//}

//Places the names of properties, procedures, and member objects for an object into a variable array.
/*Example:
t := new(ShortDate)
fmt.Println(Amembers(*t))
*/
func Amembers(zobj interface{}) (zr [][]string) {

	t := reflect.TypeOf(zobj)
	for i := 0; i < t.NumField(); i++ {
		zr = append(zr, []string{t.Field(i).Name, "Field"})
	}
	for i := 0; i < t.NumMethod(); i++ {
		zr = append(zr, []string{t.Method(i).Name, "Method"})
	}
	return
}

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

func Vartype(zobj interface{}) string {
	return reflect.TypeOf(zobj).String()
}

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
