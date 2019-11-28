package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// read 负责读取文件
// 这是一个通用的方法
func read(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	return data
}

func write(path, content string) {
	err := ioutil.WriteFile(path, []byte(content), 0755)
	if err != nil {
		log.Fatal(err)
	}
}

// 利用 VSCode 打开文件
func vscodeOpen(filename string) {
	cmd := exec.Command("code", "-r", filename)
	_, err := cmd.Output()
	if err != nil {
		panic(err.Error())
	}
}

func createFile(filePath string) error {
	if !isExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		return err
	}
	return nil
}

// 判断所给路径文件/文件夹是否存在(返回true是存在)
func isExist(path string) bool {
	_, err := os.Stat(path) // os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
