package httprunner

import (
	"encoding/json"
	"encoding/xml"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func jsonEquals(t *testing.T, expected, actual string, msgAndArgs ...interface{}) {
	var expectedJSONAsInterface, actualJSONAsInterface interface{}
	var flattenExpected = []byte(flattenString(expected))
	var flattenActual = []byte(flattenString(actual))

	if err := json.Unmarshal(flattenExpected, &expectedJSONAsInterface); err != nil {
		t.Errorf("cannot decode expected JSON string: %s", err.Error())
	}
	if err := json.Unmarshal(flattenActual, &actualJSONAsInterface); err != nil {
		t.Errorf("cannot decode actual JSON string: %s", err.Error())
	}

	require.Equal(t, expectedJSONAsInterface, actualJSONAsInterface, msgAndArgs...)
}

func xmlEquals(t *testing.T, expected, actual string, msgAndArgs ...interface{}) {
	var expectedJSONAsInterface, actualJSONAsInterface interface{}
	var flattenExpected = []byte(flattenString(expected))
	var flattenActual = []byte(flattenString(actual))

	if err := xml.Unmarshal(flattenExpected, &expectedJSONAsInterface); err != nil {
		t.Errorf("cannot decode expected JSON string: %s", err.Error())
	}
	if err := xml.Unmarshal(flattenActual, &actualJSONAsInterface); err != nil {
		t.Errorf("cannot decode actual JSON string: %s", err.Error())
	}

	require.Equal(t, expectedJSONAsInterface, actualJSONAsInterface, msgAndArgs...)
}

func flattenString(text string) string {
	return strings.Map(func(r rune) rune {
		if r == '\n' || r == '\t' {
			return -1
		}
		return r
	}, text)
}
