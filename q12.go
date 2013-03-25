/*
(1) Map function map() 函数是一个接受一个函数和一个列表作为参数的函 数。函数应用于列表中的每个元素,而一个新的包含有计算结果的列表被返 回。因此:
map(f(), (a1, a2, . . . , an−1, an)) = (f(a1), f(a2), . . . , f(an−1), f(an))
1. 编写Go中的简单的map()函数。它能工作于操作整数的函数就可以了。
2. 扩展代码使其工作于字符串列表。
*/

package main

import (
    "fmt"
    "strings"
)

func main() {
    f := func (a int) int {
        return a * a
    }
    f_str := func (s string) string {
        return s + s
    }
    slice := []int{1, 2, 3, 4}
    str := "妈卖批"
    str_mapped := my_map_str(f_str, str)
    my_map(f, slice)
    fmt.Printf("%d\n", slice)
    fmt.Printf("%s\n", str_mapped)
}

func my_map(f func(int) int, slice []int) {
    for key, value := range slice {
        slice[key] = f(value)
    }
    return
}

// 还要引入 strings 才好搞，我擦
func my_map_str(f func(string) string, str string) string {
    str_slice := strings.Split(str,"")
    for key, value := range(str_slice) {
        str_slice[key] = value + value
    }
    return strings.Join(str_slice, "")
}