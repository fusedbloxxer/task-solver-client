package ui

import "task-solver/client/src/model"

type Component interface {
}

type Resettable interface {
	Reset()
}

type TaskResultHolder interface {
	UseTaskResult(task *model.TaskResult)
}

type TaskManager interface {
	Resettable

	TaskResultHolder

	AddTaskRemovedListener(onTaskRemoved func(taskId string))

	AddTaskAddedListener(onTaskAdded func(task *model.TaskResult))
}
