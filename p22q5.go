/*
1. 编写计算一个类型是float64的slice的平均值的代码。
*/

package main

import "fmt"

func main() {
    f_slice := []float64{12.5, 13.6, 17.85}
    var sum float64
    for _, f_per := range f_slice {
        sum += f_per
    }
    average := sum / float64(len(f_slice))
    fmt.Printf("%.2f\n", average)
}