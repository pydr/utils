package utils

import "reflect"

// 判断obj是否在target中，target支持的类型array,slice,map
func Contains(collection interface{}, that interface{}) bool {
	targetValue := reflect.ValueOf(collection)
	switch reflect.TypeOf(collection).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == that {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(that)).IsValid() {
			return true
		}
	}
	return false
}
