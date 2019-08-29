package main

import "fmt"

func main() {

    var row, col int
    fmt.Print("enter rows cols: ")
    fmt.Scan(&row, &col)

    a := make([][]int, row)
    for i := range a {
        a[i] = make([]int, col)
    }

    fmt.Println("a[0][0] =", a[0][0])
 
    a[row-1][col-1] = 7

    fmt.Printf("a[%d][%d] = %d\n", row-1, col-1, a[row-1][col-1])

    a = nil
    
}