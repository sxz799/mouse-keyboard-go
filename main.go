package main

import (
	"GSAutoHSProject/record"
	"GSAutoHSProject/replay"
	"GSAutoHSProject/utils"
	"fmt"
	"os"
)

func main() {
	utils.Init() //初始化(识别设备分辨率)
	for {
		var id string
		fmt.Println("欢迎使用MouseKeyboardGo")
		fmt.Println("MouseKeyboardGo可以录制鼠标操作(移动、点击、拖动)和键盘操作")
		fmt.Println("===========================")
		fmt.Println("==   0 -> 退出程序       ==")
		fmt.Println("==   1 -> 录制脚本       ==")
		fmt.Println("==   2 -> 执行脚本       ==")
		fmt.Println("===========================")
		fmt.Println("请输入操作序号:")
		fmt.Scan(&id)

		switch id {
		case "1":
			record.DoRecord()
		case "2":
			replay.DoReplay()
		case "0":
			os.Exit(1)
		}
	}
}
