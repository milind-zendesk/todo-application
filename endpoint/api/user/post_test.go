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
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockClient := userqueries.NewMockUserQueries(ctrl)

	ctx := context.TODO()
	requestPayload := model.User{
		Name:     "Milind Shinde",
		Location: "Pune",
	}
	body, _ := json.Marshal(requestPayload)
	request, _ := http.NewRequestWithContext(ctx, http.MethodPost, "/", bytes.NewBuffer(body))
	responseWriter := httptest.NewRecorder()

	mockClient.EXPECT().InsertUserData(requestPayload).Return(nil)
	CreateHandler(mockClient)(responseWriter, request)

	resp := responseWriter.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
