package hashmapsort

import (
	"fmt"
)

// KeySort represents sort struct for key
type KeySort []interface{}

func (k KeySort) Len() int {
	return len(k)
}

func (k KeySort) Less(i, j int) bool {
	switch conveted := k[i].(type) {
	case int:
		return conveted < k[j].(int)
	case int32:
		return conveted < k[j].(int32)
	case int64:
		return conveted < k[j].(int64)
	case float32:
		return conveted < k[j].(float32)
	case float64:
		return conveted < k[j].(float64)
	default:
		return fmt.Sprintf("%v", k[i]) < fmt.Sprintf("%v", k[j])
	}
}

func (k KeySort) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}

// SortElement represents sort element struct
type SortElement struct {
	Hashkey string
	Value   interface{}
}

// ValueSort represents sort struct for Value
type ValueSort []SortElement

func (k ValueSort) Len() int {
	return len(k)
}

func (k ValueSort) Less(i, j int) bool {
	switch conveted := k[i].Value.(type) {
	case int:
		return conveted < k[j].Value.(int)
	case int32:
		return conveted < k[j].Value.(int32)
	case int64:
		return conveted < k[j].Value.(int64)
	case float32:
		return conveted < k[j].Value.(float32)
	case float64:
		return conveted < k[j].Value.(float64)
	default:
		return fmt.Sprintf("%v", k[i].Value) < fmt.Sprintf("%v", k[j].Value)
	}
}

func (k ValueSort) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}
