package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"github.com/spf13/viper"
	"task-solver/client/src/services"
	"task-solver/client/src/settings"
	"task-solver/client/src/ui"
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

	// Create services
	taskService := services.CreateTaskService()

	// Create form
	taskForm, err := ui.CreateTaskForm(&window, taskService)

	// Check if the form was created properly
	if err != nil {
		err = fmt.Errorf("could not create task form: %w", err)

		dialog.ShowInformation(
			"Fatal Error Occurred",
			err.Error(),
			window,
		)

		window.ShowAndRun()
		panic(err)
	}

	// Create history panel
	historyPanel := ui.CreateHistoryPanel(&window, taskService, taskForm)

	// The content of the window
	content := container.New(
		layout.NewBorderLayout(
			nil,
			nil,
			historyPanel.Container,
			nil,
		),
		historyPanel.Container,
		taskForm.Container,
	)

	window.SetContent(content)
	window.ShowAndRun()
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
