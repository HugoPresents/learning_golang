/*
1. 编写函数接受整数类型变参，并且每行打印一个数字。
*/

package main

import "fmt"

func main() {
    ac_int(12, 11, 58)
}

func ac_int(arg ...int) {
    for _, n := range(arg) {
        fmt.Printf("%d\n", n)
    }
}