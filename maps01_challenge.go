/* Indresh Padmonkar 
   CHALLENGE 01 - Mapping programming languages to extension names */

   package main

   import "fmt"


   func main() {

     var fileExtensions = map[string]string{
           "Python":  ".py",
           "C++":     ".cpp",
	   "Java":    ".java",
	   "Golang":  ".go",
	   "Ansimble":".yml",
     }

     fmt.Println(fileExtensions)

     delete(fileExtensions, "C++")

     fileExtensions["Julia"] = ".jl"

     fmt.Println(fileExtensions)

   }

