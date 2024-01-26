package usecase

import (
	"todo-api/model"
	"todo-api/repository"
)

type TaskUsecase interface {
	GetAllTasks(userID uint) (*[]model.TaskResponse, error)
	CreateTask(task model.Task) (*model.TaskResponse, error)
}

type taskUsecase struct {
	tr repository.TaskRepository
}

func NewTaskUsecase(tr repository.TaskRepository) TaskUsecase {
	return &taskUsecase{tr}
}

func (tu *taskUsecase) GetAllTasks(userID uint) (*[]model.TaskResponse, error) {
	var tasks []model.Task

	if err := tu.tr.GetAllTasks(&tasks, userID); err != nil {
		return nil, err
	}

	var resTasks []model.TaskResponse
	for _, task := range tasks {
		resTask := model.TaskResponse{
			ID:        task.ID,
			Title:     task.Title,
			Completed: task.Completed,
		}
		resTasks = append(resTasks, resTask)
	}

	return &resTasks, nil
}

func (tu *taskUsecase) CreateTask(task model.Task) (*model.TaskResponse, error) {
	if err := tu.tr.CreateTask(&task); err != nil {
		return &model.TaskResponse{}, err
	}

	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		Completed: task.Completed,
	}
	return &resTask, nil
}
