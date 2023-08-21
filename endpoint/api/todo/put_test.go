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

func Test_UpdateData(t *testing.T) {

	ctrl := gomock.NewController(t)
	mockClient := queries.NewMockQueries(ctrl)

	ctx := context.TODO()
	requestPayload := model.Todos{
		Title:    "Work",
		Status:   "Done",
		Priority: "High",
		UserID:   1,
	}
	body, _ := json.Marshal(requestPayload)
	request, _ := http.NewRequestWithContext(ctx, http.MethodPost, "/", bytes.NewBuffer(body))
	request = mux.SetURLVars(request, map[string]string{"id": "1"})
	responseWriter := httptest.NewRecorder()

	mockClient.EXPECT().UpdateTodoData(1, requestPayload).Return(nil)
	UpdateHandler(mockClient)(responseWriter, request)

	resp := responseWriter.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
