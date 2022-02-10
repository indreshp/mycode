/*
CHALLENGE 01 - Write a program that issues a Linux command to the local host. 
You can try anything you'd like, but something like, 
ping -c 4 or nslookup google.com 
might be fun to try.
*/


package main

import (
     "fmt"
     "os/exec"
//     "io/ioutil"
//     "strings"
     "log"
)

func main() {

    out, err := exec.Command("nslookup", "google.com").Output()

    if err != nil {
	    log.Fatal(err)
    }

    fmt.Println(string(out))

}

