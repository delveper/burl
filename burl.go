// Package burl  built is meant to build URL with functional options.
package burl

import (
	"fmt"
	"net/url"
	"path"
)

const DefaultScheme = "https"

type Option = func(*url.URL)

// New makes URL for all needs - [scheme:][//[userinfo@]host][/]path[?query][#fragment]
func New(options ...Option) (u *url.URL) {
	u.Scheme = DefaultScheme

	for i := range options {
		options[i](u)
	}

	return u
}

func WithScheme(scheme string) Option {
	return func(u *url.URL) {
		u.Scheme = scheme
	}
}

func WithUser(username string) Option {
	return func(u *url.URL) {
		u.User = url.User(username)
	}
}

func WithUserPassword(username, password string) Option {
	return func(u *url.URL) {
		u.User = url.UserPassword(username, password)
	}
}

func WithHost(host string) Option {
	return func(u *url.URL) {
		u.Host = host
	}
}

func WithPath(paths ...string) Option {
	return func(u *url.URL) {
		u.Path = path.Join(paths...)
	}
}

func WithValue(key string, value any) Option {
	return func(u *url.URL) {
		values := u.Query()
		values.Add(key, fmt.Sprint(value))
		u.RawQuery = values.Encode()
	}
}
