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

var IdTaskPoint uint = 68

func TestTaskPointCreate(t *testing.T) {
	// подготовка запроса
	api := api.Api{}
	newTaskPoint := entity.TaskPoint{
		Title:  "Test TaskPoint",
		TaskID: IdTask,
	}
	jsonBody, err := json.Marshal(newTaskPoint)
	if err != nil {
		t.Fatalf("Error marshaling taskPoint to JSON: %v", err)
	}
	req, err := http.NewRequest("POST", "/api/taskpoint", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()

	ctx := &context.Context{
		Request:  req,
		Response: res,
	}
	// создание taskPoint
	api.TaskPointCreate(ctx)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.Code)
	}

	// Проверка все ли так в бд

	// для этого сначала нужно узнать с какой ID создался taskPoint
	type Body struct {
		ID uint
	}
	var ResponseBody Body
	errResponse := json.NewDecoder(res.Body).Decode(&ResponseBody)
	if errResponse != nil {
		t.Errorf("Expected response body")
	}
	IdTaskPoint = ResponseBody.ID
	// начало проверок
	taskPoint := storage.TaskPointGet(IdTaskPoint)
	if err != nil {
		t.Fatalf("Error getting taskPoint from storage: %s", err.Error())
	}

	if taskPoint.Title != newTaskPoint.Title {
		t.Errorf("Expected taskPoint title %s, got %s", newTaskPoint.Title, taskPoint.Title)
	}

	if taskPoint.TaskID != newTaskPoint.TaskID {
		t.Errorf("Expected taskPoint user ID %d, got %d", newTaskPoint.TaskID, taskPoint.TaskID)
	}
}

func TestTaskPoints(t *testing.T) {
	api := api.Api{}
	res := httptest.NewRecorder()
	ctx := &context.Context{
		Response: httptest.NewRecorder(),
	}
	api.TaskPoints(ctx)
	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.Code)
	}
}

func TestTaskPoint(t *testing.T) {
	api := api.Api{}
	req, err := http.NewRequest("GET", "/api/taskpoint/"+strconv.Itoa(int(IdTask)), nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	ctx := &context.Context{
		Request:  req,
		Response: res,
	}
	api.TaskPoint(ctx)
	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.Code)
	}
}

func TestTaskPointUpdate(t *testing.T) {
	api := api.Api{}
	newTaskPoint := entity.TaskPoint{
		Title:  "Test TaskPoint",
		TaskID: IdTask,
	}
	jsonBody, err := json.Marshal(newTaskPoint)
	if err != nil {
		t.Fatalf("Error marshaling taskPoint to JSON: %v", err)
	}
	req, err := http.NewRequest("PUT", "/api/taskpoint/"+strconv.Itoa(int(IdTaskPoint)), bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	ctx := &context.Context{
		Request:  req,
		Response: res,
	}
	api.TaskPointUpdate(ctx)
	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.Code)
	}
	// начало проверок
	taskPoint := storage.TaskPointGet(IdTaskPoint)
	if err != nil {
		t.Fatalf("Error getting taskPoint from storage: %s", err.Error())
	}

	if taskPoint.Title != newTaskPoint.Title {
		t.Errorf("Expected taskPoint title %s, got %s", newTaskPoint.Title, taskPoint.Title)
	}

	if taskPoint.TaskID != newTaskPoint.TaskID {
		t.Errorf("Expected taskPoint user ID %d, got %d", newTaskPoint.TaskID, taskPoint.TaskID)
	}
}

func TestTaskPointDelete(t *testing.T) {
	api := api.Api{}
	req, err := http.NewRequest("DELETE", "/api/taskpoint/"+strconv.Itoa(int(IdTaskPoint)), nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	ctx := &context.Context{
		Request:  req,
		Response: res,
	}
	api.TaskPointDelete(ctx)
	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.Code)
	}
}
