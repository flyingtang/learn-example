
package main

import (
"sync"
)

var once sync.Once

//func main() {
//
//	for i, v := range make([]string, 10) {
//		once.Do(onces)
//		fmt.Println("count:", v, "---", i)
//	}
//	for i := 0; i < 10; i++ {
//
//		go func() {
//			once.Do(onced)
//			fmt.Println("213")
//		}()
//	}
//	time.Sleep(4000)
//}
//func onces() {
//	fmt.Println("onces")
//}
//func onced() {
//	fmt.Println("onced")
//
//}
//onces
//count:  --- 0
//.
//.
//.
//count:  --- 9
//213
//.
//.
//.
//213

//  sync.once 中的 once.Do(fn)  只执行一次
