package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
	"fmt"
)

type TaskRepository interface {
	Store(task *model.Task) error
	Update(taskID int, task *model.Task) error
	Delete(id int) error
	GetByID(id int) (*model.Task, error)
	GetList() ([]model.Task, error)
	GetTaskCategory(id int) ([]model.TaskCategory, error)
}

type taskRepository struct {
	filebased *filebased.Data
}

func NewTaskRepo(filebasedDb *filebased.Data) *taskRepository {
	return &taskRepository{
		filebased: filebasedDb,
	}
}

func (t *taskRepository) Store(task *model.Task) error {
	err := t.filebased.StoreTask(*task)
	if err != nil {
		return fmt.Errorf("failed to store task: %v", err)
	}
	return nil
}

func (t *taskRepository) Update(taskID int, task *model.Task) error {
	existingTask, err := t.filebased.GetTaskByID(taskID)
	if err != nil {
		return fmt.Errorf("task with id %d not found: %v", taskID, err)
	}

	existingTask.Title = task.Title
	existingTask.Description = task.Description
	existingTask.DueDate = task.DueDate
	err = t.filebased.StoreTask(*existingTask)
	if err != nil {
		return fmt.Errorf("failed to update task: %v", err)
	}
	return nil
}

func (t *taskRepository) Delete(id int) error {
	task, err := t.filebased.GetTaskByID(id)
	if err != nil {
		return fmt.Errorf("task with id %d not found: %v", id, err)
	}

	err = t.filebased.DeleteTask(*task)
	if err != nil {
		return fmt.Errorf("failed to delete task: %v", err)
	}
	return nil
}

func (t *taskRepository) GetByID(id int) (*model.Task, error) {
	task, err := t.filebased.GetTaskByID(id)
	if err != nil {
		return nil, fmt.Errorf("task with id %d not found: %v", id, err)
	}
	return task, nil
}

func (t *taskRepository) GetList() ([]model.Task, error) {
	tasks, err := t.filebased.GetAllTasks()
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks: %v", err)
	}
	return tasks, nil
}

func (t *taskRepository) GetTaskCategory(id int) ([]model.TaskCategory, error) {
	categories, err := t.filebased.GetTaskCategoriesByTaskID(id)
	if err != nil {
		return nil, fmt.Errorf("task categories for task ID %d not found: %v", id, err)
	}
	return categories, nil
}
