package util

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var ps string

func init() {
	if os.IsPathSeparator('\\') { //前边的判断是否是系统的分隔符
		ps = "\\"
	} else {
		ps = "/"
	}
}

// ReadeFile 读取文件
func ReadeFile(path string) ([]byte, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer file.Close()
	template, err := ioutil.ReadAll(file)
	return template, nil
}

// MkdirAll 创建目录 有BUG  不一定都是存在当前目录下的
func MkdirAll(dirs ...string) error {
	path, _ := os.Getwd()
	l := len(dirs)
	if l == 0 {
		return nil
	}

	for i := 0; i < len(dirs); i++ {
		path = path + ps + dirs[i]
	}
	println(path)
	err := os.MkdirAll(path, os.ModePerm) //在当前目录下生成AuthFormFile目录
	return err
}

// Create 创建文件 自动处理不同操作系统的文件路径分隔符
// !!! 创建的文件没有 close 请手动关闭
func Create(filePath ...string) (*os.File, error) {
	l := len(filePath)
	if l == 0 {
		return nil, nil
	}
	fp, _ := os.Getwd()
	for i := 0; i < l; i++ {
		fp = fp + ps + filePath[i]
	}
	file, err := os.Create(fp)
	// defer file.Close()
	return file, err
}

// WalkDir 获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。
func WalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	// passFile := strings.ToUpper("DS_Store")
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		//if err != nil { //忽略错误
		// return err
		//}

		if fi.IsDir() { // 忽略目录
			return nil
		}

		if !strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}

		return nil
	})

	return files, err
}

// 获取路径文件的 SHA256 值
func HashSHA256File(filePath string) (string, error) {
	var hashValue string
	file, err := os.Open(filePath)
	if err != nil {
		return hashValue, err
	}
	defer file.Close()
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return hashValue, err
	}
	hashInBytes := hash.Sum(nil)
	hashValue = hex.EncodeToString(hashInBytes)
	return hashValue, nil

}

// CopyFile 复制文件 将文件 fromFile 拷贝到 toFile
func CopyFile(toFile, fromFile string) (written int64, err error) {
	src, err := os.Open(fromFile)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(toFile, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}
