package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type TaskTitleComponent struct {
	Label     *widget.Label
	Container *fyne.Container
}

func CreateTaskTitle() *TaskTitleComponent {
	taskTitle := &TaskTitleComponent{}
	taskTitle.Label = widget.NewLabel("Task")
	taskTitle.Label.TextStyle = fyne.TextStyle{Bold: true}
	taskTitle.Container = container.NewHBox(
		layout.NewSpacer(),
		taskTitle.Label,
		layout.NewSpacer(),
	)
	return taskTitle
}
