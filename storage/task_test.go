package storage_test

import (
	"taskmaster/entity"
	"taskmaster/storage"
	"testing"
)

var User_id uint = 4
var Task_id uint

// как замокать бд
func TestTaskCreate(t *testing.T) {
	ready := false
	newTask := entity.Task{
		Title:  "Test Task",
		Ready:  &ready,
		UserID: User_id,
	}
	createdTask := storage.TaskCreate(newTask)
	Task_id = createdTask.ID
	if createdTask.Title != newTask.Title {
		t.Errorf("Expected title to be %s, but got %s", newTask.Title, createdTask.Title)
	}
	if *createdTask.Ready != *newTask.Ready {
		t.Errorf("Expected ready status to be %t, but got %t", *newTask.Ready, *createdTask.Ready)
	}
	if createdTask.UserID != newTask.UserID {
		t.Errorf("Expected user ID to be %d, but got %d", newTask.UserID, createdTask.UserID)
	}
}

func TestTaskGetAll(t *testing.T) {
	tasks := storage.TaskGetAll()

	if len(tasks) <= 0 {
		t.Error("Expected at least one task, but got none")
	}
}

func TestTaskGet(t *testing.T) {
	testTaskID := Task_id
	task := storage.TaskGet(testTaskID)

	if task.ID != testTaskID {
		t.Errorf("Expected task ID to be %d, but got %d", testTaskID, task.ID)
	}
}

func TestTaskUpdate(t *testing.T) {
	testTaskID := Task_id
	ready := false
	updatedTask := entity.Task{
		Title:  "Updated Task",
		Ready:  &ready,
		UserID: User_id,
	}
	task := storage.TaskUpdate(updatedTask, testTaskID)

	if task.Title != updatedTask.Title {
		t.Errorf("Expected updated title to be %s, but got %s", updatedTask.Title, task.Title)
	}
	if *task.Ready != *updatedTask.Ready {
		t.Errorf("Expected updated ready status to be %t, but got %t", *updatedTask.Ready, *task.Ready)
	}
	if task.UserID != updatedTask.UserID {
		t.Errorf("Expected updated user ID to be %d, but got %d", updatedTask.UserID, task.UserID)
	}
}

func TestTaskDelete(t *testing.T) {
	testTaskID := Task_id
	task := storage.TaskDelete(testTaskID)

	if task.ID != 0 {
		t.Errorf("Error during deletion")
	}
}
