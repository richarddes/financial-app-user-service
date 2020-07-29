package handler_test

import (
	"net/http"
	"testing"
	"time"
	"user-service/config"
	"user-service/handler"
)

func TestHandleDeleteAcc(t *testing.T) {
	mockEnv := config.NewMockEnv()

	rr := handler.NewRecorder(t, "POST", "/api/users/delete-acc", nil, handler.HandleDeleteAcc(mockEnv))

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d but gor %d instead", http.StatusOK, rr.Code)
	}

	authCookie := rr.Result().Cookies()[0]

	if authCookie == nil {
		t.Errorf("The authentication cookie hasn't been changed")
	}

	if !authCookie.Expires.Equal(time.Unix(0, 0)) {
		t.Errorf("Expected authentication cookie to expire at Unix(0,0) but got %v instead", authCookie.Expires)
	}
}
