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
	//robotgo.Move()
}
func MouseDrag(step model.Operation) {
	robotgo.Toggle("left")
	robotgo.MilliSleep(50)
	robotgo.Move(step.X, step.Y)
	robotgo.MilliSleep(1200)
	robotgo.Toggle("left", "up")
	//robotgo.DragSmooth(step.X, step.Y, 0.9, 0.9)

}
func MouseClick(step model.Operation) {
	robotgo.MoveClick(step.X, step.Y, step.MouseType)
}
func Input(msg string) {
	robotgo.TypeStr(msg)
}
