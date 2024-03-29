package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: StructToMap
//@description: 利用反射将结构体转化为map
//@param: obj interface{}
//@return: map[string]interface{}

func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	data := make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		if obj1.Field(i).Tag.Get("mapstructure") != "" {
			data[obj1.Field(i).Tag.Get("mapstructure")] = obj2.Field(i).Interface()
		} else {
			data[obj1.Field(i).Name] = obj2.Field(i).Interface()
		}
	}
	return data
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ArrayToString
//@description: 将数组格式化为字符串
//@param: array []interface{}
//@return: string

func ArrayToString(array []interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}
func IntArrayToString(array []int) (result string) {
	for _, i := range array {
		result += strconv.Itoa(i)
	}
	return
}
func StringArrayToString(arr []string) (result string) {
	for _, v := range arr {
		result += v
	}
	return
}
func BlankStringArrayToString(arr []string) (result string) {
	for k, v := range arr {
		if k == len(arr)-1 {
			result += v
		} else {
			result += v
			result += ","
		}
	}
	return
}

//字符串fenge
func StringToStringArray(strArr string, sep string) []string {
	result := strings.Split(strArr, sep)
	return result
}
