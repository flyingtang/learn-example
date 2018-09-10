
package interview

import (
	"sort"
)

// 获取一个map 下有序key下的value

func GetOrderInt(m map[string]int) (res []int){

	var keys []string
	for key, _ := range m {
		keys = append(keys, key)
	}
	
	sort.Strings(keys)
	
	for _, value := range keys {
		res = append(res, m[value])
	}
	return 
}
