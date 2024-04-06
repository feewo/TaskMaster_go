package storage_test

import (
	"testing"

	"taskmaster/entity"
	"taskmaster/storage"
)

var TPTask_id uint = 25
var TaskPoint_id uint

func TestTaskPointCreate(t *testing.T) {
	ready := true
	taskPoint := entity.TaskPoint{
		Title:  "Test Task",
		Ready:  &ready,
		TaskID: TPTask_id,
	}
	createdTaskPoint := storage.TaskPointCreate(taskPoint)
	TaskPoint_id = createdTaskPoint.ID
	if createdTaskPoint == nil {
		t.Errorf("TaskPointCreate did not return a valid entity")
	}
}

func TestTaskPointGetAll(t *testing.T) {
	taskPoints := storage.TaskPointGetAll()
	if len(taskPoints) == 0 {
		t.Errorf("TaskPointGetAll did not return any task points")
	}
}

func TestTaskPointGet(t *testing.T) {
	taskPoint := storage.TaskPointGet(TaskPoint_id)
	if taskPoint == nil {
		t.Errorf("TaskPointGet did not return a valid entity")
	}
}

func TestTaskPointUpdate(t *testing.T) {
	ready := true
	taskPoint := entity.TaskPoint{
		Title:  "Updated Task",
		Ready:  &ready,
		TaskID: TPTask_id,
	}
	updatedTaskPoint := storage.TaskPointUpdate(taskPoint, TaskPoint_id)
	if updatedTaskPoint == nil {
		t.Errorf("TaskPointUpdate did not return a valid entity")
	}
}

func TestTaskPointDelete(t *testing.T) {
	deletedTaskPoint := storage.TaskPointDelete(TaskPoint_id)
	if deletedTaskPoint == nil {
		t.Errorf("TaskPointDelete did not return a valid entity")
	}
}
