/*
1. 编写函数，返回其（两个）参数正确的（自然）数字顺序：
f(7,2) → 2,7
f(2,7) → 2,7
*/

package main
import "fmt"

func main() {
    // 实在是想把下面两行合成一行，求破
    a, b :=  sort(7, 2)
    fmt.Printf("%d → %d\n", a, b)
}

func sort(a, b int) (int, int) {
    if b > a {
        return a, b
    }
    return b, a
}