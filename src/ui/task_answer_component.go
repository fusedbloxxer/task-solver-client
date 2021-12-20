package ui

import (
	"fmt"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/spf13/viper"
	"task-solver/client/src/constants"
	"task-solver/client/src/model"
)

type TaskAnswerComponent struct {
	Label    *widget.Label
	FormItem *widget.FormItem
	LiveData binding.String
}

func CreateTaskAnswer() *TaskAnswerComponent {
	taskAnswer := &TaskAnswerComponent{}

	taskAnswer.LiveData = binding.NewString()
	taskAnswer.Label = widget.NewLabelWithData(taskAnswer.LiveData)
	taskAnswer.FormItem = widget.NewFormItem("Answer", taskAnswer.Label)
	taskAnswer.Reset()

	return taskAnswer
}

func (taskAnswer *TaskAnswerComponent) Reset() {
	_ = taskAnswer.LiveData.Set(viper.GetString(constants.UndefinedTextValue))
}

func (taskAnswer *TaskAnswerComponent) UseTaskResult(taskResult *model.TaskResult) {
	answer := fmt.Sprintf("%f", taskResult.Answer)
	_ = taskAnswer.LiveData.Set(answer)
}
