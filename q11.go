/*
1. 斐波那契数列以： 1; 1; 2; 3; 5; 8; 13; : : : 开始。或者用数学形式表达：
x1 = 1; x2 = 1; xn = xn-1 + xn-2
编写一个接受 int 值的函数，并给出这个值得到的斐波那契数列。
*/

package main

import "fmt"

func main() {
    fmt.Printf("%v\n", leonardo_slice(10))
}

func leonardo_slice(n int) [10]int {
    var slice = [10]int{}
    // 这样赋值要帅气些噶
    slice[0], slice[1] = n, n
    /*
    slice[0] = n
    slice[1] = n
    */
    for i:= 2; i<10; i++ {
        slice[i] = slice[i-2] + slice[i-1]
    }
    return slice
}