/* Alta3 Research | RZFeeser
   Methods - defining bookup() to increase the value of books by 1  */

// Go program to illustrate the
// method with struct type receiver

package main


import "fmt"


type author struct{

    name string
    branch string
    books int
    salary int

}

func (a *author) bookup() {
   a.books++    // increment by 1

}


func (a author) show() {

    fmt.Println("Author's name: ", a.name)
    fmt.Println("Branch Name: ", a.branch)
    fmt.Println("Number of books: ", a.books)
    fmt.Println("Salary: ", a.salary)

}


func main() {

	res := author{
		name: "Indresh",
		branch: "Raleigh",
		books: 10,
		salary: 25000,
            }

       res.show()

       res.bookup()

       res.show()


}







