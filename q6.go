/*
1. 编写一个函数用于计算一个 float64 类型的 slice 的平均值。
*/

package main

import "fmt"

func main() {
    f_slice := []float64{12.04, 12.04}
    fmt.Printf("%f\n", average(f_slice))
}

func average(f_slice []float64) float64 {
    if len(f_slice) == 0 {
        return 0
    }
    var sum float64 = 0
    for _, f := range f_slice {
        sum += f
    }
    average := sum / float64(len(f_slice))
    return average
}