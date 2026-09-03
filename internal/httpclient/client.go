package httpclient

import (
	"net/http"
	"time"
)

const DefaultTimeout = 10 * time.Second

type Options struct {
	Timeout             time.Duration
	MaxIdleConns        int
	MaxIdleConnsPerHost int
	InsecureSkipVerify  bool
}

func New(opts Options) *http.Client {
	if opts.Timeout == 0 {
		opts.Timeout = DefaultTimeout
	}

	if opts.MaxIdleConns == 0 {
		opts.MaxIdleConns = 100
	}

	if opts.MaxIdleConnsPerHost == 0 {
		opts.MaxIdleConnsPerHost = 20
	}

	transport := &http.Transport{
		MaxIdleConns:        opts.MaxIdleConns,
		MaxIdleConnsPerHost: opts.MaxIdleConnsPerHost,
	}

	return &http.Client{
		Timeout:   opts.Timeout,
		Transport: transport,

		// Do not automatically follow redirects.
		CheckRedirect: func(
			req *http.Request,
			via []*http.Request,
		) error {
			return http.ErrUseLastResponse
		},
	}
}
