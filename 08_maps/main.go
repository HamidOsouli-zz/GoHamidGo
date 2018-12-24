package main

import "fmt"

func main() {
	// define map
	//emails := make(map[string] string)

	// assign key values
	//emails["bob"] = "bob@gmail.com"
	//emails["sharon"] = "sharon@gmail.com"
	//emails["mike"] = "mike@gmail.com"
	// DECLARE map at the initialization time
	emails := map[string]string{"bob": "bob@gmil.com"}
	emails["mike"] = "mike@gmail.com"
	fmt.Println(emails)
	fmt.Println(len(emails))
	// delete one
	delete(emails, "bob")
	fmt.Println(emails)
	fmt.Println(len(emails))

}
