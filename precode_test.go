package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenOk(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	body := responseRecorder.Body.String()

	require.NotEmpty(t, body)
	cafeList := strings.Split(body, ",")
	assert.Len(t, cafeList, 2)
}

func TestMainHandlerWhenCityIsWrong(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=2&city=tula", nil)
	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	require.Equal(t, "wrong city value", responseRecorder.Body.String())
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {

	totalCount := 4
	requestedCount := 10

	// Создаём запрос к сервису
	req := httptest.NewRequest("GET", "/cafe?count="+strconv.Itoa(requestedCount)+"&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	require.NotEmpty(t, responseRecorder.Body.String())
	cafeList := strings.Split(responseRecorder.Body.String(), ",")
	require.Len(t, cafeList, totalCount)

}
