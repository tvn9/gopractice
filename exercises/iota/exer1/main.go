// Iota
package main

import "fmt"

func main() {
    const (
        nov = iota + 11
        oct
        sep
    )
    fmt.Println(nov, oct, sep)
}
