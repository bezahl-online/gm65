package api

import (
	"net/http"
	"testing"

	"github.com/deepmap/oapi-codegen/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func TestGetTest(t *testing.T) {
	result := testutil.NewRequest().Get("/test").Go(t, e)
	assert.Equal(t, http.StatusOK, result.Code())
}
