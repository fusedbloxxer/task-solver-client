package ui

import (
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type TaskSubmitComponent struct {
	SubmitButton *widget.Button
}

func CreateTaskSubmit() *TaskSubmitComponent {
	taskSubmit := &TaskSubmitComponent{}
	taskSubmit.SubmitButton = widget.NewButtonWithIcon("Solve", theme.ComputerIcon(), nil)
	return taskSubmit
}
