package util

import (
	// "bytes"
	// "encoding/binary"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	// "strings"
)

// Int64ToString Int64 转 String
func Int64ToString(i int64) string {
	s := strconv.FormatInt(i, 10)
	return s
}

// Int64ToString Int64 转 String
func IntToString(i int) string {
	s := strconv.Itoa(i)
	return s
}

// StringToFloat64 String 转 Float64
func StringToFloat64(s string) float64 {

	i, err := strconv.ParseFloat(s, 64)
	ErrToErrs(err, 400)
	return i
}

// StringToInt64 String 转 Int64
func StringToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	ErrToErrs(err, 400)
	return i
}

// StringToUint String 转 Uint 莫名觉得有问题
func StringToUint(s string) uint {
	i, err := strconv.ParseUint(s, 10, 64)
	ErrToErrs(err, 400)
	return uint(i)
}

// StringToUint64 String 转 Uint64 莫名觉得有问题
func StringToUint64(s string) uint64 {
	i, err := strconv.ParseUint(s, 10, 64)
	ErrToErrs(err, 400)
	return i
}

// Uint64ToString Uint64 转 String
func Uint64ToString(i uint64) string {
	return strconv.FormatUint(i, 10)
}

// StringToMD5 String 转 MD5
func StringToMD5(s string) string {
	data := []byte(s)
	return fmt.Sprintf("%x", md5.Sum(data))
}

// StructToJSON struct 转 json []byte
func StructToJSON(i interface{}) []byte {
	j, err := json.Marshal(i)
	ErrToErrs(err, 400)
	return j
}

// StructToJSONS struct 转 json 文本
func StructToJSONS(i interface{}) string {
	return string(StructToJSON(i))
}

// JSONToStruct Json 转 Struct
func JSONToStruct(j []byte, v interface{}) error {
	data := j
	err := json.Unmarshal(data, v)
	return err
}

// JSONStringToStruct Json String 转 Struct
func JSONStringToStruct(s string, v interface{}) error {

	return JSONToStruct([]byte(s), v)
}

// Round Float64 四舍五入 保留 n 个小数位 算法很迷茫有空再看
func Round(f float64, n int) float64 {
	pow10N := math.Pow10(n)
	return math.Trunc((f+0.5/pow10N)*pow10N) / pow10N
}
