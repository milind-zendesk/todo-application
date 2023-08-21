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

func Test_PostData(t *testing.T) {

	ctrl := gomock.NewController(t)
	ctx := context.TODO()
	mockClient := queries.NewMockQueries(ctrl)

	requestPayload := model.Todos{
		Title:    "Work",
		Status:   "Done",
		Priority: "High",
		UserID:   3,
	}

	body, _ := json.Marshal(requestPayload)
	request, _ := http.NewRequestWithContext(ctx, http.MethodPost, "/", bytes.NewBuffer(body))
	responseWriter := httptest.NewRecorder()
	mockClient.EXPECT().InsertTodoData(requestPayload).Return(nil)

	// Call
	CreateHandler(mockClient)(responseWriter, request)

	// Check
	resp := responseWriter.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

}
