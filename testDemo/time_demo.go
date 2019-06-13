/*time包测试*/

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())            // 2019-06-13 11:04:44.1948166 +0800 CST m=+0.005025701
	fmt.Println(time.Now().UTC())      // 2019-06-13 03:08:39.5448048 +0000 UTC
	fmt.Println(time.Now().UnixNano()) // 1560395319544804800
	// fmt.Println(time.Duration)	// time.Duration is not a expression
	fmt.Println(time.UnixDate) // Mon Jan _2 15:04:05 MST 2006
	fmt.Println(time.Second)
	fmt.Println(time.RFC3339) // 2006-01-02T15:04:05Z07:00
	fmt.Println(time.Stamp)   // Jan _2 15:04:05
}
