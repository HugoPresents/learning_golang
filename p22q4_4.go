/*
4. 编写一个Go程序可以逆转字符串,例如“foobar”被打印成“raboof”。
提 示:不幸的是你需要知道一些关于转换的内容,参阅 “转换” 第 63 页的 内容。
*/

package main

import "fmt"

func main() {
    str := "foobar"
    str_slice := []rune(str)
    //fmt.Printf("%v\n", str_slice)
    for i := len(str_slice) - 1; i >= 0; i -- {
        // 感觉也不是很好~
        // 实际上只管打印的话还行
        fmt.Printf("%s", string(str_slice[i]))
    }
    fmt.Printf("\n")
}