// goarxiv v0.1
// Made by Devin Carr (github.com/devincarr/goarxiv),
// Licensed under MIT,
// Utilizes the arXiv.org API (http://arxiv.org/help/api/user-manual)
package goarxiv

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	atom "golang.org/x/tools/blog/atom"
)

// Search represents a single instance of a search query to arxiv.org
type Search struct {
	Query  url.Values
	Url    url.URL
	Result *atom.Feed
}

// New constructs and returns a new Search.
func New() *Search {
	return &Search{
		Query: url.Values{},
		Url: url.URL{
			Scheme: "http",
			Host:   "export.arxiv.org",
			Path:   "api/query",
		},
		Result: &atom.Feed{},
	}
}

// Build checks and completes the search query string.
func (s *Search) Build() error {
	query := s.Query.Encode()
	if query == "" {
		return errors.New("Search query is empty")
	} else {
		s.Url.RawQuery = query
		return nil
	}
}

// Get parses the Url and completes an HTTP request.
func (s *Search) Get() (*atom.Feed, error) {
	if s.Url.RawQuery == "" {
		s.Build()
	}
	resp, errHttp := http.Get(s.Url.String())
	defer resp.Body.Close()
	if errHttp == nil {
		result, errRead := Read(resp)
		if errRead == nil {
			s.Result = result
			return result, nil
		} else {
			return &atom.Feed{}, errRead
		}
	} else {
		return &atom.Feed{}, errHttp
	}
}

// AddQuery adds a new search query (key=value) to the Search.
func (s *Search) AddQuery(key string, value string) error {
	if s.Query.Get(key) != "" {
		return errors.New("Key already exists")
	} else {
		s.Query.Add(key, value)
		return nil
	}
}

// Read the Http response and return an atom.Feed
func Read(res *http.Response) (*atom.Feed, error) {
	result := atom.Feed{}
	body, errIoRead := ioutil.ReadAll(res.Body)
	if errIoRead == nil {
		errUnmarshal := xml.Unmarshal([]byte(body), &result)
		if errUnmarshal == nil {
			return &result, nil
		} else {
			return &atom.Feed{}, errUnmarshal
		}
	} else {
		return &atom.Feed{}, errIoRead
	}
}
