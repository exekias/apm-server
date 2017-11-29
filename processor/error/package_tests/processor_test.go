package package_tests

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	er "github.com/elastic/apm-server/processor/error"
	"github.com/elastic/apm-server/tests"
)

// ensure all valid documents pass through the whole validation and transformation process
func TestProcessorOK(t *testing.T) {
	requestInfo := []tests.RequestInfo{
		{Name: "TestProcessErrorMininmalPayloadException", Path: "tests/data/valid/error/minimal_payload_exception.json"},
		{Name: "TestProcessErrorMininmalPayloadLog", Path: "tests/data/valid/error/minimal_payload_log.json"},
		{Name: "TestProcessErrorMininmalApp", Path: "tests/data/valid/error/minimal_app.json"},
		{Name: "TestProcessErrorFull", Path: "tests/data/valid/error/payload.json"},
		{Name: "TestProcessErrorNullValues", Path: "tests/data/valid/error/null_values.json"},
	}

	req := httptest.NewRequest("GET", "/", nil)
	tests.TestProcessRequests(t, er.NewProcessor(req), requestInfo, map[string]string{})
}

// ensure invalid documents fail the json schema validation already
func TestProcessorFailedValidation(t *testing.T) {
	data, err := tests.LoadInvalidData("error")
	assert.Nil(t, err)
	req := httptest.NewRequest("GET", "/", nil)
	err = er.NewProcessor(req).Validate(data)
	assert.NotNil(t, err)
}
