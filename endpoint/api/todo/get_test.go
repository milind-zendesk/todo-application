package todo

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-application/endpoint/api/todo/queries"
	"todo-application/model"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_GetData(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockClient := queries.NewMockQueries(ctrl)

	ctx := context.TODO()
	request, _ := http.NewRequestWithContext(ctx, http.MethodPost, "/", nil)
	responseWriter := httptest.NewRecorder()

	mockClient.EXPECT().GetAllTodosData().Return([]model.Todos{
		{
			Id:       1,
			Title:    "Work",
			Status:   "Done",
			Priority: "High",
			UserID:   3,
		},
		{
			Id:       2,
			Title:    "Breakfast",
			Status:   "Pending",
			Priority: "High",
			UserID:   4,
		},
		{
			Id:       3,
			Title:    "Exercise",
			Status:   "Ongoing",
			Priority: "High",
			UserID:   5,
		},
	}, nil)
	GetAllHandlers(mockClient)(responseWriter, request)

	resp := responseWriter.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	body, _ := io.ReadAll(resp.Body)
	assert.Contains(t, string(body), "Breakfast")
}

func Test_GetDataReturnsErrors(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockClient := queries.NewMockQueries(ctrl)

	ctx := context.TODO()
	request, _ := http.NewRequestWithContext(ctx, http.MethodPost, "/", nil)
	responseWriter := httptest.NewRecorder()

	mockClient.EXPECT().GetAllTodosData().Return([]model.Todos{}, errors.New("Something Broke"))
	GetAllHandlers(mockClient)(responseWriter, request)

	resp := responseWriter.Result()
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}
