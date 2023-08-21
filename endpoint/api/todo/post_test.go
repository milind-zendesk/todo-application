package todo

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-application/endpoint/api/todo/queries"
	"todo-application/model"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockClient := queries.NewMockQueries(ctrl)

	ctx := context.TODO()
	requestPayload := model.Todos{
		Title:    "Work",
		Status:   "not started",
		Priority: "high",
		UserID:   1,
	}
	body, _ := json.Marshal(requestPayload)
	request, _ := http.NewRequestWithContext(ctx, http.MethodPost, "/", bytes.NewBuffer(body))
	responseWriter := httptest.NewRecorder()

	mockClient.EXPECT().InsertTodoData(requestPayload).Return(nil)
	CreateHandler(mockClient)(responseWriter, request)

	resp := responseWriter.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
