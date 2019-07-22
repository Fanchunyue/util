// Package util 自用常用函数包
package util

// ErrorStruct 封装后的错误结构体
type ErrorStruct struct {
	Errcode int
	Errmsg  string
}

// RetErr 创建一个错误
func RetErr(code int, message string) *ErrorStruct {
	err := new(ErrorStruct)
	err.Errcode = code
	err.Errmsg = message
	return err
}

var ErrAppIDZero = RetErr(400, "APP ID 不能为空")
var ErrUIDZero = RetErr(401, "UID 不能为空")
var ErrWIDZero = RetErr(402, "WID 不能为空")
