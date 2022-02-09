

package main

import (
    "fmt"
    "os"
)


func main() {

	arguments := os.Args
	for i, a := range arguments {
            fmt.Printf("argument %s at index %d\n", a,i)

	}

}
