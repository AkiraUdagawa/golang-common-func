package main

import (
	"TestProject/func/command"
)

func main() {
	//コンソールに表示
	command.Echo("hoge")

	//MD5 Hashを生成
	var md5text = command.Md5("hoge")
	command.Echo(md5text)
}
