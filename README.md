# goarxiv [![Build Status](https://travis-ci.org/DevinCarr/goarxiv.svg)](https://travis-ci.org/DevinCarr/goarxiv) [![GoDoc](https://godoc.org/github.com/devincarr/goarxiv?status.svg)](http://godoc.org/github.com/devincarr/goarxiv)

A Go API wrapper around the arXiv.org [API](arxiv.org/help/api/user-manual). Provides simple access to the search query and returns in a simple Atom XML struct format.

## How to
```shell
go get github.com/devincarr/goarxiv
```
and
```Go
import github.com/devincarr/goarxiv
```
#### New Search
```Go
// Build a new search struct
s := goarxiv.New()

// Add some search parameters
s.AddQuery("search_query","cat:cs.LG")

// Get the results and print first article
result, error := s.Get()
if error != nil { error }
fmt.Println(result.Entry[0].Title)

// "Hierarchical Reinforcement Learning with the MAXQ Value Function
  Decomposition"
```

#### Result Format
The results from the arXiv.org API are in Atom XMl format and goarxiv utilizes `golang.org/x/tools/blog/atom` for the results parsing. See [godoc/atom](https://godoc.org/golang.org/x/tools/blog/atom) for the format of the results.

## License
The MIT License (MIT)
