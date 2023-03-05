package sliceSet

import (
	"fmt"
	"reflect"
)

func InterfaceToMap(a interface{}) []map[string]interface{} {
	aas := make([]map[string]interface{}, 0)
	as := reflect.ValueOf(a)
	if as.Kind() != reflect.Slice {
		fmt.Println("不是切片")
		return nil
	}
	n := as.Len()
	for i := 0; i < n; i++ {
		val := as.Index(i)
		t := val.Type()
		if t.Kind() != reflect.Struct {
			fmt.Println("不是结构体")
			return nil
		}
		num := val.NumField()
		m := make(map[string]interface{})
		for j := 0; j < num; j++ {
			m[t.Field(j).Name] = val.Field(j).Interface()
		}
		aas = append(aas, m)
	}
	return aas
}

func Mines(a, b interface{}) []map[string]interface{} {
	aas := InterfaceToMap(a)
	bbs := InterfaceToMap(b)
	data := make(map[string]bool)
	for _, bs := range bbs {
		for _, val := range bs {
			b_t := reflect.ValueOf(val)
			switch b_t.Kind() {
			case reflect.String:
				data[val.(string)] = true
			}
		}
	}
	var j int
	for _, as := range aas {
		for _, val := range as {
			a_t := reflect.ValueOf(val)
			switch a_t.Kind() {
			case reflect.String:
				if _, ok := data[val.(string)]; !ok {
					aas[j] = as
					j++
				}
			}
		}
	}
	return aas[:j]
}
