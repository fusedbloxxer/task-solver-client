package ui

import (
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"task-solver/client/src/model"
	"task-solver/client/src/utils"
)

type TaskContextComponent struct {
	Entry    *widget.Entry
	FormItem *widget.FormItem
	LiveData binding.String
}

func CreateTaskContext() *TaskContextComponent {
	taskContext := &TaskContextComponent{}

	taskContext.LiveData = binding.NewString()
	taskContext.Entry = widget.NewEntryWithData(taskContext.LiveData)
	taskContext.Entry.MultiLine = true
	taskContext.FormItem = widget.NewFormItem("Context", taskContext.Entry)

	return taskContext
}

func (taskContext *TaskContextComponent) Reset() {
	taskContext.Entry.Text = ""
	taskContext.Entry.Refresh()
}

func (taskContext *TaskContextComponent) UseTaskResult(taskResult *model.TaskResult) {
	taskContext.Entry.Text = utils.CompressContext(taskResult.Task.Context)
	taskContext.Entry.Refresh()
}
