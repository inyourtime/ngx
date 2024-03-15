package restful

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppIndex(t *testing.T) {
	setup()

	NewAppAPIHandler(e).Init()

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	resp, err := e.Test(req)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
