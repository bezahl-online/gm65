package api

import (
	"net/http"
	"testing"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func TestEnableCode(t *testing.T) {
	var codeType CodeType = CodeTypeEan13
	var request EnableCodeJSONRequestBody = EnableCodeJSONRequestBody{
		CodeType: &codeType,
	}
	result := testutil.NewRequest().Post("/enable_code").WithJsonBody(request).WithAcceptJson().Go(t, e)
	for result.Code() == http.StatusGone {
		time.Sleep(100 * time.Millisecond)
		result = testutil.NewRequest().Post("/enable_code").WithJsonBody(request).WithAcceptJson().Go(t, e)
	}
	assert.Equal(t, http.StatusOK, result.Code())
}
