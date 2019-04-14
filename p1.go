/*
题目：
已知 f(1) = 1, x>=2, f(x) = f(x-1) + [数字x中包含的“1”的个数], 求使 f(x) = x 的最小值.
*/

package main

import (
	"fmt"
	"time"
)

func main(){
	//test(fx)

	fx()
}

func test(fn func()){
	now, N := time.Now(), 10000
	for i:=0; i < N; i++{
		fn()
	}
	fmt.Println("cost time:", time.Now().Sub(now).Seconds(), "(s)")
}

func fx() {
	start, end, total, mod := 0, 9, 0, 10
	val := 0
	countOne := func(n int) (count int) {
		if n % mod == 0 {			
			n /= mod // 减少一次循环
			for n > 0 {
				if n % mod == 1 {
					count++
				}
				n /= mod
			}
			val = count
			return
		}
		count = val
		if n % mod == 1 {
			count++
		}
		return
	}
	for {
		rangeCount := countOne(start / mod) * mod + 1
		if total + rangeCount >= end {
			for i:=start; i<end; i++{
				tmp := countOne(i)
				if total + tmp == i {
					fmt.Println(i)
					break
				}
				total += tmp
			}
			break
		}
		total += rangeCount
		start += mod
		end += mod
	}
}
