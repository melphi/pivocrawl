package testit

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	// Given
	ctx := StartTestContext()
	req := httptest.NewRequest("GET", "/health", nil)

	// When
	resp, _ := ctx.API.Test(req, 1)

	// Then
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
