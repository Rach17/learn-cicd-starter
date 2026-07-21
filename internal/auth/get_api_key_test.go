package auth

import (
	"errors"
	"net/http"
	"testing"
)

// func TestSplit(t *testing.T) {
//     tests := map[string]struct {
//         input string
//         sep   string
//         want  []string
//     }{
//         "simple":       {input: "a/b/c", sep: "/", want: []string{"a", "b", "c"}},
//         "wrong sep":    {input: "a/b/c", sep: ",", want: []string{"a/b/c"}},
//         "no sep":       {input: "abc", sep: "/", want: []string{"abc"}},
//         "trailing sep": {input: "a/b/c/", sep: "/", want: []string{"a", "b", "c"}},
//     }

//     for name, tc := range tests {
//         t.Run(name, func(t *testing.T) {
//             got := Split(tc.input, tc.sep)
//             diff := cmp.Diff(tc.want, got)
//             if diff != "" {
//                 t.Fatalf(diff)
//             }
//         })
//     }
// }

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input http.Header
		want  struct {
			apiKey string
			err    error
		}
	}{
		"no auth header": {
			input: http.Header{},
			want: struct {
				apiKey string
				err    error
			}{
				apiKey: "",
				err:    ErrNoAuthHeaderIncluded,
			},
		},
		"invalid auth header": {
			input: http.Header{
				"Authorization": []string{"Bearer"},
			},
			want: struct {
				apiKey string
				err    error
			}{
				apiKey: "",
				err:    errors.New("malformed authorization header"),
			},
		},
		"valid auth header": {
			input: http.Header{
				"Authorization": []string{"ApiKey 123"},
			},
			want: struct {
				apiKey string
				err    error
			}{
				apiKey: "123",
				err:    nil,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			gotAPIKey, gotErr := GetAPIKey(tc.input)
			if gotAPIKey != tc.want.apiKey {
				t.Fatalf("GetAPIKey(%v) = (%v, %v), want (%v, %v)", tc.input, gotAPIKey, gotErr, tc.want.apiKey, tc.want.err)
			}
			if tc.want.err != nil {
				if gotErr == nil || gotErr.Error() != tc.want.err.Error() {
					t.Fatalf("GetAPIKey(%v) = (%v, %v), want (%v, %v)", tc.input, gotAPIKey, gotErr, tc.want.apiKey, tc.want.err)
				}
			} else if gotErr != nil {
				t.Fatalf("GetAPIKey(%v) = (%v, %v), want (%v, %v)", tc.input, gotAPIKey, gotErr, tc.want.apiKey, tc.want.err)
			}
		})
	}
}
