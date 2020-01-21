package controllers_test

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jonatascabral/jokes-app/pkg/models"
	"github.com/jonatascabral/jokes-app/pkg/routes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoot(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := routes.LoadRoutes()
	body := gin.H{
		"message": "pong",
	}
	response := performRequest(router, http.MethodGet, "/api/")
	var responseJson map[string]string
	err := json.Unmarshal([]byte(response.Body.String()), &responseJson)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, body["message"], responseJson["message"])
}

func TestGetJokes(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := routes.LoadRoutes()
	response := performRequest(router, http.MethodGet, "/api/jokes/")
	var responseJson []models.Joke
	err := json.Unmarshal([]byte(response.Body.String()), &responseJson)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Greater(t, len(responseJson), 0)
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}