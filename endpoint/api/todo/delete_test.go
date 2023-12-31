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
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Test_deleteData(t *testing.T) {

	ctrl := gomock.NewController(t)
	mockClient := queries.NewMockQueries(ctrl)

	ctx := context.TODO()
	requestPayload := model.User{
		Id:       1,
		Name:     "Milind Shinde",
		Location: "Pune",
	}
	body, _ := json.Marshal(requestPayload)
	request, _ := http.NewRequestWithContext(ctx, http.MethodPost, "/", bytes.NewBuffer(body))
	request = mux.SetURLVars(request, map[string]string{"id": "1"})
	responseWriter := httptest.NewRecorder()

	mockClient.EXPECT().DeleteTodoData(1).Return(nil)
	DeleteHandler(mockClient)(responseWriter, request)

	resp := responseWriter.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

}
