package user

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	userqueries "todo-application/endpoint/api/user/userQueries"
	"todo-application/model"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockClient := userqueries.NewMockUserQueries(ctrl)

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

	mockClient.EXPECT().UpdateUserData(1, requestPayload).Return(nil)
	UpdateHandler(mockClient)(responseWriter, request)

	resp := responseWriter.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
