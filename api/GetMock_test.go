package api

import (
	"net/http"
	"testing"

	"github.com/deepmap/oapi-codegen/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func TestGetMock(t *testing.T) {
	result := testutil.NewRequest().Get("/mock?code=1234567890123").Go(t, e)
	assert.Equal(t, http.StatusOK, result.Code())
}
