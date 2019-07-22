// Package util 自用常用函数包
package util

import (
	"log"
	"math/rand"
	"time"
)

// BaseForm base表单结构体 只有一个 ID form:"id"
type BaseForm struct {
	ID int64 `form:"id"`
}

// 获取随机字符串的类型常亮
const (
	RandKindNum   = iota // RandKindNUM 纯数字
	RandKindLower        // 小写字母
	RandKindUpper        // 大写字母
	RandKindAll          // 数字、大小写字母
)

// Rand 获取随机字符串
func Rand(size int, kind int) string {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	isAll := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}

	return string(result)
}

// GetNowTime 获取当前时间
func GetNowTime() int64 {
	return time.Now().Unix()
}

// GetNowUtcTime 获取 UTC 时间
func GetNowUtcTime() int64 {
	return time.Now().UTC().Unix()
}

// RemoveDuplicatesAndEmpty 排序文本数组并删除空项
func RemoveDuplicatesAndEmpty(a []string) (ret []string) {
	i := 0
	j := 0
	for i = 0; i < len(a); i++ {
		if i == 0 {
			ret = append(ret, a[0])

		}
		has := false
		s := ""
		for j = 0; j < len(ret); j++ {
			log.Println(ret)
			if ret[j] == a[i] || a[i] == "" {
				has = true

			} else {
				if a[i] != "" {
					s = a[i]
				}
			}
		}
		if !has {
			ret = append(ret, s)
		}
	}
	return
}

// ErrToErrs 检查错误
// 该方法判断了 err 为 nil 的情况
func ErrToErrs(err error, errCode ...int) *ErrorStruct {
	if err != nil {
		code := 0
		for i := 0; i < len(errCode); i++ {
			code = errCode[i]
		}
		if code == 0 {
			code = 400
		}

		return RetErr(code, err.Error())
	}
	return nil

}

// ErrNotNil 判断 err 是否为空
// 返回bool
// 不为空则先行打印
func ErrNotNil(err error) bool {
	if err != nil {
		log.Println(err.Error())
		return true
	}
	return false
}
