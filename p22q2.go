/*
1. 创建一个基于for的简单的循环。使其循环10次,并且使用fmt包打印 出计数器的值。

2. 用goto改写1的循环。保留字for不可使用。

3. 再次改写这个循环,使其遍历一个 array,并将这个 array 打印到屏幕上。
*/
package main

import "fmt"

// 必须要有main 函数体
func main() {
    /*for i := 0; i < 10; i++ {
        fmt.Printf("%d\n", i)
    }
    */
    /*
    i := 0
L:
    fmt.Printf("%d\n", i)
    i ++
    // if 没有括号
    if i < 10 {
        goto L
    }
    */
    arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    for _, i := range arr {
        fmt.Printf("%d\n", i)
    }
}
