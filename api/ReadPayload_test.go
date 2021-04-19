package api

import (
	"net/http"
	"testing"

	"github.com/deepmap/oapi-codegen/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func TestReadPayload(t *testing.T) {
	var read *Read = &Read{}
	testCode := "7613034471888"
	scanner.SetCode(testCode)
	result := testutil.NewRequest().Get("/read").Go(t, e)
	if assert.Equal(t, http.StatusOK, result.Code()) {
		err := result.UnmarshalBodyToObject(read)
		if assert.NoError(t, err) {
			assert.Equal(t, testCode, *read.Payload)
		}
	}
}
