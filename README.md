[![Build Status](https://travis-ci.org/howood/linkedhashmap.svg?branch=master)](https://travis-ci.org/howood/linkedhashmap)
[![GitHub release](http://img.shields.io/github/release/howood/linkedhashmap.svg?style=flat-square)][release]
[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/howood/linkedhashmap)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]

[release]: https://github.com/howood/linkedhashmap/releases
[license]: https://github.com/howood/linkedhashmap/blob/master/LICENSE

# linkedhashmap

LinkedHashMap provides map with order of elements.

# Install

```
$ go get -u github.com/howood/linkedhashmap
```

# Usage

```
	// create new
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
	// sort with key
	lMap.SortKeyAsc()
	lMap.SortKeyDesc()

	// sort with value
	lMap.SortValueAsc()
	lMap.SortValueDesc()

```
