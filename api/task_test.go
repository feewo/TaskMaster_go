// unit test for TaskCreate function in Api

package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"taskmaster/api"
	"taskmaster/context"
	"taskmaster/entity"
	"taskmaster/storage"
)

var IdTask uint = 68

func TestTaskCreate(t *testing.T) {
	// подготовка запроса
	api := api.Api{}
	newTask := entity.Task{
		Title:  "Test Task",
		UserID: 1,
	}
	jsonBody, err := json.Marshal(newTask)
	if err != nil {
		t.Fatalf("Error marshaling task to JSON: %v", err)
	}
	req, err := http.NewRequest("POST", "/api/task", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()

	ctx := &context.Context{
		Request:  req,
		Response: res,
	}
	// создание task
	api.TaskCreate(ctx)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.Code)
	}

	// Проверка все ли так в бд

	// для этого сначала нужно узнать с какой ID создался task
	type Body struct {
		ID uint
	}
	var ResponseBody Body
	errResponse := json.NewDecoder(res.Body).Decode(&ResponseBody)
	if errResponse != nil {
		t.Errorf("Expected response body")
	}
	IdTask = ResponseBody.ID
	// начало проверок
	task := storage.TaskGet(IdTask)
	if err != nil {
		t.Fatalf("Error getting task from storage: %s", err.Error())
	}

	if task.Title != newTask.Title {
		t.Errorf("Expected task title %s, got %s", newTask.Title, task.Title)
	}

	if task.UserID != newTask.UserID {
		t.Errorf("Expected task user ID %d, got %d", newTask.UserID, task.UserID)
	}
}

func TestTasks(t *testing.T) {
	api := api.Api{}
	res := httptest.NewRecorder()
	ctx := &context.Context{
		Response: httptest.NewRecorder(),
	}
	api.Tasks(ctx)
	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.Code)
	}
}

func TestTask(t *testing.T) {
	api := api.Api{}
	req, err := http.NewRequest("GET", "/api/task/"+strconv.Itoa(int(IdTask)), nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	ctx := &context.Context{
		Request:  req,
		Response: res,
	}
	api.Task(ctx)
	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.Code)
	}
}

func TestTaskDelete(t *testing.T) {
	api := api.Api{}
	req, err := http.NewRequest("DELETE", "/api/task/"+strconv.Itoa(int(IdTask)), nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	ctx := &context.Context{
		Request:  req,
		Response: res,
	}
	api.TaskDelete(ctx)
	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.Code)
	}
}
