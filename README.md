[![GitHub release](http://img.shields.io/github/release/howood/linkedhashmap.svg?style=flat-square)][release]
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/howood/linkedhashmap)
[![Build Status](https://github.com/howood/linkedhashmap/actions/workflows/test.yml/badge.svg?branch=master)][actions]
[![Test Coverage](https://api.codeclimate.com/v1/badges/203a651b28ac2017b4a1/test_coverage)](https://codeclimate.com/github/howood/linkedhashmap/test_coverage)
[![Go Report Card](https://goreportcard.com/badge/github.com/howood/linkedhashmap)](https://goreportcard.com/report/github.com/howood/linkedhashmap)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]

[release]: https://github.com/howood/linkedhashmap/releases
[actions]: https://github.com/howood/linkedhashmap/actions
[license]: https://github.com/howood/linkedhashmap/blob/master/LICENSE

# linkedhashmap

LinkedHashMap provides map with order of elements.

# Install

```
$ go get -u github.com/howood/linkedhashmap
```

# Usage

```
	// Create new
	lMap := NewLinkedHashMap()

	// Set map data
	lMap.Put("tokyo", 3)
	lMap.Put("yokohama", 23)
	lMap.Put("kyoto", 44)
	lMap.Put("sapporo", 65)
	lMap.Put("fukuoka", 76)

	// Get map data
	val := lMap.Get("sendai") // interface{}

	// Remove map data
	lMap.Remove("sapporo")

	// Get all keys and values
	keys := lMap.Keys()
	values := lMap.Values()

	// Itretor
	for lm := range lMap.Iter() {
		log.Printf("%v: %v", lm.Key, lm.Value)
	}

	// Sort with key
	lMap.SortKeyAsc()
	lMap.SortKeyDesc()

	// Sort with value
	lMap.SortValueAsc()
	lMap.SortValueDesc()

```
