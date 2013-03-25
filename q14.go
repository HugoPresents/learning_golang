/*
1. 编写一个针对int类型的slice冒泡排序的函数。
*/

package main

import "fmt"

func main() {
    slice := []int{1, 3, 52, 45, 33, 12, 32, 22, 11, 5411, 343215, 3213, 24132, 21, 22, 14, 21}
    fmt.Printf("%v before sort\n", slice)
    sort(slice)
    fmt.Printf("%v after sort\n", slice)
}

func sort(slice []int) {
    for key, value := range(slice) {
        for i := 0; i < key; i ++ {
            if value < slice[i] {
                slice[key], slice[i] = slice[i], slice[key]
            }
        }
    }
}
// 待完成~碎觉先。。。
func fast_sort(slice []int) {
    key := slice[0]
    i, j := 0, len(slice) - 1
    j_find := false
    i_find := true
    for j = i; j-- {
        if slice[j] > slice[i] {
            slice[j], slice[i] = slice[i], slice[j]
            j_find = true
            break
        }
    }
    if()
    for i = j; i ++ {
        if slice[i] < slice
    }
}