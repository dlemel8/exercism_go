package flatten

import (
    "reflect"
)

func Flatten(slice interface{}) []interface{} {
    res := []interface{}{}
    flatten(slice, &res)
    return res
}

func flatten(current interface{}, result *[]interface{}) {
    val := reflect.ValueOf(current)
    if val.Kind() == reflect.Int {
        *result = append(*result, val.Interface())
    }
    if val.Kind() == reflect.Slice {
        for i := 0; i < val.Len(); i++ {
            flatten(val.Index(i).Interface(), result)
        }
    }
}
