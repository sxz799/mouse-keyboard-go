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

	handleDrag := func() {

		if lastOperation.Type == "mouseDrag" {
			steps = append(steps, lastOperation)
		}
	}

	robotgo.EventHook(hook.MouseDrag, []string{}, func(event hook.Event) {
		var operation model.Operation
		operation.X = int(float64(event.X) / utils.Dpi)
		operation.Y = int(float64(event.Y) / utils.Dpi)
		operation.Type = "mouseDrag"
		nowTime := time.Now()
		operation.WaitTime = nowTime.Sub(lastTime)
		lastTime = time.Now()
		lastOperation = operation
	})

	robotgo.EventHook(hook.MouseMove, []string{}, func(event hook.Event) {
		handleDrag()
		var operation model.Operation
		operation.X = int(float64(event.X) / utils.Dpi)
		operation.Y = int(float64(event.Y) / utils.Dpi)
		operation.Type = "mouseMove"
		nowTime := time.Now()
		operation.WaitTime = nowTime.Sub(lastTime)
		lastTime = time.Now()
		lastOperation = operation

		steps = append(steps, operation)

	})

	robotgo.EventHook(hook.MouseDown, []string{}, func(event hook.Event) {
		handleDrag()
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
		fmt.Println("click", operation)
	})
	robotgo.EventHook(hook.KeyDown, []string{"esc"}, func(event hook.Event) {
		robotgo.EventEnd()
	})
	robotgo.EventHook(hook.KeyDown, []string{}, func(event hook.Event) {
		handleDrag()
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
	if ok && len(steps) > 1 {
		var finalSteps []model.Operation
		finalSteps = append(finalSteps, steps[0])
		for i := 1; i < len(steps)-1; i++ {

			if steps[i-1].Type == "mouseMove" && steps[i+1].Type == "mouseMove" && steps[i].WaitTime < time.Millisecond*50 {
				continue
			} else {
				finalSteps = append(finalSteps, steps[i])
			}

		}
		finalSteps = append(finalSteps, steps[len(steps)-1])
		marshal, _ := json.Marshal(finalSteps)
		ioutil.WriteFile("./script.txt", marshal, 0666)
		fmt.Println("脚本录制完毕并写入script.txt文件")
		fmt.Println("===========================================================")
	} else {
		fmt.Println("脚本为空")
		fmt.Println("===========================================================")
	}
}
