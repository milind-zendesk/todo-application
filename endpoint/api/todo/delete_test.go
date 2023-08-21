package todo

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-application/endpoint/api/todo/queries"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestDeleteTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockClient := queries.NewMockQueries(ctrl)

	ctx := context.TODO()
	request, _ := http.NewRequestWithContext(ctx, http.MethodDelete, "/", nil)
	request = mux.SetURLVars(request, map[string]string{"id": "1"})
	responseWriter := httptest.NewRecorder()

	mockClient.EXPECT().DeleteTodoData(1).Return(nil)
	DeleteHandler(mockClient)(responseWriter, request)

	resp := responseWriter.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
