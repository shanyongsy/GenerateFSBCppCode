package main

import (
	"fmt"
	"os"

	generate "generate-code/generate-detail"
)

func main() {
	// 检查输入参数
	if len(os.Args) < 3 {
		fmt.Println("Usage: generate-detail <type> <name>")
		return
	}

	// 输入参数1：生成类型
	generateType := os.Args[1]

	// 输入参数2：名称
	name := os.Args[2]

	// 根据生成类型生成代码
	code := ""
	switch generateType {
	case "msg":
		// 生成消息代码
		code = generate.GenerateMsgCode(name)
	case "db":
		// 生成数据库代码
		code = generate.GenerateDBCode(name)
	case "ui":
		// 生成界面代码
		code = generate.GenerateUICode(name)
	default:
		fmt.Println("Unknown generate type: ", generateType)
		return
	}

	// 判断 ./product/ 目录是否存在, 不存在则创建
	productPath := "./product/"
	_, err := os.Stat(productPath)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(productPath, os.ModePerm)
		}
	}

	// 输出生成的代码到文件, 文件位置为 ./product/
	fileName := productPath + generateType + "_" + name + ".h"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Create file error: ", err)
		return
	}
	defer file.Close()

	file.WriteString(code)
	fmt.Println("Generate code to file: ", fileName)
}
