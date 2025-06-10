package auth

import (
	"testing"
	"net/http"
	"reflect"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		input http.Header
		want string
	}{
		{
			input: http.Header{
				"Authorization": {"ApiKey 123"},
			},
			want: "123",
		},
		{
			input: http.Header{
				"Authorization": {""},
			},
			want: "",
		},
	}

	for _, tc := range tests {
		got, _ := GetAPIKey(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("Expected: %v, Got: %v", tc.want, got)
		}
	}
}
