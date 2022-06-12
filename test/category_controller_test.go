package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"

	"sudutkampus/gorestfulapi/app"
	"sudutkampus/gorestfulapi/controller"
	"sudutkampus/gorestfulapi/helper"
	"sudutkampus/gorestfulapi/middleware"
	"sudutkampus/gorestfulapi/model/domain"
	"sudutkampus/gorestfulapi/repository"
	"sudutkampus/gorestfulapi/service"

	"github.com/go-playground/validator/v10"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/gorestfulapitest")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func truncateCategory(db *sql.DB) {
	db.Exec("TRUNCATE categories")
}

func TestCreateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name": "Gadget"}`)
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	router.ServeHTTP(recorder, request)

	var responseBody map[string]interface{}
	body, err := io.ReadAll(recorder.Result().Body)
	helper.PanicIfError(err)
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
}

func TestCreateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name": ""}`)
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	router.ServeHTTP(recorder, request)

	var responseBody map[string]interface{}
	body, err := io.ReadAll(recorder.Result().Body)
	helper.PanicIfError(err)
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusBadRequest, int(responseBody["code"].(float64)))
	assert.Equal(t, "Bad Request", responseBody["status"])
}

func TestUpdateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, err := db.Begin()
	helper.PanicIfError(err)
	categoryRepository := repository.NewCategoryRepository()
	newCategory := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name": "Gadget Update"}`)
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(newCategory.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	router.ServeHTTP(recorder, request)

	var responseBody map[string]interface{}
	body, err := io.ReadAll(recorder.Result().Body)
	helper.PanicIfError(err)
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, newCategory.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Gadget Update", responseBody["data"].(map[string]interface{})["name"])
}

func TestUpdateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, err := db.Begin()
	helper.PanicIfError(err)
	categoryRepository := repository.NewCategoryRepository()
	newCategory := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name": ""}`)
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(newCategory.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	router.ServeHTTP(recorder, request)

	var responseBody map[string]interface{}
	body, err := io.ReadAll(recorder.Result().Body)
	helper.PanicIfError(err)
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusBadRequest, int(responseBody["code"].(float64)))
	assert.Equal(t, "Bad Request", responseBody["status"])
}

func TestGetOneCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, err := db.Begin()
	helper.PanicIfError(err)
	categoryRepository := repository.NewCategoryRepository()
	newCategory := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/"+strconv.Itoa(newCategory.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	router.ServeHTTP(recorder, request)

	var responseBody map[string]interface{}
	body, err := io.ReadAll(recorder.Result().Body)
	helper.PanicIfError(err)
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, newCategory.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, newCategory.Name, responseBody["data"].(map[string]interface{})["name"])
}

func TestGetOneCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	router := setupRouter(db)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/1", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	router.ServeHTTP(recorder, request)

	var responseBody map[string]interface{}
	body, err := io.ReadAll(recorder.Result().Body)
	helper.PanicIfError(err)
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusNotFound, int(responseBody["code"].(float64)))
	assert.Equal(t, "Not Found", responseBody["status"])
}

func TestDeleteCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, err := db.Begin()
	helper.PanicIfError(err)
	categoryRepository := repository.NewCategoryRepository()
	newCategory := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/"+strconv.Itoa(newCategory.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	router.ServeHTTP(recorder, request)

	var responseBody map[string]interface{}
	body, err := io.ReadAll(recorder.Result().Body)
	helper.PanicIfError(err)
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
}

func TestDeleteCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	router := setupRouter(db)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/1", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	router.ServeHTTP(recorder, request)

	var responseBody map[string]interface{}
	body, err := io.ReadAll(recorder.Result().Body)
	helper.PanicIfError(err)
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusNotFound, int(responseBody["code"].(float64)))
	assert.Equal(t, "Not Found", responseBody["status"])
}

func TestGetAllCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, err := db.Begin()
	helper.PanicIfError(err)
	categoryRepository := repository.NewCategoryRepository()
	newCategory := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	router.ServeHTTP(recorder, request)

	var responseBody map[string]interface{}
	body, err := io.ReadAll(recorder.Result().Body)
	helper.PanicIfError(err)
	json.Unmarshal(body, &responseBody)

	var categories = responseBody["data"].([]interface{})

	categoryResponse := categories[0].(map[string]interface{})

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, newCategory.Id, int(categoryResponse["id"].(float64)))
	assert.Equal(t, newCategory.Name, categoryResponse["name"])
}

func TestUnauthorized(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	router := setupRouter(db)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SALAH")

	router.ServeHTTP(recorder, request)

	var responseBody map[string]interface{}
	body, err := io.ReadAll(recorder.Result().Body)
	helper.PanicIfError(err)
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusUnauthorized, int(responseBody["code"].(float64)))
	assert.Equal(t, "Unauthorized", responseBody["status"])
}
