package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIServer_HandleHello(t *testing.T) {
	s := New((NewConfig()))
	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/hello", nil)
	assert.Equal(t, nil, err)

	s.handleHello().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Hello")
}
