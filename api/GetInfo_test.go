package api

import (
	"net/http"
	"testing"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func TestGetInfo(t *testing.T) {
	result := testutil.NewRequest().Get("/info").Go(t, e)
	for result.Code()==http.StatusGone {
		time.Sleep(100*time.Millisecond)
		result = testutil.NewRequest().Get("/info").Go(t, e)
	}
	if assert.Equal(t, http.StatusOK, result.Code()) {
		var scannerInfo *ScannerInfo = &ScannerInfo{}
		err := result.UnmarshalBodyToObject(&scannerInfo)
		if assert.NoError(t, err) {
			assert.True(t, len(*scannerInfo.HardwareVersion) > 1)
		}
	}
}
