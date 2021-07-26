package main

import (
	"reflect"
	"regexp"
	"strings"
)

func GetStructFields(e interface{}) []string {
	var fields []string
	value := reflect.ValueOf(e).Elem()
	for i := 0; i < value.NumField(); i++ {
		fields = append(fields, value.Type().Field(i).Name)
	}
	return fields
}

func ToSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func ArrayToSnakeCase(a []string) []string {
	var converted []string
	for _, str := range a {
		converted = append(converted, ToSnakeCase(str))
	}
	return converted
}
