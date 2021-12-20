package ui

import (
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/spf13/viper"
	"task-solver/client/src/constants"
	"task-solver/client/src/model"
)

type TaskIdComponent struct {
	Label    *widget.Label
	FormItem *widget.FormItem
	LiveData binding.String
}

func CreateTaskId() *TaskIdComponent {
	taskId := &TaskIdComponent{}

	taskId.LiveData = binding.NewString()
	taskId.Label = widget.NewLabelWithData(taskId.LiveData)
	taskId.FormItem = widget.NewFormItem("Id", taskId.Label)
	taskId.Reset()

	return taskId
}

func (taskId *TaskIdComponent) Reset() {
	_ = taskId.LiveData.Set(viper.GetString(constants.UndefinedTextValue))
}

func (taskId *TaskIdComponent) UseTaskResult(taskResult *model.TaskResult) {
	_ = taskId.LiveData.Set(taskResult.Id)
}
