package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	count := 5
	city := "moscow"
	url := fmt.Sprintf("/cafe?count=%d&city=%s", count, city)
	req := httptest.NewRequest("GET", url, nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	require.Equal(t, http.StatusOK, responseRecorder.Code,
		"expected http status %d got %d", http.StatusOK, responseRecorder.Code)

	responseCount := len(strings.Split(responseRecorder.Body.String(), ","))
	assert.Equal(t, totalCount, responseCount,
		"expected length %d got %d", totalCount, responseCount)
}

func TestMainHandlerWhenCountWrong(t *testing.T) {
	count := "one"
	city := "moscow"
	url := fmt.Sprintf("/cafe?count=%s&city=%s", count, city)
	req := httptest.NewRequest("GET", url, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code,
		"expected http status %d got %d", http.StatusBadRequest, responseRecorder.Code)

	expectedError := "wrong count value"
	assert.Equal(t, expectedError, responseRecorder.Body.String(),
		"expected body %s got %s", expectedError, responseRecorder.Body.String())
}

func TestMainHandler(t *testing.T) {
	count := 4
	city := "moscow"
	url := fmt.Sprintf("/cafe?count=%d&city=%s", count, city)

	req := httptest.NewRequest("GET", url, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code,
		"expected http status %d got %d", http.StatusBadRequest, responseRecorder.Code)

	require.NotNil(t, responseRecorder.Body, "expected not empty body")

	expectedBody := "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент"
	assert.Equal(t, expectedBody, responseRecorder.Body.String(),
		"expected body %s got %s", expectedBody, responseRecorder.Body.String())
}

func TestMainHandlerWhenCityNotFound(t *testing.T) {
	city := "test"
	count := 1
	url := fmt.Sprintf("/cafe?count=%d&city=%s", count, city)
	req := httptest.NewRequest("GET", url, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code,
		"expected http status %d got %d", http.StatusBadRequest, responseRecorder.Code)

	expectedError := "wrong city value"
	assert.Equal(t, expectedError, responseRecorder.Body.String(),
		"expected body %s got %s", expectedError, responseRecorder.Body.String())
}
