package httprunner

import (
	"testing"
)

func Test_Equals_JSON(t *testing.T) {
	data1 := `{"a":"a","b":"b"}`
	data2 := `{
		"a": "a",
		"b": "b"
	}`

	jsonEquals(t, data1, data2)
}

func Test_Equals_JSON_EmptyArray(t *testing.T) {
	data1 := `{"data":{}}`
	data2 := `{"data": {}}`

	jsonEquals(t, data1, data2)
}

func Test_Equals_XML(t *testing.T) {
	data1 := `<r><a>a</a><b>b</b></r>`
	data2 := `<r>
		<a>a</a>
		<b>b</b>
	</r>`

	xmlEquals(t, data1, data2)
}
