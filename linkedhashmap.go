package linkedhashmap

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	iHS "github.com/howood/linkedhashmap/internal/hashmapsort"
	"sort"
)

type hashDataType string

const (
	hashDatTypeKeys   hashDataType = "keys"
	hashDatTypeValues hashDataType = "values"
)

// Element represents linked hashmap element
type Element struct {
	Key   interface{}
	Value interface{}
}

// LinkedHashMap represents linked hashmap entity
type LinkedHashMap struct {
	elements map[string]*Element
	address  []string
}

// NewLinkedHashMap create LinkedHashMap pointer
func NewLinkedHashMap() *LinkedHashMap {
	return &LinkedHashMap{
		elements: make(map[string]*Element, 0),
		address:  make([]string, 0),
	}
}

// Put puts a new element into linked hashmap
func (jc *LinkedHashMap) Put(key, value interface{}) {
	hashkey := jc.generateHash(key)
	jc.elements[hashkey] = &Element{
		Key:   key,
		Value: value,
	}
	jc.addHashkeyToAddress(hashkey)
}

// Get gets a element from linked hashmap
func (jc *LinkedHashMap) Get(key interface{}) interface{} {
	hashkey := jc.generateHash(key)
	data := jc.elements[hashkey]
	return data.Value
}

// Remove removes a element from linked hashmap
func (jc *LinkedHashMap) Remove(key interface{}) {
	hashkey := jc.generateHash(key)
	delete(jc.elements, hashkey)
	jc.removeHashkeyToAddress(hashkey)
}

// Iter iterates over each elements of linked hashmap
func (jc *LinkedHashMap) Iter() <-chan *Element {
	ch := make(chan *Element, len(jc.address))
	go func() {
		defer close(ch)
		for _, hashkey := range jc.address {
			ch <- jc.elements[hashkey]
		}
	}()
	return ch
}

// Keys gets all keys of linked hashmap
func (jc *LinkedHashMap) Keys() []interface{} {
	return jc.getKeysOrValues(hashDatTypeKeys)
}

// Values gets all values of linked hashmap
func (jc *LinkedHashMap) Values() []interface{} {
	return jc.getKeysOrValues(hashDatTypeValues)
}

// SortKeyAsc sorts linked hashmap ascending order with key
func (jc *LinkedHashMap) SortKeyAsc() {
	keys := jc.Keys()
	sort.Sort(iHS.KeySort(keys))
	jc.address = make([]string, 0, len(keys))
	for _, key := range keys {
		hashkey := jc.generateHash(key)
		jc.address = append(jc.address, hashkey)
	}
}

// SortKeyDesc sorts linked hashmap descending order with key
func (jc *LinkedHashMap) SortKeyDesc() {
	keys := jc.Keys()
	sort.Sort(sort.Reverse(iHS.KeySort(keys)))
	jc.address = make([]string, 0, len(keys))
	for _, key := range keys {
		hashkey := jc.generateHash(key)
		jc.address = append(jc.address, hashkey)
	}
}

// SortValueAsc sorts linked hashmap ascending order with value
func (jc *LinkedHashMap) SortValueAsc() {
	sortvalue := iHS.ValueSort{}
	for k, v := range jc.elements {
		sortvalue = append(sortvalue, iHS.SortElement{
			Hashkey: k,
			Value:   v.Value,
		})
	}
	sort.Sort(sortvalue)
	jc.address = make([]string, 0, len(sortvalue))
	for _, sortelement := range sortvalue {
		jc.address = append(jc.address, sortelement.Hashkey)
	}
}

// SortValueDesc sorts linked hashmap descending order with value
func (jc *LinkedHashMap) SortValueDesc() {
	sortvalue := iHS.ValueSort{}
	for k, v := range jc.elements {
		sortvalue = append(sortvalue, iHS.SortElement{
			Hashkey: k,
			Value:   v.Value,
		})
	}
	sort.Sort(sort.Reverse(sortvalue))
	jc.address = make([]string, 0, len(sortvalue))
	for _, sortelement := range sortvalue {
		jc.address = append(jc.address, sortelement.Hashkey)
	}
}

func (jc *LinkedHashMap) generateHash(key interface{}) string {
	hasher := md5.New()
	hasher.Write([]byte(fmt.Sprintf("%v", key)))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (jc *LinkedHashMap) addHashkeyToAddress(hashkey string) {
	if jc.isInAddress(hashkey) == false {
		jc.address = append(jc.address, hashkey)
	}
}

func (jc *LinkedHashMap) removeHashkeyToAddress(hashkey string) {
	newaddress := make([]string, 0, len(jc.address)-1)
	for _, key := range jc.address {
		if key != hashkey {
			newaddress = append(newaddress, key)
		}
	}
	jc.address = newaddress
}

func (jc *LinkedHashMap) isInAddress(hashkey string) bool {
	for _, key := range jc.address {
		if hashkey == key {
			return true
		}
	}
	return false
}

func (jc *LinkedHashMap) getKeysOrValues(datatype hashDataType) []interface{} {
	hashdatas := make([]interface{}, 0, len(jc.address))
	for _, hashkey := range jc.address {
		elem := jc.elements[hashkey]
		switch datatype {
		case hashDatTypeKeys:
			hashdatas = append(hashdatas, elem.Key)
		case hashDatTypeValues:
			hashdatas = append(hashdatas, elem.Value)
		}
	}
	return hashdatas
}
