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

func MouseClick(step model.Operation) {
	robotgo.MoveClick(step.X, step.Y, step.MouseType)
}
func Input(msg string) {
	robotgo.TypeStr(msg)
}
