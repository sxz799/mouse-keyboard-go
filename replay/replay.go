package replay

import (
	"GSAutoHSProject/model"
	"GSAutoHSProject/utils"
	"encoding/json"
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"io/ioutil"
	"os"
	"time"
)

func DoReplal() {
	fmt.Println("正在执行中...(您可以按 Ctrl+Q 结束本程序)")
	go func() {
		robotgo.EventHook(hook.KeyDown, []string{"q", "ctrl"}, func(event hook.Event) {
			robotgo.EventEnd()
			os.Exit(1)
		})
		s := robotgo.EventStart()
		<-robotgo.EventProcess(s)
	}()
	f1, _ := ioutil.ReadFile("./script.txt")
	var steps []model.Operation
	json.Unmarshal(f1, &steps)
	if len(steps) < 1 {
		fmt.Println("录制内容为空")
	}
	for _, step := range steps {
		time.Sleep(step.WaitTime)
		switch step.Type {
		case "mouseClick":
			utils.MouseClick(step)
		case "keyboard":
			utils.Input(step.InputMsg)
		}
	}
	fmt.Println("============================================================")
}
