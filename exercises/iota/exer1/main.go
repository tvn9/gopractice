// Iota
package main

import "fmt"

func main() {
    const (
        nov = (11 - iota)
        oct
        sep
    )
    fmt.Println(nov, oct, sep)
}
