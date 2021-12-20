package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type TaskContext struct {
	TaskUnitsLive           binding.StringList
	TaskUnitComponents      []*TaskUnit
	AppendTaskUnitButton    *widget.Button
	AppendTaskContextButton *widget.Button
	RemoveTaskContextButton *widget.Button
	Container               *fyne.Container
	InnerContainer          *fyne.Container
}

func CreateTaskContext(onAppend func(), onRemove func()) *TaskContext {
	taskContext := &TaskContext{
		TaskUnitComponents: []*TaskUnit{},
	}

	// Create append task unit button
	taskContext.AppendTaskUnitButton = widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		taskContext.AppendTaskUnit(true)
	})

	// Create append task context button
	taskContext.AppendTaskContextButton = widget.NewButtonWithIcon("", theme.MoveDownIcon(), onAppend)

	// Create remove task context button
	taskContext.RemoveTaskContextButton = widget.NewButtonWithIcon("", theme.DeleteIcon(), onRemove)

	// Create the inner container for task unit components
	taskContext.InnerContainer = container.New(
		layout.NewHBoxLayout(),
	)

	// Create a horizontal layout
	taskContext.Container = container.NewHBox(
		taskContext.AppendTaskContextButton,
		taskContext.RemoveTaskContextButton,
		taskContext.InnerContainer,
		layout.NewSpacer(),
		taskContext.AppendTaskUnitButton,
	)

	// Create the array of task units
	taskContext.AppendTaskUnit(false)
	taskContext.TaskUnitComponents[len(taskContext.TaskUnitComponents)-1].Button.Hide()
	return taskContext
}

func (taskContext *TaskContext) AddListeners(onAppend func(), onRemove func()) *TaskContext {
	taskContext.AppendTaskContextButton.OnTapped = onAppend
	taskContext.RemoveTaskContextButton.OnTapped = onRemove
	return taskContext
}

func (taskContext *TaskContext) AppendTaskUnit(isRemovable bool) *TaskContext {
	var onTapped func()
	taskUnit := CreateTaskUnit(nil)

	if isRemovable {
		onTapped = func() {
			// Find the task unit with the same reference
			var taskIndex int
			for index, unit := range taskContext.TaskUnitComponents {
				if unit == taskUnit {
					taskIndex = index
					break
				}
			}

			// Remove the task unit from the UI
			taskContext.InnerContainer.Remove(taskContext.TaskUnitComponents[taskIndex].Container)

			// Remove element at pos taskIndex
			taskContext.TaskUnitComponents = append(
				taskContext.TaskUnitComponents[:taskIndex],
				taskContext.TaskUnitComponents[taskIndex+1:]...,
			)
		}
	}

	taskUnit.Button.OnTapped = onTapped

	// Place the task unit append button after the new task unit
	taskContext.InnerContainer.Add(taskUnit.Container)

	taskContext.TaskUnitComponents = append(
		taskContext.TaskUnitComponents,
		taskUnit,
	)

	return taskContext
}
