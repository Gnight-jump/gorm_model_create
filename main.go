package main

import (
	"modelgenerator/conf"
	"modelgenerator/generate"
)

func main() {

	/**
	使用方法：修改配置文件，参考conf.yaml文件注释修改
	*/

	// 读取初始化配置文件
	conf.Init()
	// 生成目标表单
	generate.Genertate(conf.Tconf)
}
