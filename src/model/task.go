package model

type Task struct {
	Context [][]string `json:"context"`
	Index   int64      `json:"index"`
}

type TaskResult struct {
	Id     string  `json:"id"`
	Task   Task    `json:"task"`
	Answer float64 `json:"answer"`
}
