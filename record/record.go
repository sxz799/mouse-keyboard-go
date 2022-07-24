package record

import (
	"GSAutoHSProject/model"
	"GSAutoHSProject/utils"
	"encoding/json"
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"io/ioutil"
	"time"
)

func DoRecord() {
	fmt.Println("开始录制=======按下esc表示录制结束")
	var steps []model.Operation
	lastTime := time.Now()
	var lastOperation model.Operation

	robotgo.EventHook(hook.MouseMove, []string{}, func(event hook.Event) {

		var operation model.Operation
		operation.X = int(float64(event.X) / utils.Dpi)
		operation.Y = int(float64(event.Y) / utils.Dpi)
		operation.Type = "mouseMove"
		nowTime := time.Now()
		operation.WaitTime = nowTime.Sub(lastTime)
		lastTime = time.Now()
		lastOperation = operation
		if lastOperation.Type != "mouseMove" || lastOperation.WaitTime > time.Millisecond*300 {
			steps = append(steps, operation)
		}

	})

	robotgo.EventHook(hook.MouseDown, []string{}, func(event hook.Event) {
		var operation model.Operation
		if event.Button == 1 {
			operation.MouseType = "left"
		} else {
			operation.MouseType = "right"
		}
		operation.X = int(float64(event.X) / utils.Dpi)
		operation.Y = int(float64(event.Y) / utils.Dpi)
		operation.Type = "mouseClick"
		nowTime := time.Now()
		operation.WaitTime = nowTime.Sub(lastTime)
		lastTime = time.Now()
		steps = append(steps, operation)
		lastOperation = operation
	})
	robotgo.EventHook(hook.KeyDown, []string{"esc"}, func(event hook.Event) {
		robotgo.EventEnd()
	})
	robotgo.EventHook(hook.KeyDown, []string{}, func(event hook.Event) {
		var operation model.Operation
		operation.Type = "keyboard"
		operation.InputMsg = string(event.Keychar)
		nowTime := time.Now()
		operation.WaitTime = nowTime.Sub(lastTime)
		lastTime = time.Now()
		steps = append(steps, operation)
		lastOperation = operation
	})

	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)
	ok := robotgo.AddEvents("esc")
	if ok {
		marshal, _ := json.Marshal(steps)
		ioutil.WriteFile("./script.txt", marshal, 0666)
		fmt.Println("脚本录制完毕并写入script.txt文件")
		fmt.Println("============================================================")
	}
}
