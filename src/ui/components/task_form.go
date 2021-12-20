package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type TaskForm struct {
	TaskContextComponents []*TaskContext
	Container             *fyne.Container
	Content               *widget.Form
}

func CreateTaskForm() *TaskForm {
	taskForm := &TaskForm{}

	// Create the first taskContext
	firstTaskContext := CreateTaskContext(nil, nil)

	// Create components
	taskForm.TaskContextComponents = []*TaskContext{
		firstTaskContext,
	}

	// Bundle form entries
	//taskForm.Content = widget.NewForm(
	//	taskForm.TaskContextComponents.FormItem,
	//)

	// Define container for form
	taskForm.Container = container.NewVBox(
		firstTaskContext.Container,
	)

	return taskForm
}

func (taskForm *TaskForm) AddTaskContext(addListeners bool) *TaskForm {
	taskContext := CreateTaskContext(nil, nil)

	var (
		onAppend func()
		onRemove func()
	)

	if addListeners {
		onAppend = func() {

		}
		onRemove = func() {

		}
	}

	taskContext.AddListeners(onAppend, onRemove)
	return taskForm
}

//func (taskForm *TaskForm) AddTaskContext() *TaskForm {
//	taskContext := CreateTaskUnit()
//
//	taskForm.TaskContextComponents = append(
//		taskForm.TaskContextComponents,
//		taskContext,
//	)
//
//	taskForm.Content.AppendItem(
//		taskContext.FormItem,
//	)
//
//	return taskForm
//}
//
//func (taskForm *TaskForm) RemoveTaskContext
