package httpclient

import (
	"net/http"
	"testing"
	"time"
)

func TestNew_Defaults(t *testing.T) {
	client := New(Options{})

	tests := []struct {
		name string
		got  any
		want any
	}{
		{
			name: "default timeout",
			got:  client.Timeout,
			want: DefaultTimeout,
		},
		{
			name: "transport configured",
			got:  client.Transport != nil,
			want: true,
		},
		{
			name: "redirect policy configured",
			got:  client.CheckRedirect != nil,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Fatalf("got %v, want %v", tt.got, tt.want)
			}
		})
	}
}

func TestNew_DoesNotFollowRedirects(t *testing.T) {
	client := New(Options{
		Timeout: 5 * time.Second,
	})

	req, err := http.NewRequest(
		http.MethodGet,
		"https://example.com",
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	err = client.CheckRedirect(req, nil)

	if err != http.ErrUseLastResponse {
		t.Fatalf(
			"redirect error = %v, want %v",
			err,
			http.ErrUseLastResponse,
		)
	}
}
