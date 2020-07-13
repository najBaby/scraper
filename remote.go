package scraper

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	contentJSON = "application/json"
	contentFORM = "application/x-www-form-urlencoded"
)

type (
	// Remote is
	Remote struct {
		client *http.Client
	}

	// Config is
	Config struct {
		CheckRedirect func(req *http.Request, via []*http.Request) error
		Jar           http.CookieJar
		Timeout       time.Duration
	}

	// Options is
	Options struct {
		URL    string
		Query  map[string][]string
		Header map[string]string
		Body   interface{}
	}
)

// NewRemote is
func NewRemote(config Config) *Remote {
	return &Remote{
		client: &http.Client{
			Jar:           config.Jar,
			Timeout:       config.Timeout,
			CheckRedirect: config.CheckRedirect,
		},
	}
}

// GET is
func (remote *Remote) GET(config Options) (*http.Response, error) {
	return remote.do(http.MethodGet, contentFORM, config)
}

// HEAD is
func (remote *Remote) HEAD(config Options) (*http.Response, error) {
	return remote.do(http.MethodHead, contentFORM, config)
}

// POST is
func (remote *Remote) POST(config Options) (*http.Response, error) {
	return remote.do(http.MethodPost, contentJSON, config)
}

// PUT is
func (remote *Remote) PUT(config Options) (*http.Response, error) {
	return remote.do(http.MethodPut, contentJSON, config)
}

// PATCH is
func (remote *Remote) PATCH(config Options) (*http.Response, error) {
	return remote.do(http.MethodPatch, contentJSON, config)
}

// DELETE is
func (remote *Remote) DELETE(config Options) (*http.Response, error) {
	return remote.do(http.MethodDelete, contentJSON, config)
}

func (remote *Remote) do(method string, content string, config Options) (*http.Response, error) {
	url, body := config.parse()

	request, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		panic(err)
	}

	for k, v := range config.Header {
		request.Header.Add(k, v)
	}

	request.Header.Set("Content-Type", content)

	res, err := remote.client.Do(request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (options *Options) parse() (*url.URL, io.Reader) {

	b, err := json.Marshal(options.Body)
	if err != nil {
		panic(err)
	}

	u, err := url.Parse(options.URL)
	if err != nil {
		panic(err)
	}

	q := u.Query()
	for k, v := range options.Query {
		for _, t := range v {
			q.Add(k, t)
		}
	}
	u.RawQuery = q.Encode()

	bufReader := bytes.NewReader(b)
	return u, bufReader
}
