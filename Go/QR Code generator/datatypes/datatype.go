package datatype

import "regexp"


type DataType string
var regexNumber *regexp.Regexp
var regexAlphaNumeric *regexp.Regexp


const (
	Number DataType = "number"
	AlphaNumeric DataType = "alphanumeric"
	Byte DataType = "byte"
)

func WhatIsDataType(s string) DataType{
 return Number
}