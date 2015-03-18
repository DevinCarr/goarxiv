// goarxiv_test v0.1
// Made by Devin Carr (github.com/devincarr/goarxiv),
// Licensed under MIT,
// Utilizes the arXiv.org API (http://arxiv.org/help/api/user-manual)
package goarxiv

import "testing"

func TestSearchAdd(t *testing.T) {
	s := New()

	s.AddQuery("start", "0")
	s.AddQuery("end", "1")
	text := s.Query.Encode()
	if text != "end=1&start=0" {
		t.Error("query had value of", text)
	}
}

func TestSearchBuild(t *testing.T) {
	s := New()

	s.AddQuery("start", "0")
	s.AddQuery("end", "1")
	s.Build()
	textValues := s.Url.Query()
	if textValues.Get("start") != "0" {
		t.Error("query start had value of", textValues)
	}
	if textValues.Get("end") != "1" {
		t.Error("query end had value of", textValues)
	}
}

func TestSearchGet(t *testing.T) {
	s := New()

	s.AddQuery("max_results", "1")
	s.AddQuery("search_query", "all:electron")
	s.AddQuery("start", "0")
	res, err := s.Get()
	if err != nil {
		t.Error("Get returned an error:", err)
	}
	if res.Title == "" {
		t.Error("results returned empty")
	}
}
