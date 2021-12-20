package ui

import (
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"task-solver/client/src/model"
)

type TaskIndexComponent struct {
	SelectEntry *widget.SelectEntry
	FormItem    *widget.FormItem
	Options     binding.StringList
}

func CreateTaskIndex() *TaskIndexComponent {
	taskIndex := &TaskIndexComponent{}

	taskIndex.Options = binding.NewStringList()
	options, _ := taskIndex.Options.Get()
	taskIndex.SelectEntry = widget.NewSelectEntry(options)
	taskIndex.FormItem = widget.NewFormItem("Index", taskIndex.SelectEntry)

	// Listen to changes on inpu	t data
	taskIndex.Options.AddListener(binding.NewDataListener(func() {
		// Get info from live data
		options, _ := taskIndex.Options.Get()

		// Update the component
		taskIndex.SelectEntry.SetOptions(options)

		// Mark selected text
		if len(options) != 0 {
			taskIndex.SelectEntry.Entry.Text = options[0]
		}

		// Refresh the component
		taskIndex.SelectEntry.Refresh()
	}))

	taskIndex.Reset()

	return taskIndex
}

func (taskIndex *TaskIndexComponent) Reset() {
	options, _ := taskIndex.Options.Get()

	// Mark selected text
	if len(options) != 0 {
		taskIndex.SelectEntry.Entry.Text = options[0]
	}

	taskIndex.SelectEntry.Entry.Refresh()
}

func (taskIndex *TaskIndexComponent) UseTaskResult(taskResult *model.TaskResult) {
	taskIndex.SelectEntry.Entry.Text = strconv.FormatInt(taskResult.Task.Index, 10)
	taskIndex.SelectEntry.Entry.Refresh()
}
