package ui

import (
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"task-solver/client/src/model"
)

type TaskRemoveComponent struct {
	DeleteButton *widget.Button
}

func CreateTaskRemove() *TaskRemoveComponent {
	taskRemove := &TaskRemoveComponent{}
	taskRemove.DeleteButton = widget.NewButtonWithIcon("Remove", theme.DeleteIcon(), nil)
	taskRemove.Reset()
	return taskRemove
}

func (taskRemove *TaskRemoveComponent) Reset() {
	taskRemove.DeleteButton.Hide()
}

func (taskRemove *TaskRemoveComponent) UseTaskResult(taskResult *model.TaskResult) {
	taskRemove.DeleteButton.Show()
}
