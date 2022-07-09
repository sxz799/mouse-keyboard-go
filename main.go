package main

import (
	"GSAutoHSProject/record"
	"GSAutoHSProject/replay"
	"GSAutoHSProject/utils"
	"os"

	"fmt"
)

func main() {
	utils.Init() //初始化(识别设备分辨率)
	for {
		var id string
		fmt.Println("欢迎使用MouseKeyboardGo")
		fmt.Println("MouseKeyboardGo可以录制鼠标点击和键盘输出(不支持拖动、不支持鼠标悬停)")
		fmt.Println("============================================================")
		fmt.Println("0.退出")
		fmt.Println("1.录制操作")
		fmt.Println("2.执行录制的脚本")
		fmt.Println("============================================================")
		fmt.Scan(&id)
		switch id {
		case "1":
			record.DoRecord()
		case "2":
			replay.DoReplal()
		case "0":
			os.Exit(1)
		}

	}

}
