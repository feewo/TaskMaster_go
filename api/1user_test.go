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
	"taskmaster/db"
	"taskmaster/entity"
	"taskmaster/storage"
)

var IdUser uint = 1
var TokenUser string

func TestUsertCreate(t *testing.T) {
	db.Migrate()
	// подготовка запроса
	api := api.Api{}
	newUser := entity.User{
		Surname:    "Иванов",
		Name:       "Иван",
		Patronymic: "Иванович",
		Login:      "_",
		Email:      "test@pochta.ru",
		Password:   "1",
	}
	jsonBody, err := json.Marshal(newUser)
	if err != nil {
		t.Fatalf("Error marshaling User to JSON: %v", err)
	}
	req, err := http.NewRequest("POST", "/api/user", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()

	ctx := &context.Context{
		Request:  req,
		Response: res,
	}
	// создание User
	api.UserCreate(ctx)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.Code)
	}

	// Проверка все ли так в бд

	// для этого сначала нужно узнать с какой ID создался User
	type Body struct {
		ID uint
	}
	var ResponseBody Body
	errResponse := json.NewDecoder(res.Body).Decode(&ResponseBody)
	if errResponse != nil {
		t.Errorf("Expected response body")
	}
	IdUser = ResponseBody.ID
	// начало проверок
	User := storage.UserGet(IdUser)
	if err != nil {
		t.Fatalf("Error getting User from storage: %s", err.Error())
	}

	if User.Surname != newUser.Surname {
		t.Errorf("Expected User surname %s, got %s", newUser.Surname, User.Surname)
	}

	if User.Name != newUser.Name {
		t.Errorf("Expected User name ID %s, got %s", newUser.Name, User.Name)
	}

	if User.Patronymic != newUser.Patronymic {
		t.Errorf("Expected User patronymic ID %s, got %s", newUser.Patronymic, User.Patronymic)
	}

	if User.Login != newUser.Login {
		t.Errorf("Expected User login ID %s, got %s", newUser.Login, User.Login)
	}

	if User.Email != newUser.Email {
		t.Errorf("Expected User email ID %s, got %s", newUser.Email, User.Email)
	}

	if User.Role != "user" {
		t.Errorf("Expected User role ID %s, got %s", newUser.Role, "user")
	}
}

func TestUsers(t *testing.T) {
	api := api.Api{}
	res := httptest.NewRecorder()
	ctx := &context.Context{
		Response: httptest.NewRecorder(),
	}
	api.Users(ctx)
	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.Code)
	}
}

func TestUser(t *testing.T) {
	api := api.Api{}
	req, err := http.NewRequest("GET", "/api/user/"+strconv.Itoa(int(IdUser)), nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	ctx := &context.Context{
		Request:  req,
		Response: res,
	}
	api.User(ctx)
	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.Code)
	}
}

func TestUserUpdate(t *testing.T) {
	api := api.Api{}
	newUser := entity.User{
		Surname:    "Петров",
		Name:       "Петр",
		Patronymic: "Петрович",
		Login:      ".",
		Email:      "proverka@mail.com",
		Password:   "12",
	}
	jsonBody, err := json.Marshal(newUser)
	if err != nil {
		t.Fatalf("Error marshaling User to JSON: %v", err)
	}
	req, err := http.NewRequest("PUT", "/api/user/"+strconv.Itoa(int(IdUser)), bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	ctx := &context.Context{
		Request:  req,
		Response: res,
	}
	api.UserUpdate(ctx)
	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.Code)
	}
	// начало проверок
	User := storage.UserGet(IdUser)
	if err != nil {
		t.Fatalf("Error getting User from storage: %s", err.Error())
	}

	if User.Surname != newUser.Surname {
		t.Errorf("Expected User surname %s, got %s", newUser.Surname, User.Surname)
	}

	if User.Name != newUser.Name {
		t.Errorf("Expected User name ID %s, got %s", newUser.Name, User.Name)
	}

	if User.Patronymic != newUser.Patronymic {
		t.Errorf("Expected User patronymic ID %s, got %s", newUser.Patronymic, User.Patronymic)
	}

	if User.Login != newUser.Login {
		t.Errorf("Expected User login ID %s, got %s", newUser.Login, User.Login)
	}

	if User.Email != newUser.Email {
		t.Errorf("Expected User email ID %s, got %s", newUser.Email, User.Email)
	}

	if User.Role != "user" {
		t.Errorf("Expected User role ID %s, got %s", newUser.Role, "user")
	}
}

func TestUserAuth(t *testing.T) {
	api := api.Api{}
	newUser := entity.User{
		Login:    ".",
		Password: "12",
	}
	jsonBody, err := json.Marshal(newUser)
	if err != nil {
		t.Fatalf("Error marshaling User to JSON: %v", err)
	}
	req, err := http.NewRequest("POST", "/api/token", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	ctx := &context.Context{
		Request:  req,
		Response: res,
	}
	api.UserAuth(ctx)
	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.Code)
	}
}

func TestIsAuth(t *testing.T) {
	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/user", nil)
	if err != nil {
		t.Fatal(err)
	}
	ctx := &context.Context{
		Response: res,
		Request:  req,
	}
	api.IsAuth(ctx, "GET")
	if res.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, res.Code)
	}
}

func TestUserDelete(t *testing.T) {
	api := api.Api{}
	req, err := http.NewRequest("DELETE", "/api/user/"+strconv.Itoa(int(IdUser)), nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	ctx := &context.Context{
		Request:  req,
		Response: res,
	}
	api.UserDelete(ctx)
	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.Code)
	}
}
