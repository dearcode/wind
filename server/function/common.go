package function

import (
	"reflect"
	"strings"
)

//TableName 解析table.id中的table.
func TableName(n interface{}) interface{} {
	name := reflect.ValueOf(n).String()
	return reflect.ValueOf(strings.Split(name, ".")[0]).Convert(reflect.TypeOf(n)).Interface()
}

//TableNameToLower 转小写.
func TableNameToLower(n interface{}) interface{} {
	name := reflect.ValueOf(n).String()
	return reflect.ValueOf(strings.ToLower(strings.Split(name, ".")[0])).Convert(reflect.TypeOf(n)).Interface()
}

//TableColumn 解析table.id到tableid.
func TableColumn(n interface{}) interface{} {
	name := reflect.ValueOf(n).String()
	return reflect.ValueOf(strings.Replace(name, ".", "", 1)).Convert(reflect.TypeOf(n)).Interface()
}
