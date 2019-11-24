package linkedhashmap

import (
	"testing"
)

func Test_LinkedHashedMap(t *testing.T) {
	lMap := NewLinkedHashMap()

	lMap.Put("tokyo", 3)
	lMap.Put("yokohama", 23)
	lMap.Put("kyoto", 44)
	lMap.Put("sapporo", 65)
	lMap.Put("fukuoka", 76)
	lMap.Put("hiroshima", 74)
	lMap.Put("sendai", 67)
	lMap.Put("osaka", 89)

	t.Log(lMap.Keys())
	t.Log(lMap.Values())
	count := 0
	for lm := range lMap.Iter() {
		t.Logf("%v: %v", lm.Key, lm.Value)
		if count == 4 {
			if key, ok := lm.Key.(string); ok {
				if key != "fukuoka" {
					t.Fatalf("invalid key :%v - %v", key, "fukuoka")
				}
			} else {
				t.Fatal("key type error")
			}
		}
		count++
	}
	t.Log("success test1")

	if v := lMap.Get("sendai"); v != nil {
		if value, ok := v.(int); ok {
			if value != 67 {
				t.Fatalf("invalid value :%v - %v", value, 67)
			}
		} else {
			t.Fatal("value type error")
		}
	} else {
		t.Fatal("not exist key error")
	}
	t.Log("success test2")

	lMap.Remove("sapporo")
	lMap.Remove("xxxxxx")

	t.Log(lMap.Keys())
	t.Log(lMap.Values())
	count = 0
	for lm := range lMap.Iter() {
		t.Logf("%v: %v", lm.Key, lm.Value)
		if count == 4 {
			if key, ok := lm.Key.(string); ok {
				if key != "hiroshima" {
					t.Fatalf("invalid key :%v - %v", key, "hiroshima")
				}
			} else {
				t.Fatal("key type error")
			}
		}
		count++
	}
	t.Log("success test3")

	t.Log("success LinkedHashedMap")
}