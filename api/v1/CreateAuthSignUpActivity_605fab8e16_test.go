package v1

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/usememos/memos/store"
)

// TestCreateAuthSignUpActivity_605fab8e16 tests the createAuthSignUpActivity method
func TestCreateAuthSignUpActivity_605fab8e16(t *testing.T) {
	e := echo.New()
	s := &APIV1Service{
		Store: new(MockStore),
	}
	user := &store.User{
		ID:       "test-id",
		Username: "test-user",
	}
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
	c := e.NewContext(req, httptest.NewRecorder())
	c.Set("user", user)

	t.Run("Successful creation", func(t *testing.T) {
		err := s.createAuthSignUpActivity(c, user)
		if err != nil {
			t.Errorf("Failed to create activity: %v", err)
		} else {
			t.Log("Success: Activity created")
		}
	})

	t.Run("Failure in marshalling payload", func(t *testing.T) {
		badUser := &store.User{
			ID:       "test-id",
			Username: string([]byte{0x80, 0x81, 0x82}),
		}
		err := s.createAuthSignUpActivity(c, badUser)
		if err == nil || !errors.Is(err, json.ErrUnsupportedValue) {
			t.Errorf("Expected json.ErrUnsupportedValue, got: %v", err)
		} else {
			t.Log("Success: Correct error returned when marshalling payload fails")
		}
	})

	t.Run("Failure in creating activity", func(t *testing.T) {
		s.Store = &MockStore{
			CreateActivityErr: errors.New("test error"),
		}
		err := s.createAuthSignUpActivity(c, user)
		if err == nil || !strings.Contains(err.Error(), "failed to create activity") {
			t.Errorf("Expected 'failed to create activity' error, got: %v", err)
		} else {
			t.Log("Success: Correct error returned when creating activity fails")
		}
	})
}

type MockStore struct {
	CreateActivityErr error
}

func (m *MockStore) CreateActivity(ctx context.Context, activity *store.Activity) (*store.Activity, error) {
	if m.CreateActivityErr != nil {
		return nil, m.CreateActivityErr
	}
	return activity, nil
}
