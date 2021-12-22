package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"task-solver/client/src/model"
)

type ITaskService interface {
	SolveTask(task model.Task) (*model.TaskResult, error)

	GetAllTasks() ([]model.TaskResult, error)

	GetAllIndexes() ([]string, error)

	RemoveTask(taskId string) error

	RemoveAllTasks() error
}

type TaskService struct {
}

func CreateTaskService() ITaskService {
	taskService := &TaskService{}
	return taskService
}

func (taskService *TaskService) GetAllIndexes() ([]string, error) {
	res, err := http.Get(viper.GetString("api.baseUrl") + "/tasks/indexes")

	if err != nil {
		return nil, fmt.Errorf("could not retrieve task indexes: %w", err)
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		message, _ := ioutil.ReadAll(res.Body)
		return nil, fmt.Errorf("could not retrieve the indexes: %s", string(message))
	}

	indexes := make([]int64, 0)
	if err := json.NewDecoder(res.Body).Decode(&indexes); err != nil {
		return nil, err
	}

	stringIndexes := make([]string, 0, len(indexes))
	for _, index := range indexes {
		stringIndexes = append(stringIndexes, strconv.FormatInt(index, 10))
	}

	sort.Strings(stringIndexes)
	return stringIndexes, nil
}

func (taskService *TaskService) SolveTask(task model.Task) (*model.TaskResult, error) {
	req, err := json.Marshal(task)

	if err != nil {
		return nil, fmt.Errorf("could not serialize the task: %w", err)
	}

	res, err := http.Post(
		viper.GetString("api.baseUrl")+"/tasks/solve",
		"application/json",
		bytes.NewBuffer(req),
	)

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		message, _ := ioutil.ReadAll(res.Body)
		return nil, fmt.Errorf("could not solve the task: %s", string(message))
	}

	var taskResult model.TaskResult
	if err := json.NewDecoder(res.Body).Decode(&taskResult); err != nil {
		return nil, fmt.Errorf("could not deserialize the task result: %w", err)
	}

	return &taskResult, nil
}

func (taskService *TaskService) RemoveTask(taskId string) error {
	req, err := http.NewRequest("DELETE", viper.GetString("api.baseUrl")+"/tasks/"+taskId, nil)

	if err != nil {
		return fmt.Errorf("could not create the delete request: %w", err)
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil || res.StatusCode != http.StatusOK {
		return fmt.Errorf("could not delete the task %s: %w", taskId, err)
	}

	return nil
}

func (taskService *TaskService) GetAllTasks() ([]model.TaskResult, error) {
	res, err := http.Get(viper.GetString("api.baseUrl") + "/tasks")

	if err != nil {
		return nil, fmt.Errorf("could not retrieve all tasks: %w", err)
	}

	body, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not retrieve the tasks: %s", string(body))
	}

	taskResults := make([]model.TaskResult, 0)

	if err := json.Unmarshal(body, &taskResults); err != nil {
		return nil, fmt.Errorf("could not deserialize the tasks: %w", err)
	}

	return taskResults, nil
}

func (taskService *TaskService) RemoveAllTasks() error {
	req, err := http.NewRequest("DELETE", viper.GetString("api.baseUrl")+"/tasks", nil)

	if err != nil {
		return fmt.Errorf("could not create delete request: %w", err)
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return fmt.Errorf("could not delete the tasks: %w", err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("could not delete the tasks: %s", string(body))
	}

	return nil
}
