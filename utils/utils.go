package utils

import (
	"GSAutoHSProject/model"
	"github.com/go-vgo/robotgo"
)

var Dpi float64

func Init() {
	x1, _ := robotgo.GetScreenSize()
	x2, _ := robotgo.GetScaleSize()
	Dpi = float64(x2) / float64(x1)
}

func MouseMove(step model.Operation) {
	robotgo.Move(step.X, step.Y)
}
func MouseDrag(step model.Operation) {
	robotgo.Toggle("left")
	robotgo.MilliSleep(50)
	robotgo.MoveSmooth(step.X, step.Y, 0.7, 0.7)
	robotgo.MilliSleep(1200)
	robotgo.Toggle("left", "up")

}
func MouseClick(step model.Operation) {
	robotgo.MoveClick(step.X, step.Y, step.MouseType)
}

func KeyboardDown(step model.Operation) {
	robotgo.KeyDown(step.Key)
}

func KeyboardDownWithCtrl(step model.Operation) {
	robotgo.KeySleep = 100
	robotgo.KeyTap(step.Key, "ctrl")
}

func KeyboardDownWithAlt(step model.Operation) {
	robotgo.KeySleep = 100
	robotgo.KeyTap(step.Key, "alt")
}

func InputStr(msg string) {
	robotgo.TypeStr(msg)
}
