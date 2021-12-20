package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"task-solver/client/src/model"
	"task-solver/client/src/services"
	"task-solver/client/src/utils"
)

type TaskForm struct {
	// Form Items
	TaskContextComponent *TaskContextComponent
	TaskIndexComponent   *TaskIndexComponent
	TaskSubmitComponent  *TaskSubmitComponent
	TaskRemoveComponent  *TaskRemoveComponent

	// Presentational components
	TaskAnswerComponent *TaskAnswerComponent
	TaskTitleComponent  *TaskTitleComponent
	TaskIdComponent     *TaskIdComponent

	// Structure
	Content   *widget.Form
	Container *fyne.Container

	// Dependencies
	Parent      *fyne.Window
	TaskService services.ITaskService

	// External events to execute
	onTaskAdded   func(task *model.TaskResult)
	onTaskRemoved func(taskId string)
}

func CreateTaskForm(parent *fyne.Window, service services.ITaskService) (*TaskForm, error) {
	taskForm := &TaskForm{
		Parent:        parent,
		TaskService:   service,
		onTaskAdded:   func(result *model.TaskResult) {},
		onTaskRemoved: func(taskId string) {},
	}

	taskForm.createComponents()

	taskForm.createForm()

	taskForm.attachButtonListeners()

	taskForm.bundleComponents()

	if err := taskForm.PopulateWithData(); err != nil {
		return nil, err
	}

	return taskForm, nil
}

func (taskForm *TaskForm) AddTaskRemovedListener(onTaskRemoved func(taskId string)) {
	taskForm.onTaskRemoved = onTaskRemoved
}

func (taskForm *TaskForm) AddTaskAddedListener(onTaskAdded func(task *model.TaskResult)) {
	taskForm.onTaskAdded = onTaskAdded
}

func (taskForm *TaskForm) PopulateWithData() error {
	indexes, err := taskForm.TaskService.GetAllIndexes()

	if err != nil {
		return fmt.Errorf("could not populate with data: %w", err)
	}

	_ = taskForm.TaskIndexComponent.Options.Set(indexes)

	return nil
}

func (taskForm *TaskForm) UseTaskResult(taskResult *model.TaskResult) {
	taskForm.TaskContextComponent.UseTaskResult(taskResult)
	taskForm.TaskRemoveComponent.UseTaskResult(taskResult)
	taskForm.TaskAnswerComponent.UseTaskResult(taskResult)
	taskForm.TaskIndexComponent.UseTaskResult(taskResult)
	taskForm.TaskIdComponent.UseTaskResult(taskResult)
}

func (taskForm *TaskForm) Reset() {
	taskForm.TaskContextComponent.Reset()
	taskForm.TaskRemoveComponent.Reset()
	taskForm.TaskAnswerComponent.Reset()
	taskForm.TaskIndexComponent.Reset()
	taskForm.TaskIdComponent.Reset()
}

func (taskForm *TaskForm) createComponents() *TaskForm {
	taskForm.TaskContextComponent = CreateTaskContext()
	taskForm.TaskAnswerComponent = CreateTaskAnswer()
	taskForm.TaskSubmitComponent = CreateTaskSubmit()
	taskForm.TaskRemoveComponent = CreateTaskRemove()
	taskForm.TaskIndexComponent = CreateTaskIndex()
	taskForm.TaskTitleComponent = CreateTaskTitle()
	taskForm.TaskIdComponent = CreateTaskId()
	return taskForm
}

func (taskForm *TaskForm) createForm() *TaskForm {
	taskForm.Content = widget.NewForm(
		taskForm.TaskIdComponent.FormItem,
		taskForm.TaskAnswerComponent.FormItem,
		taskForm.TaskContextComponent.FormItem,
		taskForm.TaskIndexComponent.FormItem,
	)
	return taskForm
}

func (taskForm *TaskForm) attachButtonListeners() *TaskForm {
	taskForm.attachOnSubmit()
	taskForm.attachOnRemove()
	return taskForm
}

func (taskForm *TaskForm) attachOnSubmit() {
	taskForm.TaskSubmitComponent.SubmitButton.OnTapped = func() {
		go func() {
			index, err := strconv.Atoi(taskForm.TaskIndexComponent.SelectEntry.Text)

			if err != nil {
				fmt.Printf("could not parse the index in the form: %s", err.Error())
			}

			taskResult, err := taskForm.TaskService.SolveTask(model.Task{
				Context: utils.ExtractContext(taskForm.TaskContextComponent.Entry.Text),
				Index:   int64(index),
			})

			if err != nil {
				dialog.ShowInformation(
					"Error occurred",
					fmt.Errorf("could not solve the task: %w", err).Error(),
					*taskForm.Parent,
				)
				return
			}

			taskForm.UseTaskResult(taskResult)
			taskForm.onTaskAdded(taskResult)
		}()
	}
}

func (taskForm *TaskForm) attachOnRemove() {
	taskForm.TaskRemoveComponent.DeleteButton.OnTapped = func() {
		go func() {
			taskId, _ := taskForm.TaskIdComponent.LiveData.Get()

			dialog.ShowConfirm(
				"Delete the task?",
				"Are you sure you want to delete the task "+taskId,
				func(positive bool) {
					if !positive {
						return
					}

					if err := taskForm.TaskService.RemoveTask(taskId); err != nil {
						dialog.ShowInformation(
							"Error occurred",
							fmt.Errorf("could not remove the task: %w", err).Error(),
							*taskForm.Parent,
						)
						return
					}

					taskForm.Reset()
					taskForm.onTaskRemoved(taskId)
				},
				*taskForm.Parent,
			)
		}()
	}
}

func (taskForm *TaskForm) bundleComponents() *TaskForm {
	taskForm.Container = container.NewVBox(
		taskForm.TaskTitleComponent.Container,
		taskForm.Content,
		container.NewHBox(
			layout.NewSpacer(),
			taskForm.TaskSubmitComponent.SubmitButton,
			taskForm.TaskRemoveComponent.DeleteButton,
			layout.NewSpacer(),
		),
	)
	return taskForm
}
