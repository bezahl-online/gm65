package api

import (
	"net/http"
	"testing"

	"github.com/deepmap/oapi-codegen/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func TestLight(t *testing.T) {
	var lightSwitch SwitchOpt = SwitchOpt_std
	var request LightJSONRequestBody = LightJSONRequestBody{
		Set: &lightSwitch,
	}
	result := testutil.NewRequest().Post("/light").WithJsonBody(request).WithAcceptJson().Go(t, e)
	assert.Equal(t, http.StatusOK, result.Code())
}
