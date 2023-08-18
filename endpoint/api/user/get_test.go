package user

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	userqueries "todo-application/endpoint/api/user/userQueries"
	"todo-application/model"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllData(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockClient := userqueries.NewMockUserQueries(ctrl)

	ctx := context.TODO()
	request, _ := http.NewRequestWithContext(ctx, http.MethodPost, "/", nil)
	responseWriter := httptest.NewRecorder()

	expected_result := []model.User{
		{
			Id:       1,
			Name:     "Milind Shinde",
			Location: "Pune",
		},
		{
			Id:       2,
			Name:     "Devesh Chinchole",
			Location: "Jalgaon",
		},
		{
			Id:       3,
			Name:     "Luke Josh",
			Location: "Melbourn",
		},
	}

	mockClient.EXPECT().GetAllUsersData().Return(expected_result, nil)
	GetAllHandler(mockClient)(responseWriter, request)

	resp := responseWriter.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	actual_result := []model.User{}
	err := json.Unmarshal(body, &actual_result)
	assert.NoError(t, err)
	assert.Equal(t, expected_result, actual_result)
}

func Test_GetData(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockClient := userqueries.NewMockUserQueries(ctrl)

	ctx := context.TODO()
	request, _ := http.NewRequestWithContext(ctx, http.MethodGet, "/1", nil)
	request = mux.SetURLVars(request, map[string]string{"id": "1"})
	responseWriter := httptest.NewRecorder()

	expected_result := model.User{
		Id:       1,
		Name:     "Milind Shinde",
		Location: "Pune",
	}
	mockClient.EXPECT().GetUserData(1).Return(expected_result, nil)
	GetHandler(mockClient)(responseWriter, request)

	resp := responseWriter.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	actual_result := model.User{}
	err := json.Unmarshal(body, &actual_result)
	assert.NoError(t, err)

	assert.Equal(t, expected_result, actual_result)
}

func Test_GetUserTodoData(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockClient := userqueries.NewMockUserQueries(ctrl)

	ctx := context.TODO()
	request, _ := http.NewRequestWithContext(ctx, http.MethodGet, "/", nil)
	request = mux.SetURLVars(request, map[string]string{"id": "1"})
	responseWriter := httptest.NewRecorder()

	expected_result := model.UserTodoDetails{
		Id:       1,
		Name:     "Milind Shinde",
		Location: "Pune",
		Todos: []model.Todos{
			{
				Id:       47,
				Title:    "Breakfast",
				Status:   "Done",
				Priority: "medium",
				UserID:   1,
			},
			{
				Id:       48,
				Title:    "Work",
				Status:   "in progress",
				Priority: "low",
				UserID:   1,
			},
			{
				Id:       49,
				Title:    "Shopping",
				Status:   "not started",
				Priority: "high",
				UserID:   1,
			},
		},
		TotalCount: 3,
		Priorities: map[string]int{
			"high":   1,
			"low":    1,
			"medium": 1,
		},
	}
	mockClient.EXPECT().GetUserTodosData(1).Return(expected_result, nil)
	GetUserTodosHandler(mockClient)(responseWriter, request)

	resp := responseWriter.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	actual_result := model.UserTodoDetails{}
	err := json.Unmarshal(body, &actual_result)
	assert.NoError(t, err)

	assert.Equal(t, expected_result, actual_result)
}
