package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Use go log package
// Initialize a local log
func InitLogConf() error {
	y, m, d := time.Now().Date()
	fileName := fmt.Sprintf("%d-%d-%d_award.log", y, m, d)									// 返回一个字符串
	f , err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)		// 打开这个文件
	if err != nil {
		fmt.Errorf("file open error, %v", err)									// 如果这个错误不为空，打印
		return err
	}

	// display line numbers
	log.SetFlags(log.Lshortfile | log.LstdFlags)									// 设置log的输出选项
	log.SetOutput(f)

	return nil
}
