package utils

import (
	"io/ioutil"
	"log"
	"os"
)

func GetPwd() string {
	pwd, _ := os.Getwd()
	return pwd
}

// PathExists 判断所给路径文件/文件夹是否存在

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// GetFileSize 获取文件大小
func GetFileSize(filepath string) int64 {
	file, err := os.Stat(filepath)
	if err != nil {
		return 0
	}
	return file.Size()
}

// ListDirByType 查看指定目录下所有文件
//
// resultType: 1 文件  2 目录 其他全部
func ListDirByType(path string, resultType uint) (result []string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}
	checked := true
	for _, f := range files {
		switch resultType {
		case 1:
			checked = !f.IsDir()
		case 2:
			checked = f.IsDir()
		default:
			checked = true
		}
		if checked {
			result = append(result, f.Name())
		}
	}
	return
}

// Mkdir 创建目录
func Mkdir(path string) bool {
	if PathExists(path) {
		return true
	}
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Printf("创建文件夹发生错误：%v", err)
		return false
	}
	return true
}

// IsDir 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

// RemoveFile 删除文件
func RemoveFile(filepath string) error {
	return os.Remove(filepath)
}

func RenameFile(oldPath, newPath string) error {
	return os.Rename(oldPath, newPath)
}
