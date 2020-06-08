package utils

import (
	"net/url"
)

type GrafanaURL struct {
	DashboardURL string
	ExternalURL  string
	Username     string
	Password     string
}

func GrafanaDashboardToInternalURL(dashboard_url, username, password string) string {
	res := GetGrafanaDashboardURL(&GrafanaURL{
		DashboardURL: dashboard_url,
		Username:     username,
		Password:     password,
	})

	return res
}

func GrafanaDashboardToExternalURL(dashboard_url, external_url string) string {
	res := GetGrafanaDashboardURL(&GrafanaURL{
		DashboardURL: dashboard_url,
		ExternalURL:  external_url,
	})

	return res
}

func GrafanaDashboardExternalToInternalURL(dashboard_url, internal_url, username, password string) string {
	res := GetGrafanaDashboardURL(&GrafanaURL{
		DashboardURL: dashboard_url,
		ExternalURL:  internal_url,
		Username:     username,
		Password:     password,
	})

	return res
}

// GetGrafanaDashboardURL ...
func GetGrafanaDashboardURL(payload *GrafanaURL) string {
	_dashboard_url := GetParsedURL(payload.DashboardURL)
	_external_url := GetParsedURL(payload.ExternalURL)

	res := _dashboard_url
	if payload.ExternalURL != "" {
		if payload.Username != "" && payload.Password == "" {
			res = &url.URL{
				Scheme:     _external_url.Scheme,
				Opaque:     _external_url.Opaque,
				User:       url.User(payload.Username),
				Host:       _external_url.Host,
				Path:       _external_url.Path + _dashboard_url.Path,
				RawPath:    _dashboard_url.RawPath,
				ForceQuery: _dashboard_url.ForceQuery,
				RawQuery:   _dashboard_url.RawQuery,
				Fragment:   _dashboard_url.Fragment,
			}
			return res.String()
		} else if payload.Username != "" && payload.Password != "" {
			res = &url.URL{
				Scheme:     _external_url.Scheme,
				Opaque:     _external_url.Opaque,
				User:       url.UserPassword(payload.Username, payload.Password),
				Host:       _external_url.Host,
				Path:       _external_url.Path + _dashboard_url.Path,
				RawPath:    _dashboard_url.RawPath,
				ForceQuery: _dashboard_url.ForceQuery,
				RawQuery:   _dashboard_url.RawQuery,
				Fragment:   _dashboard_url.Fragment,
			}
			return res.String()
		} else {
			res = &url.URL{
				Scheme:     _external_url.Scheme,
				Opaque:     _external_url.Opaque,
				User:       nil,
				Host:       _external_url.Host,
				Path:       _external_url.Path + _dashboard_url.Path,
				RawPath:    _dashboard_url.RawPath,
				ForceQuery: _dashboard_url.ForceQuery,
				RawQuery:   _dashboard_url.RawQuery,
				Fragment:   _dashboard_url.Fragment,
			}
			return res.String()
		}
	} else {
		if payload.Username != "" && payload.Password == "" {
			res = &url.URL{
				Scheme:     _dashboard_url.Scheme,
				Opaque:     _dashboard_url.Opaque,
				User:       url.User(payload.Username),
				Host:       _dashboard_url.Host,
				Path:       _dashboard_url.Path,
				RawPath:    _dashboard_url.RawPath,
				ForceQuery: _dashboard_url.ForceQuery,
				RawQuery:   _dashboard_url.RawQuery,
				Fragment:   _dashboard_url.Fragment,
			}
			return res.String()
		} else if payload.Username != "" && payload.Password != "" {
			res = &url.URL{
				Scheme:     _dashboard_url.Scheme,
				Opaque:     _dashboard_url.Opaque,
				User:       url.UserPassword(payload.Username, payload.Password),
				Host:       _dashboard_url.Host,
				Path:       _dashboard_url.Path,
				RawPath:    _dashboard_url.RawPath,
				ForceQuery: _dashboard_url.ForceQuery,
				RawQuery:   _dashboard_url.RawQuery,
				Fragment:   _dashboard_url.Fragment,
			}
			return res.String()
		} else {
			res = &url.URL{
				Scheme:     _dashboard_url.Scheme,
				Opaque:     _dashboard_url.Opaque,
				User:       nil,
				Host:       _dashboard_url.Host,
				Path:       _dashboard_url.Path,
				RawPath:    _dashboard_url.RawPath,
				ForceQuery: _dashboard_url.ForceQuery,
				RawQuery:   _dashboard_url.RawQuery,
				Fragment:   _dashboard_url.Fragment,
			}
			return res.String()
		}
	}
}
