package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/spf13/viper"
	"image/color"
	"task-solver/client/src/settings"
	"task-solver/client/src/ui/components"
)

func main() {
	var err error
	var temp *fyne.App

	// Load all settings and create the app
	if temp, err = createApp(); err != nil {
		panic(fmt.Errorf("could not create the app: %w", err))
	}

	// Create the main window
	application := *temp
	window := createWindow(application)

	// Add the left panel containing the task functionality
	taskListSeparatorContainer := createTaskListPanel()
	taskFormContainer := components.CreateTaskForm()

	// The content of the window
	content := container.New(
		layout.NewBorderLayout(
			nil,
			nil,
			taskListSeparatorContainer,
			nil,
		),
		taskListSeparatorContainer,
		taskFormContainer.Container,
	)

	// TODO: add form content

	window.SetContent(content)
	window.ShowAndRun()
}

func createTaskListPanel() *fyne.Container {
	// Create button functionalities for the task list
	taskListOptions := container.New(
		layout.NewHBoxLayout(),
		widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
			fmt.Println("click on add button")
		}),
		widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
			fmt.Println("click on delete button")
		}),
		widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), func() {
			fmt.Println("click on refresh button")
		}),
	)

	// Create the task list container
	taskListContainer := container.New(
		layout.NewVBoxLayout(),
		taskListOptions,
		// TODO: add live data list
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
	return taskListSeparatorContainer
}

// Create the window for the app
func createWindow(app fyne.App) fyne.Window {
	title := viper.GetString("window.title")

	window := app.NewWindow(title)

	sizes := viper.Sub("window.size")

	windowSize := fyne.NewSize(
		float32(sizes.GetFloat64("width")),
		float32(sizes.GetFloat64("height")),
	)

	window.Resize(windowSize)
	return window
}

func createApp() (*fyne.App, error) {
	if err := settings.LoadSettings(); err != nil {
		return nil, fmt.Errorf("could not load settings: %w", err)
	}

	application := app.New()
	return &application, nil
}
