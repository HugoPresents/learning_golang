/*
1. 编写一个函数,计算intslice([]int)中的最大值。
2. 编写一个函数,计算intslice([]int)中的最小值。
*/
package main

import "fmt"

func main() {
    slice := []int{12, 23, 11, 3, 6, 57, 44}
    max, min := max_min(slice)
    fmt.Printf("max is %d, min is %d\n", max, min)
}

func max_min(slice []int) (int, int) {
    if len(slice) == 0 {
        return 0, 0
    }
    max := slice[0]
    min := slice[0]
    for _, value := range(slice) {
        if value > max {
            max = value
        }
        if value < min {
            min = value
        }
    }
    return max, min
}