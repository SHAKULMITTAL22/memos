package v1

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// TestRegisterIdentityProviderRoutes_f046fed06a tests the registerIdentityProviderRoutes method
func TestRegisterIdentityProviderRoutes_f046fed06a(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	s := &APIV1Service{}

	// Test case: GetIdentityProviderList
	req = httptest.NewRequest(http.MethodGet, "/idp", nil)
	c = e.NewContext(req, rec)
	s.registerIdentityProviderRoutes(e.Group("/"))
	if assert.NoError(t, s.GetIdentityProviderList(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		t.Log("Success: GetIdentityProviderList")
	} else {
		t.Error("Failure: GetIdentityProviderList, Arguments: ", c)
	}

	// Test case: CreateIdentityProvider
	req = httptest.NewRequest(http.MethodPost, "/idp", nil)
	c = e.NewContext(req, rec)
	s.registerIdentityProviderRoutes(e.Group("/"))
	if assert.NoError(t, s.CreateIdentityProvider(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		t.Log("Success: CreateIdentityProvider")
	} else {
		t.Error("Failure: CreateIdentityProvider, Arguments: ", c)
	}

	// TODO: Add more test cases for GetIdentityProvider, UpdateIdentityProvider, DeleteIdentityProvider
}
