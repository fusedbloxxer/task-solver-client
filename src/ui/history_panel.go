package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strings"
	"task-solver/client/src/model"
	"task-solver/client/src/services"
)

type HistoryPanel struct {
	// The container containing all child elements
	Container *fyne.Container

	// Dependencies
	Parent       *fyne.Window
	TaskService  services.ITaskService
	ContentPanel TaskManager

	// The tasks that are stored shown locally
	Tasks binding.UntypedList
}

func CreateHistoryPanel(parent *fyne.Window, taskService services.ITaskService, contentPanel TaskManager) *HistoryPanel {
	historyPanel := &HistoryPanel{
		TaskService:  taskService,
		Parent:       parent,
		ContentPanel: contentPanel,
		Tasks:        binding.NewUntypedList(),
	}

	historyPanel.createTaskListPanel()

	historyPanel.refreshTaskListData()

	historyPanel.handleTaskEvents()

	return historyPanel
}

func (historyPanel *HistoryPanel) handleTaskEvents() *HistoryPanel {
	historyPanel.ContentPanel.AddTaskRemovedListener(func(taskId string) {
		historyPanel.refreshTaskListData()
	})

	historyPanel.ContentPanel.AddTaskAddedListener(func(taskResult *model.TaskResult) {
		historyPanel.refreshTaskListData()
	})

	return historyPanel
}

func (historyPanel *HistoryPanel) createTaskListPanel() *HistoryPanel {
	buttonPanel := historyPanel.createButtonPanel()
	taskList := historyPanel.createTaskList()

	// Create the task list container
	taskListContainer := container.New(
		layout.NewBorderLayout(
			buttonPanel,
			nil,
			nil,
			nil,
		),
		buttonPanel,
		taskList,
	)

	// Create boundary separator
	separationLine := canvas.NewLine(color.White)
	separationLine.StrokeWidth = 5

	// Create separator container
	taskListSeparatorContainer := container.New(
		layout.NewHBoxLayout(),
		taskListContainer,
		separationLine,
	)

	// Add the container to the panel
	historyPanel.Container = taskListSeparatorContainer

	return historyPanel
}

func (historyPanel *HistoryPanel) createTaskList() *fyne.Container {
	return container.New(
		layout.NewMaxLayout(),
		widget.NewListWithData(
			historyPanel.Tasks,
			func() fyne.CanvasObject {
				w := NewTappableLabel(strings.Repeat("-", 50), nil)
				w.Alignment = fyne.TextAlignCenter
				return w
			},
			func(item binding.DataItem, object fyne.CanvasObject) {
				untypedItem, _ := item.(binding.Untyped).Get()
				taskItem := untypedItem.(model.TaskResult)
				tapLabel := object.(*TapLabel)
				tapLabel.Bind(binding.BindString(&taskItem.Id))
				tapLabel.OnTap = func() {
					historyPanel.ContentPanel.UseTaskResult(&taskItem)
				}
			},
		),
	)
}

func (historyPanel *HistoryPanel) createButtonPanel() *fyne.Container {
	return container.New(
		layout.NewHBoxLayout(),
		layout.NewSpacer(),
		historyPanel.createAddButton(),
		historyPanel.createDeleteButton(),
		historyPanel.createRefreshButton(),
		layout.NewSpacer(),
	)
}

func (historyPanel *HistoryPanel) createRefreshButton() *widget.Button {
	return widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), func() {
		fmt.Println("click on refresh button")
		historyPanel.refreshTaskListData()
	})
}

func (historyPanel *HistoryPanel) refreshTaskListData() {
	go func() {
		tasks, err := historyPanel.TaskService.GetAllTasks()

		if err != nil {
			dialog.ShowInformation(
				"Error occurred",
				fmt.Errorf("could not retrieve the tasks: %w", err).Error(),
				*historyPanel.Parent,
			)
			return
		}

		untypedList := make([]interface{}, 0)

		for _, task := range tasks {
			untypedList = append(untypedList, task)
		}

		_ = historyPanel.Tasks.Set(untypedList)
	}()
}

func (historyPanel *HistoryPanel) createDeleteButton() *widget.Button {
	return widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
		fmt.Println("click on delete button")
		go func() {
			dialog.ShowConfirm("Delete ALL Tasks?",
				"Are you really sure you want to delete ALL Tasks?",
				func(positive bool) {
					if !positive {
						return
					}

					err := historyPanel.TaskService.RemoveAllTasks()

					if err != nil {
						dialog.ShowInformation(
							"Error occurred",
							fmt.Errorf("could not remove the tasks: %w", err).Error(),
							*historyPanel.Parent,
						)
						return
					}

					_ = historyPanel.Tasks.Set(make([]interface{}, 0))
				},
				*historyPanel.Parent,
			)
		}()
	})
}

func (historyPanel *HistoryPanel) createAddButton() *widget.Button {
	return widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		fmt.Println("click on add button")
		go func() {
			historyPanel.ContentPanel.Reset()
		}()
	})
}
