/*
1. 解决这个叫做Fizz-Buzz[30]的问题:
编写一个程序,打印从 1 到 100 的数字。当是三个􏰄数就打印 “Fizz” 代替数字,当是􏰅的􏰄数就打印 “Buzz”。当数字同时是 三和􏰅的􏰄数时,打印 “FizzBuzz”。
*/
package main

import "fmt"

func main() {
    for i := 0; i < 100; i++ {
        if (i % 3 == 0) && (i % 5 == 0) {
            fmt.Printf("%d FizzBuzz\n", i)
            continue
        }
        if i % 3 == 0 {
            fmt.Printf("%d Fizz\n", i)
            continue
        }
        if i % 5 == 0 {
            fmt.Printf("%d Buzz\n", i)
            continue
        }
        fmt.Printf("%d\n", i)
    }
}