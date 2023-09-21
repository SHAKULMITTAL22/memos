package v1

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// TestGetImage_a8728ef6e5 tests the GetImage function
func TestGetImage_a8728ef6e5(t *testing.T) {
	e := echo.New()

	t.Run("valid url", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/?url=https://example.com/image.jpg", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := GetImage(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		t.Log("Test case passed for valid url")
	})

	t.Run("empty url", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/?url=", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := GetImage(c)

		assert.Error(t, err)
		he, ok := err.(*echo.HTTPError)
		if ok {
			assert.Equal(t, http.StatusBadRequest, he.Code)
			assert.Equal(t, "Missing image url", he.Message)
		} else {
			t.Error(fmt.Sprintf("Expected *echo.HTTPError, got %T", err))
		}
		t.Log("Test case passed for empty url")
	})

	t.Run("invalid url", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/?url=invalid", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := GetImage(c)

		assert.Error(t, err)
		he, ok := err.(*echo.HTTPError)
		if ok {
			assert.Equal(t, http.StatusBadRequest, he.Code)
			assert.Equal(t, "Wrong url", he.Message)
		} else {
			t.Error(fmt.Sprintf("Expected *echo.HTTPError, got %T", err))
		}
		t.Log("Test case passed for invalid url")
	})

	t.Run("non-existent url", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/?url=https://example.com/nonexistent.jpg", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := GetImage(c)

		assert.Error(t, err)
		he, ok := err.(*echo.HTTPError)
		if ok {
			assert.Equal(t, http.StatusBadRequest, he.Code)
			assert.Equal(t, fmt.Sprintf("Failed to get image url: %s", url.QueryEscape("https://example.com/nonexistent.jpg")), he.Message)
		} else {
			t.Error(fmt.Sprintf("Expected *echo.HTTPError, got %T", err))
		}
		t.Log("Test case passed for non-existent url")
	})
}
