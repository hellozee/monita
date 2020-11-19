package main

import "encoding/json"

type Task struct {
	ID         int    `json:"id"`
	Desciption string `json:"description"`
}

type TaskModel struct {
	tasks        []Task
	currentIndex int
}

func (tm *TaskModel) AddTask(description string) {
	ts := Task{
		ID:         tm.currentIndex,
		Desciption: description,
	}
	tm.currentIndex++
	tm.tasks = append(tm.tasks, ts)
}

func (tm *TaskModel) DeleteTask(id int) {
	tm.tasks = append(tm.tasks[:id], tm.tasks[id+1:]...)
}

func (tm *TaskModel) ToJSON() []byte {
	bytes, _ := json.Marshal(tm.tasks)
	return bytes
}

func NewTaskModel() TaskModel {
	return TaskModel{
		currentIndex: 0,
		tasks:        []Task{},
	}
}
