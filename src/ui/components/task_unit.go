package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type TaskUnit struct {
	Entry     *widget.Entry
	Data      binding.String
	Button    *widget.Button
	Container *fyne.Container
}

func CreateTaskUnit(onTapped func()) *TaskUnit {
	taskUnit := &TaskUnit{}

	taskUnit.Data = binding.NewString()
	taskUnit.Entry = widget.NewEntryWithData(taskUnit.Data)
	taskUnit.Button = widget.NewButtonWithIcon("", theme.DeleteIcon(), onTapped)
	taskUnit.Container = container.NewHBox(
		taskUnit.Entry,
		taskUnit.Button,
	)

	return taskUnit
}
