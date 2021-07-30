package http

import (
	"errors"
	"io"
	"mime"
	"os"
	"strconv"
	"strings"

	"sextube/utils"
)

func DownloadFile(reqConfig interface{}, savePath string, fb func(length, downLen int64)) (string, error) {
	// 设置超时时间
	// config.Timeout = int(time.Second * 60)
	config := parseToRequestConfig(reqConfig)
	// get方法获取资源
	resp, err := Request(config, nil)
	if err != nil {
		return "", err
	}
	// 读取服务器返回的文件大小
	filesize, err := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 32)
	if err != nil {
		return "", err
	}
	// 获取下载文件名
	_, params, err := mime.ParseMediaType(resp.Header.Get("Content-Disposition"))
	if err != nil {
		return "", err
	}
	// 给路径补 /
	if !strings.HasSuffix(savePath, "/") {
		savePath += "/"
	}
	if !utils.Mkdir(savePath) {
		return "", errors.New("下载目录创建失败！")
	}
	var (
		filename    = params["filename"]
		tmpFilePath = savePath + filename + ".download"
	)
	// 删除已有文件
	_ = utils.RemoveFile(savePath + filename)
	// 创建文件
	file, err := os.Create(tmpFilePath)
	// 清理资源
	defer func() {
		resp.Body.Close()
		file.Close()
	}()
	if err != nil {
		return "", err
	}
	if resp.Body == nil {
		return "", errors.New("body is null")
	}
	// 下面是 io.copyBuffer() 的简化版本
	var (
		written int64
		buf     = make([]byte, 32*1024)
	)
	for {
		// 读取bytes
		nr, er := resp.Body.Read(buf)
		if nr > 0 {
			// 写入bytes
			nw, ew := file.Write(buf[0:nr])
			// 数据长度大于0
			if nw > 0 {
				written += int64(nw)
			}
			// 写入出错
			if ew != nil {
				err = ew
				break
			}
			// 读取是数据长度不等于写入的数据长度
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
		// 没有错误了快使用 callback
		if fb != nil {
			fb(filesize, written)
		}
	}
	if err == nil {
		_ = utils.RenameFile(tmpFilePath, savePath+filename)
	}
	return savePath + filename, err
}
