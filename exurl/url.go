package exurl

import (
	"net/url"
	"strings"
)

type URL struct {
	pathParts []string
	path      string
	suffix    string
	fragment  string

	url *url.URL
	Raw string
}

func Parse(rawURL string) (*URL, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	ret := &URL{
		url: u,
		Raw: rawURL,
	}

	return ret, nil
}

func (u *URL) Unwrap() *url.URL {
	return u.url
}

func (u *URL) GetScheme() string {
	return u.url.Scheme
}

// GetHost returns host or host:port
func (u *URL) GetHost() string {
	return u.url.Host
}

// GetHostname returns u.GetHost(), stripping any valid port number if present.
//
// If the result is enclosed in square brackets, as literal IPv6 addresses are,
// the square brackets are removed from the result.
func (u *URL) GetHostname() string {
	return u.url.Hostname()
}

// GetPort returns the port part of u.GetHost(), without the leading colon.
//
// If u.Host doesn't contain a valid numeric port, Port returns an empty string.
func (u *URL) GetPort() string {
	return u.url.Port()
}

func (u *URL) GetUser() *url.Userinfo {
	return u.url.User
}

func (u *URL) GetPathParts() []string {
	if u.pathParts != nil {
		return u.pathParts
	}

	var pathParts []string
	if u.url.Path != "" || u.url.RawPath != "" {
		rp := u.url.RawPath
		if rp == "" {
			rp = u.url.Path
		}

		rp = strings.TrimPrefix(rp, "/")

		if rp != "" {
			parts := strings.Split(rp, "/")
			pathParts = make([]string, len(parts))
			for i, part := range parts {
				var ep string
				if uep, err := url.PathUnescape(part); err == nil {
					ep = uep
				} else {
					ep = part // cannot decode, ignore
				}

				pathParts[i] = ep
			}
		} else {
			pathParts = []string{}
		}
	} else {
		pathParts = []string{}
	}

	u.pathParts = pathParts
	return u.pathParts
}

func (u *URL) GetPath() string {
	if u.path != "" {
		return u.path
	}

	pathParts := u.GetPathParts()
	if len(pathParts) == 0 {
		return ""
	}

	rb := strings.Builder{}

	for i, part := range pathParts {
		if i != 0 {
			rb.WriteByte('/')
		}

		rb.WriteString(url.PathEscape(part))
	}

	u.path = rb.String()
	return u.path
}

func (u *URL) GetPathLastPart() string {
	parts := u.GetPathParts()
	if len(parts) == 0 {
		return ""
	}

	return parts[len(parts)-1]
}

func (u *URL) GetSuffix() string {
	if u.suffix == "" {
		lastPart := u.GetPathLastPart()
		if lastPart == "" {
			return ""
		}

		lastIndex := strings.LastIndex(lastPart, ".")
		if lastIndex == -1 {
			return ""
		}
	}

	return u.suffix
}

// GetRawQuery returns encoded query values, without '?'
func (u *URL) GetRawQuery() string {
	return u.url.RawQuery
}

// GetQueryString returns encoded query values, without '?'
//
// TODO: please use GetRawQuery now
func (u *URL) GetQueryString() string {
	return u.url.RawQuery
}

// GetQuery parses RawQuery and returns the corresponding values.
// It silently discards malformed value pairs.
// To check errors use ParseQuery.
func (u *URL) GetQuery() url.Values {
	return u.url.Query()
}

// GetFragment returns the fragment without '#'
func (u *URL) GetFragment() string {
	if u.fragment != "" {
		return u.fragment
	}

	fragment := ""
	if u.url.Fragment != "" {
		oldRawFragment := u.url.RawFragment
		u.url.RawFragment = "" // <- 强制触发 re-escape
		fragment = u.url.EscapedFragment()
		u.url.RawFragment = oldRawFragment
	}
	u.fragment = fragment

	return fragment
}
