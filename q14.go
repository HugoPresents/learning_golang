/*
1. 编写一个针对int类型的slice冒泡排序的函数。
*/

package main

import "fmt"

func main() {
    slice := []int{243, 3, 52, 45, 33, 12, 32, 22, 11, 5411, 343215, 3213, 24132, 21, 22, 14, 21}
    fmt.Printf("%v before sort\n", slice)
    //sort(slice)
    quick_sort(0, len(slice)-1, slice)
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
func quick_sort(start, end int, slice []int) {
    fmt.Printf("start is %d, end is %d\n", start, end)
    if start == end {
        return
    }
    key := slice[start]
    mark := start
    for i := start; i <= end; i++ {
        fmt.Printf("key is %d, i is %d, value is %d\n", key, i, slice[i])
        if slice[i] < key {
            fmt.Printf("switch!\n")
            slice[mark], slice[i] = slice[i], slice[mark]
            mark = i
        }
    }
    if mark + 1 > end {
        return
    }
    //quick_sort(start, mark, slice)
    //quick_sort(mark + 1, end, slice)
}