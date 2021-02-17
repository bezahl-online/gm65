package api

import (
	"net/http"
	"testing"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func TestLight(t *testing.T) {
	var lightSwitch SwitchOpt = SwitchOpt_std
	var request LightJSONRequestBody = LightJSONRequestBody{
		Set: &lightSwitch,
	}
	result := testutil.NewRequest().Post("/light").WithJsonBody(request).WithAcceptJson().Go(t, e)
	for result.Code()==http.StatusGone {
		time.Sleep(100*time.Millisecond)
		result = testutil.NewRequest().Post("/light").WithJsonBody(request).WithAcceptJson().Go(t, e)
	}
	assert.Equal(t, http.StatusOK, result.Code())
}
