/*    switch - case and default  */

package main


import ( 
	"fmt"
        "runtime"
	"strings"
)


func main() {

	var gove string = runtime.Version()
	switch  {

             case strings.Contains(gove, "go1.17"):
		     fmt.Println("you are using the latest version of GoLang")
              case strings.Contains(gove, "go1.16"), strings.Contains(gove, "go1.15"):
		     fmt.Println("You are using an older, but acceptable version of GoLang")
	  default:
		     fmt.Println("Upgrade GoLang before you continue")

	}

}
