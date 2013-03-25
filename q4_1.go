/*
1. 建立一个Go程序打印下面的内容(到100个字符):
A
AA
AAA 
AAAA 
AAAAA 
AAAAAA 
AAAAAAA 
...
*/

package main

import "fmt"

func main() {
    /* 题干理解错误
    for i, j := 0, 1; i < 100; i++ {
        for p := 0; p < j; p++ {
            fmt.Printf("A")
        }
        fmt.Printf("\n")
        j++
    }
    */
    var sum, line = 0, 1
// 大写冒号 循环标记
G:    for {
        for i := 0; i < line; i++ {
            fmt.Printf("A")
            sum ++
            if sum == 99 {
                // 退出指定循环
                break G
            }
        }
        fmt.Printf("\n")
        line ++
    }
    /* 我觉得这个才是理解错了~我日
    str := "A"
    for i := 0; i < 100; i++ {
        fmt.Printf("%s\n", str)
        str = str + "A"
    }
    /*
}