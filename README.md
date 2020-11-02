# duckduckgogo
[![Go Report Card](https://goreportcard.com/badge/github.com/sap-nocops/duckduckgogo)](https://goreportcard.com/report/github.com/sap-nocops/duckduckgogo)
![Go](https://github.com/sap-nocops/duckduckgogo/workflows/Go/badge.svg)
[![license](https://img.shields.io/github/license/sap-nocops/duckduckgogo.svg)](LICENSE)

### Simple DuckDuckGo search api
duckduckgogo implements a search api in go-lang for DuckDuckGo scraping the html search page

### Usage 

```go
package main

import (
    "fmt"
    "github.com/sap-nocops/duckduckgogo/client"
)

func main() {
    ddg := client.NewDuckDuckGoSearchClient()
    res, err := ddg.SearchLimited("antani", 10)
    if err != nil {
        fmt.Printf("error: %v", err)
        return
    }
    for i, r := range res {
        fmt.Printf("*********** RESULT %d\n", i)
        fmt.Printf("url:     %s\n", r.FormattedUrl)
        fmt.Printf("title:   %s\n", r.Title)
        fmt.Printf("snippet: %s\n", r.Snippet)
    }
}
``` 

### Testing
If you need to mock the ddg results for testing purpose you can use the `SearchClient` interface and implementing it
in your test. For example:

```go
type DdgMock struct {
    results []client.Result
}

func (d *DdgMock) Search(query string) ([]client.Result, error) {
    return d.results, nil
}

func (d *DdgMock) SearchLimited(query string, limit int) ([]client.Result, error) {
    if limit < 0 || limit > len(d.results) {
        return nil, fmt.Errorf("invalid limit")
    }
    return d.results[0:limit], nil
}
```
